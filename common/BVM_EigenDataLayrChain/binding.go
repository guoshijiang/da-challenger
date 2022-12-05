// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// BVM_EigenDataLayrChainDisclosureProofs is an auto generated low-level Go binding around an user-defined struct.
type BVM_EigenDataLayrChainDisclosureProofs struct {
	Header               []byte
	FirstChunkNumber     uint32
	Polys                [][]byte
	MultiRevealProofs    []DataLayrDisclosureLogicMultiRevealProof
	PolyEquivalenceProof BN254G2Point
}

// DataLayrDisclosureLogicMultiRevealProof is an auto generated low-level Go binding around an user-defined struct.
type DataLayrDisclosureLogicMultiRevealProof struct {
	InterpolationPoly BN254G1Point
	RevealProof       BN254G1Point
	ZeroPoly          BN254G2Point
	ZeroPolyProof     []byte
}

// IDataLayrServiceManagerDataStoreMetadata is an auto generated low-level Go binding around an user-defined struct.
type IDataLayrServiceManagerDataStoreMetadata struct {
	HeaderHash          [32]byte
	DurationDataStoreId uint32
	GlobalDataStoreId   uint32
	BlockNumber         uint32
	Fee                 *big.Int
	Confirmer           common.Address
	SignatoryRecordHash [32]byte
}

// IDataLayrServiceManagerDataStoreSearchData is an auto generated low-level Go binding around an user-defined struct.
type IDataLayrServiceManagerDataStoreSearchData struct {
	Metadata  IDataLayrServiceManagerDataStoreMetadata
	Duration  uint8
	Timestamp *big.Int
	Index     uint32
}

// BindingMetaData contains all meta data concerning the Binding contract.
var BindingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupStoreNumber\",\"type\":\"uint32\"}],\"name\":\"RollupStoreConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"dataStoreId\",\"type\":\"uint32\"}],\"name\":\"RollupStoreInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupStoreNumber\",\"type\":\"uint32\"}],\"name\":\"RollupStoreReverted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCK_STALE_MEASURE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FRAUD_STRING\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"headerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"durationDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"globalDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fee\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"confirmer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreSearchData\",\"name\":\"searchData\",\"type\":\"tuple\"}],\"name\":\"confirmData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataManageAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"dataStoreIdToRollupStoreNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fraudProofPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dataManageAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"polys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"parse\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fraudulentStoreNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"headerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"durationDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"globalDataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"fee\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"confirmer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structIDataLayrServiceManager.DataStoreSearchData\",\"name\":\"searchData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"firstChunkNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"polys\",\"type\":\"bytes[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"interpolationPoly\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"revealProof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"zeroPoly\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"zeroPolyProof\",\"type\":\"bytes\"}],\"internalType\":\"structDataLayrDisclosureLogic.MultiRevealProof[]\",\"name\":\"multiRevealProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"polyEquivalenceProof\",\"type\":\"tuple\"}],\"internalType\":\"structBVM_EigenDataLayrChain.DisclosureProofs\",\"name\":\"disclosureProofs\",\"type\":\"tuple\"}],\"name\":\"proveFraud\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupStoreNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rollupStores\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"dataStoreId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"confirmAt\",\"type\":\"uint32\"},{\"internalType\":\"enumBVM_EigenDataLayrChain.RollupStoreStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"duration\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalOperatorsIndex\",\"type\":\"uint32\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405262015180609a5534801561001757600080fd5b50612436806100276000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80637ef01d5e11610097578063bcb8ec9011610066578063bcb8ec901461026f578063f249502914610282578063f2a8f124146102a2578063f2fde38b146102ab57600080fd5b80637ef01d5e146101cb5780638da5cb5b146101de578063b0393a37146101fc578063b537c4c71461024f57600080fd5b80635c1bba38116100d35780635c1bba38146101625780635e8b3f2d146101a7578063715018a6146101b05780637bd85879146101b857600080fd5b80631f944c8f14610105578063428bba091461012e57806346b2eb9b14610145578063485cc9551461014d575b600080fd5b61011861011336600461179e565b6102be565b6040516101259190611850565b60405180910390f35b610137609d5481565b604051908152602001610125565b610118610427565b61016061015b3660046118ca565b610443565b005b6097546101829073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610125565b61013760995481565b61016061061d565b6101606101c6366004611aac565b610631565b6101606101d9366004611b5f565b610dd2565b60335473ffffffffffffffffffffffffffffffffffffffff16610182565b61024061020a366004611bb5565b609b6020526000908152604090205463ffffffff8082169164010000000081049091169068010000000000000000900460ff1683565b60405161012593929190611bfd565b61013761025d366004611c57565b609c6020526000908152604090205481565b61016061027d366004611c7b565b61110b565b6098546101829073ffffffffffffffffffffffffffffffffffffffff1681565b610137609a5481565b6101606102b9366004611cf7565b6113c5565b6060806000845b848351101561041b576000610340826102df602082611d41565b6102ea906001611d7c565b6102f5906020611d94565b6102ff9190611dd1565b855161030b9089611dd1565b848c8c8881811061031e5761031e611de8565b90506020028101906103309190611e17565b61033b929150611dd1565b611462565b90508389898581811061035557610355611de8565b90506020028101906103679190611e17565b84906103738583611d7c565b9261038093929190611e7c565b60405160200161039293929190611ea6565b6040516020818303038152906040529350835186116103b1575061041b565b8888848181106103c3576103c3611de8565b90506020028101906103d59190611e17565b90506103e18284611d7c565b14156103fd57826103f181611ece565b93505060019150610415565b610408816001611d7c565b6104129083611d7c565b91505b506102c5565b50909695505050505050565b6040518060800160405280606081526020016123a16060913981565b600054610100900460ff16158080156104635750600054600160ff909116105b8061047d5750303b15801561047d575060005460ff166001145b6104f45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561055257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61055a611495565b6097805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560988054928516929091169190911790556064609955801561061857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b505050565b61062561151a565b61062f6000611581565b565b6000848152609b602090815260408083208151606081018352815463ffffffff80821683526401000000008204169482019490945292909183019068010000000000000000900460ff16600281111561068c5761068c611bce565b600281111561069d5761069d611bce565b90525090506001816040015160028111156106ba576106ba611bce565b1480156106d0575042816020015163ffffffff16115b6107425760405162461bcd60e51b815260206004820152602d60248201527f526f6c6c757053746f7265206d75737420626520636f6d6d697474656420616e60448201527f6420756e636f6e6669726d65640000000000000000000000000000000000000060648201526084016104eb565b825161074d906115f8565b6098546020850151604080870151606088015191517fed82c0ee00000000000000000000000000000000000000000000000000000000815260ff9093166004840152602483015263ffffffff16604482015273ffffffffffffffffffffffffffffffffffffffff9091169063ed82c0ee9060640160206040518083038186803b1580156107d957600080fd5b505afa1580156107ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108119190611f07565b1461085e5760405162461bcd60e51b815260206004820152601e60248201527f6d6574616461746120707265696d61676520697320696e636f7272656374000060448201526064016104eb565b805183516040015163ffffffff9081169116146109095760405162461bcd60e51b815260206004820152604260248201527f7365616368446174612773206461746173746f7265206964206973206e6f742060448201527f636f6e73697374656e74207769746820676976656e20726f6c6c75702073746f60648201527f7265000000000000000000000000000000000000000000000000000000000000608482015260a4016104eb565b6109138280611e17565b604051610921929190611f20565b604051908190039020835151146109a05760405162461bcd60e51b815260206004820152603260248201527f646973636c6f737572652070726f6f667320686561646572686173682070726560448201527f696d61676520697320696e636f7272656374000000000000000000000000000060648201526084016104eb565b73__$487cc7a1b58f8f6e007e8cd55d94b79fe8$__63899c1afc6109c48480611e17565b6109d46040870160208801611c57565b6109e16040880188611f30565b6109ee60608a018a611f30565b8a6080016040518963ffffffff1660e01b8152600401610a1598979695949392919061215e565b60206040518083038186803b158015610a2d57600080fd5b505af4158015610a41573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a65919061222c565b610ab15760405162461bcd60e51b815260206004820152601d60248201527f646973636c6f737572652070726f6f66732061726520696e76616c696400000060448201526064016104eb565b600073__$487cc7a1b58f8f6e007e8cd55d94b79fe8$__63e34c39d6610ad78580611e17565b6040518363ffffffff1660e01b8152600401610af492919061224e565b60206040518083038186803b158015610b0c57600080fd5b505af4158015610b20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b449190612262565b905063ffffffff8116610b5a6040850185611f30565b9050610b6c6040860160208701611c57565b63ffffffff16610b7c9190611d7c565b1115610bf05760405162461bcd60e51b815260206004820152602e60248201527f43616e206f6e6c792070726f766520646174612066726f6d207468652073797360448201527f74656d61746963206368756e6b7300000000000000000000000000000000000060648201526084016104eb565b6000610c22610c026040860186611f30565b886040518060800160405280606081526020016123a160609139516102be565b90506040518060800160405280606081526020016123a16060913951815114610cd95760405162461bcd60e51b815260206004820152604260248201527f50617273696e67206572726f722c2070726f76656e20737472696e672069732060448201527f646966666572656e74206c656e677468207468616e206672617564207374726960648201527f6e67000000000000000000000000000000000000000000000000000000000000608482015260a4016104eb565b6040518060800160405280606081526020016123a16060913980519060200120818051906020012014610d4e5760405162461bcd60e51b815260206004820152601d60248201527f70726f76656e20737472696e6720213d20667261756420737472696e6700000060448201526064016104eb565b6000878152609b60209081526040918290208054680200000000000000007fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff909116179055905163ffffffff891681527f18407f85390b98bfd30ac355138feffbf9d2519036bddba3429ec57ed328e7a8910160405180910390a150505050505050565b60975473ffffffffffffffffffffffffffffffffffffffff163314610e5f5760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79207468652073657175656e6365722063616e2073746f72652064617460448201527f610000000000000000000000000000000000000000000000000000000000000060648201526084016104eb565b805160409081015163ffffffff166000908152609c60205220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14610f335760405162461bcd60e51b815260206004820152605560248201527f446174612073746f72652065697468657220776173206e6f7420696e6974696160448201527f6c697a65642062792074686520726f6c6c757020636f6e74726163742c206f7260648201527f20697320616c726561647920636f6e6669726d65640000000000000000000000608482015260a4016104eb565b6098546040517f5189951500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690635189951590610f8d9086908690869060040161227f565b600060405180830381600087803b158015610fa757600080fd5b505af1158015610fbb573d6000803e3d6000fd5b50505050604051806060016040528082600001516040015163ffffffff168152602001609a5442610fec9190611d7c565b63ffffffff16815260200160019052609d546000908152609b6020908152604091829020835181549285015163ffffffff908116640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009094169116179190911780825591830151909182907fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff166801000000000000000083600281111561109b5761109b611bce565b021790555050609d8054835160409081015163ffffffff166000908152609c6020529081208290557fe9f8c90baa6e73f4fbc1350dac0cf673eabb9d0bf5eef014ce5fe08be7d2d7a993509091906110f283611ece565b9091555060405163ffffffff909116815260200161060f565b60975473ffffffffffffffffffffffffffffffffffffffff1633146111985760405162461bcd60e51b815260206004820152602160248201527f4f6e6c79207468652073657175656e6365722063616e2073746f72652064617460448201527f610000000000000000000000000000000000000000000000000000000000000060648201526084016104eb565b6099546111ab63ffffffff841643611dd1565b106111f85760405162461bcd60e51b815260206004820152601e60248201527f7374616b65732074616b656e2066726f6d20746f6f206c6f6e672061676f000060448201526064016104eb565b609854604080517f72d18e8d000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff16916372d18e8d916004808301926020929190829003018186803b15801561126357600080fd5b505afa158015611277573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061129b9190612262565b6098546040517fdcf49ea700000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff169063dcf49ea7906112fe90339030908990899089908e908e90600401612349565b602060405180830381600087803b15801561131857600080fd5b505af115801561132c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113509190612262565b5063ffffffff81166000818152609c60209081526040918290207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff905590519182527f957f0dd1f1ce8fbaa766e73503339f17b04cfbbd7e0db44e9460644485b813b5910160405180910390a1505050505050565b6113cd61151a565b73ffffffffffffffffffffffffffffffffffffffff81166114565760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104eb565b61145f81611581565b50565b600082841061147e57818310611478578161148d565b8261148d565b81841061148b578161148d565b835b949350505050565b600054610100900460ff166115125760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104eb565b61062f611718565b60335473ffffffffffffffffffffffffffffffffffffffff16331461062f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104eb565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600080826000015183602001518460400151856060015186608001518760a001518860c001516040516020016116db979695949392919096875260e095861b7fffffffff00000000000000000000000000000000000000000000000000000000908116602089015294861b851660248801529290941b909216602885015260a09190911b7fffffffffffffffffffffffff000000000000000000000000000000000000000016602c84015260609190911b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166038830152604c820152606c0190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209392505050565b600054610100900460ff166117955760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104eb565b61062f33611581565b600080600080606085870312156117b457600080fd5b843567ffffffffffffffff808211156117cc57600080fd5b818701915087601f8301126117e057600080fd5b8135818111156117ef57600080fd5b8860208260051b850101111561180457600080fd5b6020928301999098509187013596604001359550909350505050565b60005b8381101561183b578181015183820152602001611823565b8381111561184a576000848401525b50505050565b602081526000825180602084015261186f816040850160208701611820565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146118c557600080fd5b919050565b600080604083850312156118dd57600080fd5b6118e6836118a1565b91506118f4602084016118a1565b90509250929050565b6040516080810167ffffffffffffffff81118282101715611947577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290565b60405160e0810167ffffffffffffffff81118282101715611947577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b63ffffffff8116811461145f57600080fd5b80356118c581611997565b803560ff811681146118c557600080fd5b60008183036101408112156119d957600080fd5b6119e16118fd565b915060e08112156119f157600080fd5b506119fa61194d565b823581526020830135611a0c81611997565b60208201526040830135611a1f81611997565b60408201526060830135611a3281611997565b606082015260808301356bffffffffffffffffffffffff81168114611a5657600080fd5b6080820152611a6760a084016118a1565b60a082015260c083810135908201528152611a8460e083016119b4565b60208201526101008201356040820152611aa161012083016119a9565b606082015292915050565b6000806000806101a08587031215611ac357600080fd5b8435935060208501359250611adb86604087016119c5565b915061018085013567ffffffffffffffff811115611af857600080fd5b85016101008188031215611b0b57600080fd5b939692955090935050565b60008083601f840112611b2857600080fd5b50813567ffffffffffffffff811115611b4057600080fd5b602083019150836020828501011115611b5857600080fd5b9250929050565b60008060006101608486031215611b7557600080fd5b833567ffffffffffffffff811115611b8c57600080fd5b611b9886828701611b16565b9094509250611bac905085602086016119c5565b90509250925092565b600060208284031215611bc757600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b63ffffffff8481168252831660208201526060810160038310611c49577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b826040830152949350505050565b600060208284031215611c6957600080fd5b8135611c7481611997565b9392505050565b600080600080600060808688031215611c9357600080fd5b853567ffffffffffffffff811115611caa57600080fd5b611cb688828901611b16565b9096509450611cc99050602087016119b4565b92506040860135611cd981611997565b91506060860135611ce981611997565b809150509295509295909350565b600060208284031215611d0957600080fd5b611c74826118a1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082611d77577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60008219821115611d8f57611d8f611d12565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611dcc57611dcc611d12565b500290565b600082821015611de357611de3611d12565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611e4c57600080fd5b83018035915067ffffffffffffffff821115611e6757600080fd5b602001915036819003821315611b5857600080fd5b60008085851115611e8c57600080fd5b83861115611e9957600080fd5b5050820193919092039150565b60008451611eb8818460208901611820565b8201838582376000930192835250909392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611f0057611f00611d12565b5060010190565b600060208284031215611f1957600080fd5b5051919050565b8183823760009101908152919050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611f6557600080fd5b83018035915067ffffffffffffffff821115611f8057600080fd5b6020019150600581901b3603821315611b5857600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261201657600080fd5b830160208101925035905067ffffffffffffffff81111561203657600080fd5b803603831315611b5857600080fd5b604081833760408201600081526040808301823750600060808301525050565b60008383855260208086019550808560051b830101846000805b88811015612150577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868503018a5282357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee18936030181126120df578283fd5b8801803585526020808201359086015260408082013590860152606080820135908601526101206080612116818801828501612045565b5061010061212681840184611fe1565b9350828289015261213a8389018583611f98565b9d89019d9750505093860193505060010161207f565b509198975050505050505050565b60006101008083526121738184018b8d611f98565b9050602063ffffffff8a16818501528382036040850152818883528183019050818960051b8401018a60005b8b8110156121f7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030184526121d8828e611fe1565b6121e3858284611f98565b95870195945050509084019060010161219f565b5050858103606087015261220c81898b612065565b94505050505061221f6080830184612045565b9998505050505050505050565b60006020828403121561223e57600080fd5b81518015158114611c7457600080fd5b60208152600061148d602083018486611f98565b60006020828403121561227457600080fd5b8151611c7481611997565b60006101608083526122948184018688611f98565b915050825180516020840152602081015163ffffffff808216604086015280604084015116606086015280606084015116608086015250506bffffffffffffffffffffffff60808201511660a084015273ffffffffffffffffffffffffffffffffffffffff60a08201511660c084015260c081015160e084015250602083015161232461010084018260ff169052565b50604083015161012083015260609092015163ffffffff166101409091015292915050565b73ffffffffffffffffffffffffffffffffffffffff88811682528716602082015260ff8616604082015263ffffffff85811660608301528416608082015260c060a0820181905260009061221f9083018486611f9856fe2d5f2860204f2060295f2d202d5f2860206f2060295f2d202d5f286020512060295f2d2042495444414f204a5553542052454b5420594f55207c5f2860204f2060295f7c202d207c5f2860206f2060295f7c202d207c5f286020512060295f7ca2646970667358221220d0e91ac5de5b4dea4cc1c6702da8a58dc918f059deb94430aa8d28844976a5ab64736f6c63430008090033",
}

// BindingABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingMetaData.ABI instead.
var BindingABI = BindingMetaData.ABI

// BindingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BindingMetaData.Bin instead.
var BindingBin = BindingMetaData.Bin

// DeployBinding deploys a new Ethereum contract, binding an instance of Binding to it.
func DeployBinding(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Binding, error) {
	parsed, err := BindingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BindingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Binding{BindingCaller: BindingCaller{contract: contract}, BindingTransactor: BindingTransactor{contract: contract}, BindingFilterer: BindingFilterer{contract: contract}}, nil
}

// Binding is an auto generated Go binding around an Ethereum contract.
type Binding struct {
	BindingCaller     // Read-only binding to the contract
	BindingTransactor // Write-only binding to the contract
	BindingFilterer   // Log filterer for contract events
}

// BindingCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingSession struct {
	Contract     *Binding          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingCallerSession struct {
	Contract *BindingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BindingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingTransactorSession struct {
	Contract     *BindingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BindingRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingRaw struct {
	Contract *Binding // Generic contract binding to access the raw methods on
}

// BindingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingCallerRaw struct {
	Contract *BindingCaller // Generic read-only contract binding to access the raw methods on
}

// BindingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingTransactorRaw struct {
	Contract *BindingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBinding creates a new instance of Binding, bound to a specific deployed contract.
func NewBinding(address common.Address, backend bind.ContractBackend) (*Binding, error) {
	contract, err := bindBinding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Binding{BindingCaller: BindingCaller{contract: contract}, BindingTransactor: BindingTransactor{contract: contract}, BindingFilterer: BindingFilterer{contract: contract}}, nil
}

// NewBindingCaller creates a new read-only instance of Binding, bound to a specific deployed contract.
func NewBindingCaller(address common.Address, caller bind.ContractCaller) (*BindingCaller, error) {
	contract, err := bindBinding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingCaller{contract: contract}, nil
}

// NewBindingTransactor creates a new write-only instance of Binding, bound to a specific deployed contract.
func NewBindingTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingTransactor, error) {
	contract, err := bindBinding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingTransactor{contract: contract}, nil
}

// NewBindingFilterer creates a new log filterer instance of Binding, bound to a specific deployed contract.
func NewBindingFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingFilterer, error) {
	contract, err := bindBinding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingFilterer{contract: contract}, nil
}

// bindBinding binds a generic wrapper to an already deployed contract.
func bindBinding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BindingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Binding *BindingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Binding.Contract.BindingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Binding *BindingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Binding.Contract.BindingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Binding *BindingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Binding.Contract.BindingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Binding *BindingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Binding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Binding *BindingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Binding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Binding *BindingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Binding.Contract.contract.Transact(opts, method, params...)
}

// BLOCKSTALEMEASURE is a free data retrieval call binding the contract method 0x5e8b3f2d.
//
// Solidity: function BLOCK_STALE_MEASURE() view returns(uint256)
func (_Binding *BindingCaller) BLOCKSTALEMEASURE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "BLOCK_STALE_MEASURE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BLOCKSTALEMEASURE is a free data retrieval call binding the contract method 0x5e8b3f2d.
//
// Solidity: function BLOCK_STALE_MEASURE() view returns(uint256)
func (_Binding *BindingSession) BLOCKSTALEMEASURE() (*big.Int, error) {
	return _Binding.Contract.BLOCKSTALEMEASURE(&_Binding.CallOpts)
}

// BLOCKSTALEMEASURE is a free data retrieval call binding the contract method 0x5e8b3f2d.
//
// Solidity: function BLOCK_STALE_MEASURE() view returns(uint256)
func (_Binding *BindingCallerSession) BLOCKSTALEMEASURE() (*big.Int, error) {
	return _Binding.Contract.BLOCKSTALEMEASURE(&_Binding.CallOpts)
}

// FRAUDSTRING is a free data retrieval call binding the contract method 0x46b2eb9b.
//
// Solidity: function FRAUD_STRING() view returns(bytes)
func (_Binding *BindingCaller) FRAUDSTRING(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "FRAUD_STRING")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FRAUDSTRING is a free data retrieval call binding the contract method 0x46b2eb9b.
//
// Solidity: function FRAUD_STRING() view returns(bytes)
func (_Binding *BindingSession) FRAUDSTRING() ([]byte, error) {
	return _Binding.Contract.FRAUDSTRING(&_Binding.CallOpts)
}

// FRAUDSTRING is a free data retrieval call binding the contract method 0x46b2eb9b.
//
// Solidity: function FRAUD_STRING() view returns(bytes)
func (_Binding *BindingCallerSession) FRAUDSTRING() ([]byte, error) {
	return _Binding.Contract.FRAUDSTRING(&_Binding.CallOpts)
}

// DataManageAddress is a free data retrieval call binding the contract method 0xf2495029.
//
// Solidity: function dataManageAddress() view returns(address)
func (_Binding *BindingCaller) DataManageAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "dataManageAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataManageAddress is a free data retrieval call binding the contract method 0xf2495029.
//
// Solidity: function dataManageAddress() view returns(address)
func (_Binding *BindingSession) DataManageAddress() (common.Address, error) {
	return _Binding.Contract.DataManageAddress(&_Binding.CallOpts)
}

// DataManageAddress is a free data retrieval call binding the contract method 0xf2495029.
//
// Solidity: function dataManageAddress() view returns(address)
func (_Binding *BindingCallerSession) DataManageAddress() (common.Address, error) {
	return _Binding.Contract.DataManageAddress(&_Binding.CallOpts)
}

// DataStoreIdToRollupStoreNumber is a free data retrieval call binding the contract method 0xb537c4c7.
//
// Solidity: function dataStoreIdToRollupStoreNumber(uint32 ) view returns(uint256)
func (_Binding *BindingCaller) DataStoreIdToRollupStoreNumber(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "dataStoreIdToRollupStoreNumber", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DataStoreIdToRollupStoreNumber is a free data retrieval call binding the contract method 0xb537c4c7.
//
// Solidity: function dataStoreIdToRollupStoreNumber(uint32 ) view returns(uint256)
func (_Binding *BindingSession) DataStoreIdToRollupStoreNumber(arg0 uint32) (*big.Int, error) {
	return _Binding.Contract.DataStoreIdToRollupStoreNumber(&_Binding.CallOpts, arg0)
}

// DataStoreIdToRollupStoreNumber is a free data retrieval call binding the contract method 0xb537c4c7.
//
// Solidity: function dataStoreIdToRollupStoreNumber(uint32 ) view returns(uint256)
func (_Binding *BindingCallerSession) DataStoreIdToRollupStoreNumber(arg0 uint32) (*big.Int, error) {
	return _Binding.Contract.DataStoreIdToRollupStoreNumber(&_Binding.CallOpts, arg0)
}

// FraudProofPeriod is a free data retrieval call binding the contract method 0xf2a8f124.
//
// Solidity: function fraudProofPeriod() view returns(uint256)
func (_Binding *BindingCaller) FraudProofPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "fraudProofPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FraudProofPeriod is a free data retrieval call binding the contract method 0xf2a8f124.
//
// Solidity: function fraudProofPeriod() view returns(uint256)
func (_Binding *BindingSession) FraudProofPeriod() (*big.Int, error) {
	return _Binding.Contract.FraudProofPeriod(&_Binding.CallOpts)
}

// FraudProofPeriod is a free data retrieval call binding the contract method 0xf2a8f124.
//
// Solidity: function fraudProofPeriod() view returns(uint256)
func (_Binding *BindingCallerSession) FraudProofPeriod() (*big.Int, error) {
	return _Binding.Contract.FraudProofPeriod(&_Binding.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Binding *BindingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Binding *BindingSession) Owner() (common.Address, error) {
	return _Binding.Contract.Owner(&_Binding.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Binding *BindingCallerSession) Owner() (common.Address, error) {
	return _Binding.Contract.Owner(&_Binding.CallOpts)
}

// RollupStoreNumber is a free data retrieval call binding the contract method 0x428bba09.
//
// Solidity: function rollupStoreNumber() view returns(uint256)
func (_Binding *BindingCaller) RollupStoreNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "rollupStoreNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RollupStoreNumber is a free data retrieval call binding the contract method 0x428bba09.
//
// Solidity: function rollupStoreNumber() view returns(uint256)
func (_Binding *BindingSession) RollupStoreNumber() (*big.Int, error) {
	return _Binding.Contract.RollupStoreNumber(&_Binding.CallOpts)
}

// RollupStoreNumber is a free data retrieval call binding the contract method 0x428bba09.
//
// Solidity: function rollupStoreNumber() view returns(uint256)
func (_Binding *BindingCallerSession) RollupStoreNumber() (*big.Int, error) {
	return _Binding.Contract.RollupStoreNumber(&_Binding.CallOpts)
}

// RollupStores is a free data retrieval call binding the contract method 0xb0393a37.
//
// Solidity: function rollupStores(uint256 ) view returns(uint32 dataStoreId, uint32 confirmAt, uint8 status)
func (_Binding *BindingCaller) RollupStores(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DataStoreId uint32
	ConfirmAt   uint32
	Status      uint8
}, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "rollupStores", arg0)

	outstruct := new(struct {
		DataStoreId uint32
		ConfirmAt   uint32
		Status      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DataStoreId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.ConfirmAt = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// RollupStores is a free data retrieval call binding the contract method 0xb0393a37.
//
// Solidity: function rollupStores(uint256 ) view returns(uint32 dataStoreId, uint32 confirmAt, uint8 status)
func (_Binding *BindingSession) RollupStores(arg0 *big.Int) (struct {
	DataStoreId uint32
	ConfirmAt   uint32
	Status      uint8
}, error) {
	return _Binding.Contract.RollupStores(&_Binding.CallOpts, arg0)
}

// RollupStores is a free data retrieval call binding the contract method 0xb0393a37.
//
// Solidity: function rollupStores(uint256 ) view returns(uint32 dataStoreId, uint32 confirmAt, uint8 status)
func (_Binding *BindingCallerSession) RollupStores(arg0 *big.Int) (struct {
	DataStoreId uint32
	ConfirmAt   uint32
	Status      uint8
}, error) {
	return _Binding.Contract.RollupStores(&_Binding.CallOpts, arg0)
}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_Binding *BindingCaller) Sequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "sequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_Binding *BindingSession) Sequencer() (common.Address, error) {
	return _Binding.Contract.Sequencer(&_Binding.CallOpts)
}

// Sequencer is a free data retrieval call binding the contract method 0x5c1bba38.
//
// Solidity: function sequencer() view returns(address)
func (_Binding *BindingCallerSession) Sequencer() (common.Address, error) {
	return _Binding.Contract.Sequencer(&_Binding.CallOpts)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x7ef01d5e.
//
// Solidity: function confirmData(bytes data, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData) returns()
func (_Binding *BindingTransactor) ConfirmData(opts *bind.TransactOpts, data []byte, searchData IDataLayrServiceManagerDataStoreSearchData) (*types.Transaction, error) {
	return _Binding.contract.Transact(opts, "confirmData", data, searchData)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x7ef01d5e.
//
// Solidity: function confirmData(bytes data, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData) returns()
func (_Binding *BindingSession) ConfirmData(data []byte, searchData IDataLayrServiceManagerDataStoreSearchData) (*types.Transaction, error) {
	return _Binding.Contract.ConfirmData(&_Binding.TransactOpts, data, searchData)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x7ef01d5e.
//
// Solidity: function confirmData(bytes data, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData) returns()
func (_Binding *BindingTransactorSession) ConfirmData(data []byte, searchData IDataLayrServiceManagerDataStoreSearchData) (*types.Transaction, error) {
	return _Binding.Contract.ConfirmData(&_Binding.TransactOpts, data, searchData)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _sequencer, address _dataManageAddress) returns()
func (_Binding *BindingTransactor) Initialize(opts *bind.TransactOpts, _sequencer common.Address, _dataManageAddress common.Address) (*types.Transaction, error) {
	return _Binding.contract.Transact(opts, "initialize", _sequencer, _dataManageAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _sequencer, address _dataManageAddress) returns()
func (_Binding *BindingSession) Initialize(_sequencer common.Address, _dataManageAddress common.Address) (*types.Transaction, error) {
	return _Binding.Contract.Initialize(&_Binding.TransactOpts, _sequencer, _dataManageAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _sequencer, address _dataManageAddress) returns()
func (_Binding *BindingTransactorSession) Initialize(_sequencer common.Address, _dataManageAddress common.Address) (*types.Transaction, error) {
	return _Binding.Contract.Initialize(&_Binding.TransactOpts, _sequencer, _dataManageAddress)
}

// Parse is a paid mutator transaction binding the contract method 0x1f944c8f.
//
// Solidity: function parse(bytes[] polys, uint256 startIndex, uint256 length) returns(bytes)
func (_Binding *BindingTransactor) Parse(opts *bind.TransactOpts, polys [][]byte, startIndex *big.Int, length *big.Int) (*types.Transaction, error) {
	return _Binding.contract.Transact(opts, "parse", polys, startIndex, length)
}

// Parse is a paid mutator transaction binding the contract method 0x1f944c8f.
//
// Solidity: function parse(bytes[] polys, uint256 startIndex, uint256 length) returns(bytes)
func (_Binding *BindingSession) Parse(polys [][]byte, startIndex *big.Int, length *big.Int) (*types.Transaction, error) {
	return _Binding.Contract.Parse(&_Binding.TransactOpts, polys, startIndex, length)
}

// Parse is a paid mutator transaction binding the contract method 0x1f944c8f.
//
// Solidity: function parse(bytes[] polys, uint256 startIndex, uint256 length) returns(bytes)
func (_Binding *BindingTransactorSession) Parse(polys [][]byte, startIndex *big.Int, length *big.Int) (*types.Transaction, error) {
	return _Binding.Contract.Parse(&_Binding.TransactOpts, polys, startIndex, length)
}

// ProveFraud is a paid mutator transaction binding the contract method 0x7bd85879.
//
// Solidity: function proveFraud(uint256 fraudulentStoreNumber, uint256 startIndex, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, (bytes,uint32,bytes[],((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2]),bytes)[],(uint256[2],uint256[2])) disclosureProofs) returns()
func (_Binding *BindingTransactor) ProveFraud(opts *bind.TransactOpts, fraudulentStoreNumber *big.Int, startIndex *big.Int, searchData IDataLayrServiceManagerDataStoreSearchData, disclosureProofs BVM_EigenDataLayrChainDisclosureProofs) (*types.Transaction, error) {
	return _Binding.contract.Transact(opts, "proveFraud", fraudulentStoreNumber, startIndex, searchData, disclosureProofs)
}

// ProveFraud is a paid mutator transaction binding the contract method 0x7bd85879.
//
// Solidity: function proveFraud(uint256 fraudulentStoreNumber, uint256 startIndex, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, (bytes,uint32,bytes[],((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2]),bytes)[],(uint256[2],uint256[2])) disclosureProofs) returns()
func (_Binding *BindingSession) ProveFraud(fraudulentStoreNumber *big.Int, startIndex *big.Int, searchData IDataLayrServiceManagerDataStoreSearchData, disclosureProofs BVM_EigenDataLayrChainDisclosureProofs) (*types.Transaction, error) {
	return _Binding.Contract.ProveFraud(&_Binding.TransactOpts, fraudulentStoreNumber, startIndex, searchData, disclosureProofs)
}

// ProveFraud is a paid mutator transaction binding the contract method 0x7bd85879.
//
// Solidity: function proveFraud(uint256 fraudulentStoreNumber, uint256 startIndex, ((bytes32,uint32,uint32,uint32,uint96,address,bytes32),uint8,uint256,uint32) searchData, (bytes,uint32,bytes[],((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2]),bytes)[],(uint256[2],uint256[2])) disclosureProofs) returns()
func (_Binding *BindingTransactorSession) ProveFraud(fraudulentStoreNumber *big.Int, startIndex *big.Int, searchData IDataLayrServiceManagerDataStoreSearchData, disclosureProofs BVM_EigenDataLayrChainDisclosureProofs) (*types.Transaction, error) {
	return _Binding.Contract.ProveFraud(&_Binding.TransactOpts, fraudulentStoreNumber, startIndex, searchData, disclosureProofs)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Binding *BindingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Binding.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Binding *BindingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Binding.Contract.RenounceOwnership(&_Binding.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Binding *BindingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Binding.Contract.RenounceOwnership(&_Binding.TransactOpts)
}

// StoreData is a paid mutator transaction binding the contract method 0xbcb8ec90.
//
// Solidity: function storeData(bytes header, uint8 duration, uint32 blockNumber, uint32 totalOperatorsIndex) returns()
func (_Binding *BindingTransactor) StoreData(opts *bind.TransactOpts, header []byte, duration uint8, blockNumber uint32, totalOperatorsIndex uint32) (*types.Transaction, error) {
	return _Binding.contract.Transact(opts, "storeData", header, duration, blockNumber, totalOperatorsIndex)
}

// StoreData is a paid mutator transaction binding the contract method 0xbcb8ec90.
//
// Solidity: function storeData(bytes header, uint8 duration, uint32 blockNumber, uint32 totalOperatorsIndex) returns()
func (_Binding *BindingSession) StoreData(header []byte, duration uint8, blockNumber uint32, totalOperatorsIndex uint32) (*types.Transaction, error) {
	return _Binding.Contract.StoreData(&_Binding.TransactOpts, header, duration, blockNumber, totalOperatorsIndex)
}

// StoreData is a paid mutator transaction binding the contract method 0xbcb8ec90.
//
// Solidity: function storeData(bytes header, uint8 duration, uint32 blockNumber, uint32 totalOperatorsIndex) returns()
func (_Binding *BindingTransactorSession) StoreData(header []byte, duration uint8, blockNumber uint32, totalOperatorsIndex uint32) (*types.Transaction, error) {
	return _Binding.Contract.StoreData(&_Binding.TransactOpts, header, duration, blockNumber, totalOperatorsIndex)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Binding *BindingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Binding.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Binding *BindingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Binding.Contract.TransferOwnership(&_Binding.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Binding *BindingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Binding.Contract.TransferOwnership(&_Binding.TransactOpts, newOwner)
}

// BindingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Binding contract.
type BindingInitializedIterator struct {
	Event *BindingInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingInitialized represents a Initialized event raised by the Binding contract.
type BindingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Binding *BindingFilterer) FilterInitialized(opts *bind.FilterOpts) (*BindingInitializedIterator, error) {

	logs, sub, err := _Binding.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BindingInitializedIterator{contract: _Binding.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Binding *BindingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BindingInitialized) (event.Subscription, error) {

	logs, sub, err := _Binding.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingInitialized)
				if err := _Binding.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Binding *BindingFilterer) ParseInitialized(log types.Log) (*BindingInitialized, error) {
	event := new(BindingInitialized)
	if err := _Binding.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Binding contract.
type BindingOwnershipTransferredIterator struct {
	Event *BindingOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingOwnershipTransferred represents a OwnershipTransferred event raised by the Binding contract.
type BindingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Binding *BindingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Binding.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindingOwnershipTransferredIterator{contract: _Binding.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Binding *BindingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BindingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Binding.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingOwnershipTransferred)
				if err := _Binding.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Binding *BindingFilterer) ParseOwnershipTransferred(log types.Log) (*BindingOwnershipTransferred, error) {
	event := new(BindingOwnershipTransferred)
	if err := _Binding.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingRollupStoreConfirmedIterator is returned from FilterRollupStoreConfirmed and is used to iterate over the raw logs and unpacked data for RollupStoreConfirmed events raised by the Binding contract.
type BindingRollupStoreConfirmedIterator struct {
	Event *BindingRollupStoreConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingRollupStoreConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingRollupStoreConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingRollupStoreConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingRollupStoreConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingRollupStoreConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingRollupStoreConfirmed represents a RollupStoreConfirmed event raised by the Binding contract.
type BindingRollupStoreConfirmed struct {
	RollupStoreNumber uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupStoreConfirmed is a free log retrieval operation binding the contract event 0xe9f8c90baa6e73f4fbc1350dac0cf673eabb9d0bf5eef014ce5fe08be7d2d7a9.
//
// Solidity: event RollupStoreConfirmed(uint32 rollupStoreNumber)
func (_Binding *BindingFilterer) FilterRollupStoreConfirmed(opts *bind.FilterOpts) (*BindingRollupStoreConfirmedIterator, error) {

	logs, sub, err := _Binding.contract.FilterLogs(opts, "RollupStoreConfirmed")
	if err != nil {
		return nil, err
	}
	return &BindingRollupStoreConfirmedIterator{contract: _Binding.contract, event: "RollupStoreConfirmed", logs: logs, sub: sub}, nil
}

// WatchRollupStoreConfirmed is a free log subscription operation binding the contract event 0xe9f8c90baa6e73f4fbc1350dac0cf673eabb9d0bf5eef014ce5fe08be7d2d7a9.
//
// Solidity: event RollupStoreConfirmed(uint32 rollupStoreNumber)
func (_Binding *BindingFilterer) WatchRollupStoreConfirmed(opts *bind.WatchOpts, sink chan<- *BindingRollupStoreConfirmed) (event.Subscription, error) {

	logs, sub, err := _Binding.contract.WatchLogs(opts, "RollupStoreConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingRollupStoreConfirmed)
				if err := _Binding.contract.UnpackLog(event, "RollupStoreConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupStoreConfirmed is a log parse operation binding the contract event 0xe9f8c90baa6e73f4fbc1350dac0cf673eabb9d0bf5eef014ce5fe08be7d2d7a9.
//
// Solidity: event RollupStoreConfirmed(uint32 rollupStoreNumber)
func (_Binding *BindingFilterer) ParseRollupStoreConfirmed(log types.Log) (*BindingRollupStoreConfirmed, error) {
	event := new(BindingRollupStoreConfirmed)
	if err := _Binding.contract.UnpackLog(event, "RollupStoreConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingRollupStoreInitializedIterator is returned from FilterRollupStoreInitialized and is used to iterate over the raw logs and unpacked data for RollupStoreInitialized events raised by the Binding contract.
type BindingRollupStoreInitializedIterator struct {
	Event *BindingRollupStoreInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingRollupStoreInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingRollupStoreInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingRollupStoreInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingRollupStoreInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingRollupStoreInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingRollupStoreInitialized represents a RollupStoreInitialized event raised by the Binding contract.
type BindingRollupStoreInitialized struct {
	DataStoreId uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRollupStoreInitialized is a free log retrieval operation binding the contract event 0x957f0dd1f1ce8fbaa766e73503339f17b04cfbbd7e0db44e9460644485b813b5.
//
// Solidity: event RollupStoreInitialized(uint32 dataStoreId)
func (_Binding *BindingFilterer) FilterRollupStoreInitialized(opts *bind.FilterOpts) (*BindingRollupStoreInitializedIterator, error) {

	logs, sub, err := _Binding.contract.FilterLogs(opts, "RollupStoreInitialized")
	if err != nil {
		return nil, err
	}
	return &BindingRollupStoreInitializedIterator{contract: _Binding.contract, event: "RollupStoreInitialized", logs: logs, sub: sub}, nil
}

// WatchRollupStoreInitialized is a free log subscription operation binding the contract event 0x957f0dd1f1ce8fbaa766e73503339f17b04cfbbd7e0db44e9460644485b813b5.
//
// Solidity: event RollupStoreInitialized(uint32 dataStoreId)
func (_Binding *BindingFilterer) WatchRollupStoreInitialized(opts *bind.WatchOpts, sink chan<- *BindingRollupStoreInitialized) (event.Subscription, error) {

	logs, sub, err := _Binding.contract.WatchLogs(opts, "RollupStoreInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingRollupStoreInitialized)
				if err := _Binding.contract.UnpackLog(event, "RollupStoreInitialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupStoreInitialized is a log parse operation binding the contract event 0x957f0dd1f1ce8fbaa766e73503339f17b04cfbbd7e0db44e9460644485b813b5.
//
// Solidity: event RollupStoreInitialized(uint32 dataStoreId)
func (_Binding *BindingFilterer) ParseRollupStoreInitialized(log types.Log) (*BindingRollupStoreInitialized, error) {
	event := new(BindingRollupStoreInitialized)
	if err := _Binding.contract.UnpackLog(event, "RollupStoreInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingRollupStoreRevertedIterator is returned from FilterRollupStoreReverted and is used to iterate over the raw logs and unpacked data for RollupStoreReverted events raised by the Binding contract.
type BindingRollupStoreRevertedIterator struct {
	Event *BindingRollupStoreReverted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingRollupStoreRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingRollupStoreReverted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingRollupStoreReverted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingRollupStoreRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingRollupStoreRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingRollupStoreReverted represents a RollupStoreReverted event raised by the Binding contract.
type BindingRollupStoreReverted struct {
	RollupStoreNumber uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupStoreReverted is a free log retrieval operation binding the contract event 0x18407f85390b98bfd30ac355138feffbf9d2519036bddba3429ec57ed328e7a8.
//
// Solidity: event RollupStoreReverted(uint32 rollupStoreNumber)
func (_Binding *BindingFilterer) FilterRollupStoreReverted(opts *bind.FilterOpts) (*BindingRollupStoreRevertedIterator, error) {

	logs, sub, err := _Binding.contract.FilterLogs(opts, "RollupStoreReverted")
	if err != nil {
		return nil, err
	}
	return &BindingRollupStoreRevertedIterator{contract: _Binding.contract, event: "RollupStoreReverted", logs: logs, sub: sub}, nil
}

// WatchRollupStoreReverted is a free log subscription operation binding the contract event 0x18407f85390b98bfd30ac355138feffbf9d2519036bddba3429ec57ed328e7a8.
//
// Solidity: event RollupStoreReverted(uint32 rollupStoreNumber)
func (_Binding *BindingFilterer) WatchRollupStoreReverted(opts *bind.WatchOpts, sink chan<- *BindingRollupStoreReverted) (event.Subscription, error) {

	logs, sub, err := _Binding.contract.WatchLogs(opts, "RollupStoreReverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingRollupStoreReverted)
				if err := _Binding.contract.UnpackLog(event, "RollupStoreReverted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupStoreReverted is a log parse operation binding the contract event 0x18407f85390b98bfd30ac355138feffbf9d2519036bddba3429ec57ed328e7a8.
//
// Solidity: event RollupStoreReverted(uint32 rollupStoreNumber)
func (_Binding *BindingFilterer) ParseRollupStoreReverted(log types.Log) (*BindingRollupStoreReverted, error) {
	event := new(BindingRollupStoreReverted)
	if err := _Binding.contract.UnpackLog(event, "RollupStoreReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
