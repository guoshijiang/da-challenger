package challenger

import (
	"context"
	"github.com/Layr-Labs/datalayr/common/logging"
	"github.com/ethereum/go-ethereum/common"
	ethc "github.com/ethereum/go-ethereum/common"
	"github.com/mantlenetworkio/da-challenger/challenger"
	bsscore "github.com/mantlenetworkio/mantle/bss-core"
	"github.com/mantlenetworkio/mt-batcher/l1l2client"
	"github.com/urfave/cli"
	"time"
)

func Main(gitVersion string) func(ctx *cli.Context) error {
	return func(cliCtx *cli.Context) error {
		cfg, err := NewConfig(cliCtx)
		if err != nil {
			return err
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		logger, err := logging.GetLogger(logging.Config{})
		if err != nil {
			return err
		}
		sequencerPrivKey, _, err := bsscore.ParseWalletPrivKeyAndContractAddr(
			"DaSequencer", cfg.Mnemonic, cfg.SequencerHDPath,
			cfg.PrivateKey, cfg.EigenContractAddress,
		)
		if err != nil {
			return err
		}
		l1l2Config := &l1l2client.Config{
			L1RpcUrl:     cfg.L1EthRpc,
			ChainId:      cfg.ChainId,
			Private:      cfg.PrivateKey,
			DisableHTTP2: cfg.DisableHTTP2,
		}
		l1Client, err := l1l2client.NewL1ChainClient(ctx, l1l2Config)
		if err != nil {
			return err
		}
		l2Client, err := l1l2client.DialL2EthClientWithTimeout(ctx, cfg.L2MtlRpc, cfg.DisableHTTP2)
		if err != nil {
			return err
		}
		timeout, err := time.ParseDuration("12s")
		if err != nil {
			logger.Fatal().Err(err).Msg("Improper timeout from config")
		}
		challengerConfig := &challenger.ChallengerConfig{
			L1Client:          l1Client,
			L2Client:          l2Client,
			EigenContractAddr: ethc.Address(common.HexToAddress(cfg.EigenContractAddress)),
			Logger:            logger,
			PrivKey:           sequencerPrivKey,
			GraphProvider:     cfg.GraphProvider,
			RetrieverSocket:   cfg.RetrieverSocket,
			KzgConfig:         cfg.KzgConfig,
			LastStoreNumber:   cfg.FromStoreNumber,
			Timeout:           timeout,
		}
		clager, err := challenger.NewChallenger(ctx, challengerConfig)
		if err != nil {
			return err
		}
		if err := clager.Start(); err != nil {
			return err
		}
		defer clager.Stop()
		return nil

	}
}
