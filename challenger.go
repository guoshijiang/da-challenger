package challenger

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shurcooL/graphql"
	"google.golang.org/grpc"

	datalayr "github.com/Layr-Labs/datalayr/common/contracts"
	kzg "github.com/Layr-Labs/datalayr/common/crypto/go-kzg-bn254"
	chain "github.com/Layr-Labs/datalayr/middleware/rollup-example/utils/contracts"
	rc "github.com/mantlenetworkio/da-challenger/common/BVM_EigenDataLayrChain"

	"github.com/Layr-Labs/datalayr/common/graphView"
	pb "github.com/Layr-Labs/datalayr/common/interfaces/interfaceRetrieverServer"
	"github.com/Layr-Labs/datalayr/common/logging"
)

type KzgConfig struct {
	G1Path    string
	G2Path    string
	TableDir  string
	NumWorker int
	Order     uint64 // Order is the total size of SRS
}

type ChallengerSettings struct {
	ChainSettings     chain.ChainSettings
	RollupSettings    RollupSettings
	RetrieverSettings RetrieverSettings
	KzgConfig         KzgConfig

	LoggingConfig   logging.Config
	GraphEndpoint   string
	FromStoreNumber uint64
}

type Challenger struct {
	ChallengerSettings

	ChainClient *chain.ChainClient
	GraphClient *graphView.GraphClient
	Logger      *logging.Logger

	lastStoreNumber uint64
	Timeout         time.Duration
}

type RollupSettings struct {
	Address common.Address
}

type RetrieverSettings struct {
	Socket string
}

type Fraud struct {
	StartingIndex int
}

type DataLayrDisclosureProof struct {
	Header                    []byte
	Polys                     [][]byte
	MultirevealProofs         []datalayr.MultiRevealProof
	BatchPolyEquivalenceProof [4]*big.Int
	StartingChunkIndex        int
}

type FraudProof struct {
	DataLayrDisclosureProof
	StartingSymbolIndex int
}

//If any datastore contains this fraud string
const fraudString = "2d5f2860204f2060295f2d202d5f2860206f2060295f2d202d5f286020512060295f2d2042495444414f204a5553542052454b5420594f55207c5f2860204f2060295f7c202d207c5f2860206f2060295f7c202d207c5f286020512060295f7c"

func NewChallenger(
	chainClient *chain.ChainClient,
	graphClient *graphView.GraphClient,
	logger *logging.Logger,
	settings ChallengerSettings,
	lastStoreNumber uint64,
) *Challenger {

	timeout, err := time.ParseDuration("12s")
	if err != nil {
		logger.Fatal().Err(err).Msg("Improper timeout from config")
	}

	return &Challenger{
		ChallengerSettings: settings,
		ChainClient:        chainClient,
		GraphClient:        graphClient,
		Logger:             logger,
		lastStoreNumber:    lastStoreNumber,
		Timeout:            timeout,
	}
}

//this function starts a loop that checks for fraud twice per second
func (c *Challenger) FraudCheckLoop() {

	// TODO: Handle errors
	ticker := time.NewTicker(500 * time.Millisecond)

	for {
		<-ticker.C

		//fetch the next datastore
		store, err := c.getNextDataStore()
		if err != nil {
			continue
		}

		obj, _ := json.Marshal(store)
		c.Logger.Info().Msg("Got store:" + string(obj))

		data, frames, err := c.callRetrieve(store)
		//retrieve the data associated with the store

		if err != nil {
			c.Logger.Error().Err(err).Msg("Error getting data")
			continue
		}

		c.Logger.Info().Msg("Got data:" + hexutil.Encode(data))

		//check if the fraud string exists within the data
		fraud, exists := c.checkForFraud(store, data)
		if !exists {
			log.Println("No fraud", err)
			continue
		}

		obj, _ = json.Marshal(fraud)
		c.Logger.Info().Msg("Found fraud:" + string(obj))

		proof, err := c.constructFraudProof(store, data, fraud, frames)
		//construct the fraud proof if fraud was found

		if err != nil {
			c.Logger.Error().Err(err).Msg("Error constructing fraud")
			continue
		}

		obj, _ = json.Marshal(proof)
		c.Logger.Info().Msg("Fraud proof:" + string(obj))

		obj, _ = json.Marshal(store)
		c.Logger.Info().Msg("Store:" + string(obj))

		//post the fraud proof to the chain
		tx, err := c.postFraudProof(store, proof)
		if err != nil {
			c.Logger.Error().Err(err).Msg("Error posting fraud proof")
			continue
		}
		c.Logger.Info().Msg("Fraud proof tx:" + tx.Hash().Hex())

	}

}

//gets the next datastore from the graphql db that is after the most recent one the challenger has seen
func (c *Challenger) getNextDataStore() (*graphView.DataStore, error) {

	var query struct {
		DataStores []graphView.DataStoreGql `graphql:"dataStores(first:1,where:{storeNumber_gt: $lastStoreNumber,confirmer: $confirmer,confirmed:true})"`
	}
	variables := map[string]interface{}{
		"lastStoreNumber": graphql.String(fmt.Sprint(c.lastStoreNumber)),
		"confirmer":       graphql.String(strings.ToLower(c.RollupSettings.Address.Hex())),
	}

	client := graphql.NewClient(c.GraphClient.GetEndpoint(), nil)
	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		c.Logger.Error().Err(err).Msg("GetExpiringDataStores error")
		return nil, err
	}

	if len(query.DataStores) == 0 {
		return nil, errors.New("no new stores")
	}

	store, err := query.DataStores[0].Convert()
	if err != nil {
		return nil, errors.New("conversion error")
	}
	fmt.Println("challenger get store", store.StoreNumber)
	//set the new lastStoreNumber to the retrieved store's number
	c.lastStoreNumber = uint64(store.StoreNumber)

	return store, nil

}

func (c *Challenger) callRetrieve(store *graphView.DataStore) ([]byte, []datalayr.Frame, error) {

	socket := c.RetrieverSettings.Socket
	fmt.Println("socket", socket)
	conn, err := grpc.Dial(socket, grpc.WithInsecure())
	if err != nil {
		c.Logger.Printf("Disperser Cannot connect to %v. %v\n", socket, err)
		return nil, nil, err
	}
	defer conn.Close()
	client := pb.NewDataRetrievalClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	opt := grpc.MaxCallRecvMsgSize(1024 * 1024 * 300)
	request := &pb.FramesAndDataRequest{
		DataStoreId: store.StoreNumber,
	}

	reply, err := client.RetrieveFramesAndData(ctx, request, opt)

	if err != nil {
		return nil, nil, err
	}

	data := reply.GetData()

	framesBytes := reply.GetFrames()
	header, err := datalayr.DecodeDataStoreHeader(store.Header)
	if err != nil {
		c.Logger.Printf("Could not decode header %v. %v\n", header, err)
		return nil, nil, err
	}
	frames := make([]datalayr.Frame, header.NumSys+header.NumPar)
	for i, frameBytes := range framesBytes {
		frame, err := datalayr.DecodeFrame(frameBytes)
		if err == nil {
			frames[i] = frame
		} else {
			return nil, nil, errors.New("Does not Contain all the frames")
		}
	}
	return data, frames, nil
}

//check if the fraud string exists within the data
func (c *Challenger) checkForFraud(store *graphView.DataStore, data []byte) (*Fraud, bool) {
	dataString := hex.EncodeToString(data)
	index := strings.Index(dataString, fraudString)
	if index != -1 {
		//change char index to byte index
		return &Fraud{StartingIndex: index / 2}, true
	}
	return nil, false
}

func (c *Challenger) constructFraudProof(store *graphView.DataStore, data []byte, fraud *Fraud, frames []datalayr.Frame) (*FraudProof, error) {
	//encode data to frames here

	header, err := datalayr.DecodeDataStoreHeader(store.Header)
	if err != nil {
		c.Logger.Printf("Could not decode header %v. %v\n", header, err)
		return nil, err
	}
	config := c.ChallengerSettings.KzgConfig

	s1 := kzg.ReadG1Points(config.G1Path, config.Order, config.NumWorker)
	s2 := kzg.ReadG2Points(config.G2Path, config.Order, config.NumWorker)

	dp := datalayr.NewDisclosureProver(s1, s2)

	//there are 31 bytes per fr so there are 31*chunkLenE bytes in each chunk
	//so the i'th byte starts at the (i/(31*encoder.EncodingParams.ChunkLenE))'th chunk
	startingChunkIndex := fraud.StartingIndex / int(31*header.Degree)
	//the fraud string ends len(fraudString)/2 bytes later
	endingChunkIndex := (fraud.StartingIndex + len(fraudString)/2) / int(31*header.Degree)
	startingSymbolIndex := fraud.StartingIndex % int(31*header.Degree)
	//do some math to shift this over by the correct number of bytes
	//there are 32 bytes in the actual poly for every 31 bytes in the data, hence (startingSymbolIndex/31)*32
	//then we shift over by 1 to get past the first 0 byte, and then (startingSymbolIndex % 31)
	startingSymbolIndex = (startingSymbolIndex/31)*32 + 1 + (startingSymbolIndex % 31)

	fmt.Println("INDEX", fraud.StartingIndex, startingSymbolIndex)

	//generate parameters for proving data on chain
	//this is
	//	polys: the []byte representation of the polynomials
	//	multirevealProofs: the openings of each polynomial against the full polynomial commitment
	//  batchPolyEquivalenceProof: the proof that the `polys` are in fact represented by the commitments in `multirevealProofs`
	polys, multirevealProofs, batchPolyEquivalenceProof, err := dp.ProveBatchInterpolatingPolyDisclosure(frames[startingChunkIndex:endingChunkIndex+1], store.DataCommitment, store.Header, uint32(startingChunkIndex))
	if err != nil {
		return nil, err
	}

	disclosureProof := DataLayrDisclosureProof{
		Header:                    store.Header,
		Polys:                     polys,
		MultirevealProofs:         multirevealProofs,
		BatchPolyEquivalenceProof: batchPolyEquivalenceProof,
		StartingChunkIndex:        startingChunkIndex,
	}

	return &FraudProof{DataLayrDisclosureProof: disclosureProof, StartingSymbolIndex: startingSymbolIndex}, nil
}

//converts the fraud proof object to a format that the rollup contract can handle
func (fp *DataLayrDisclosureProof) ToDisclosureProofs() rc.BVM_EigenDataLayrChainDisclosureProofs {

	proofs := make([]rc.DataLayrDisclosureLogicMultiRevealProof, 0)

	for _, oldProof := range fp.MultirevealProofs {
		newProof := rc.DataLayrDisclosureLogicMultiRevealProof{
			InterpolationPoly: rc.BN254G1Point{X: oldProof.InterpolationPolyCommit[0], Y: oldProof.InterpolationPolyCommit[1]},
			RevealProof:       rc.BN254G1Point{X: oldProof.RevealProof[0], Y: oldProof.RevealProof[1]},
			ZeroPoly: rc.BN254G2Point{
				//preserve this ordering for dumb precompile reasons
				X: [2]*big.Int{oldProof.ZeroPolyCommit[1], oldProof.ZeroPolyCommit[0]},
				Y: [2]*big.Int{oldProof.ZeroPolyCommit[3], oldProof.ZeroPolyCommit[2]},
			},
			ZeroPolyProof: oldProof.ZeroPolyProof,
		}
		proofs = append(proofs, newProof)
	}

	return rc.BVM_EigenDataLayrChainDisclosureProofs{
		Header:            fp.Header,
		FirstChunkNumber:  uint32(fp.StartingChunkIndex),
		Polys:             fp.Polys,
		MultiRevealProofs: proofs,
		PolyEquivalenceProof: rc.BN254G2Point{
			//preserve this ordering for dumb precompile reasons
			X: [2]*big.Int{fp.BatchPolyEquivalenceProof[1], fp.BatchPolyEquivalenceProof[0]},
			Y: [2]*big.Int{fp.BatchPolyEquivalenceProof[3], fp.BatchPolyEquivalenceProof[2]},
		},
	}

}

func (c *Challenger) postFraudProof(store *graphView.DataStore, fraudProof *FraudProof) (*types.Transaction, error) {
	rollup, err := c.getRollupContractBinding()
	if err != nil {
		return nil, err
	}
	//prepare the datastore's serach data for metadata proving
	searchData := rc.IDataLayrServiceManagerDataStoreSearchData{
		Duration:  store.Duration,
		Timestamp: new(big.Int).SetUint64(uint64(store.InitTime)),
		Index:     store.Index,
		Metadata: rc.IDataLayrServiceManagerDataStoreMetadata{
			HeaderHash:          store.DataCommitment,
			DurationDataStoreId: store.DurationDataStoreId,
			GlobalDataStoreId:   store.StoreNumber,
			BlockNumber:         store.StakesFromBlockNumber,
			Fee:                 store.Fee,
			Confirmer:           common.HexToAddress(store.Confirmer),
			SignatoryRecordHash: store.SignatoryRecord,
		},
	}

	//convert the disclosure proof to chain's expected format
	disclosureProofs := fraudProof.ToDisclosureProofs()

	auth := c.ChainClient.PrepareAuthTransactor()
	//get the corresponding rollup store number
	fraudStoreNumber, err := rollup.DataStoreIdToRollupStoreNumber(&bind.CallOpts{}, store.StoreNumber)
	if err != nil {
		return nil, err
	}

	//send the fraud proof to the chain
	tx, err := rollup.ProveFraud(auth, fraudStoreNumber, new(big.Int).SetUint64(uint64(fraudProof.StartingSymbolIndex)), searchData, disclosureProofs)

	return tx, err
}

//gets the binding of the rollup contract
func (c *Challenger) getRollupContractBinding() (*rc.Binding, error) {

	rollup, err := rc.NewBinding(c.RollupSettings.Address, c.ChainClient.Client)
	if err != nil {
		return nil, err
	}

	return rollup, nil

}
