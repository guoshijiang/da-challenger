package challenger

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"

	"github.com/Layr-Labs/datalayr/common/logging"
	"github.com/Layr-Labs/datalayr/middleware/rollup-example/challenger/flags"
	"github.com/Layr-Labs/datalayr/middleware/rollup-example/utils/contracts"
)

func NewSettings(ctx *cli.Context) ChallengerSettings {
	return ChallengerSettings{
		ChainSettings: contracts.ChainSettings{
			Provider: ctx.GlobalString(flags.ChainProviderFlag.Name),
			ChainId:  ctx.GlobalUint64(flags.ChainIdFlag.Name),
			Private:  ctx.GlobalString(flags.PrivateFlag.Name),
		},
		RollupSettings: RollupSettings{
			Address: common.HexToAddress(ctx.GlobalString(flags.RollupAddressFlag.Name)),
		},
		RetrieverSettings: RetrieverSettings{
			Socket: ctx.GlobalString(flags.RetrieverEndpointFlag.Name),
		},
		KzgConfig: KzgConfig{
			G1Path:    ctx.GlobalString(flags.G1PathFlag.Name),
			G2Path:    ctx.GlobalString(flags.G2PathFlag.Name),
			TableDir:  ctx.GlobalString(flags.SrsTablePathFlag.Name),
			Order:     ctx.GlobalUint64(flags.OrderFlag.Name),
			NumWorker: ctx.GlobalInt(flags.KzgWorkersFlag.Name),
		},
		LoggingConfig:   logging.ReadCLIConfig(ctx),
		GraphEndpoint:   ctx.GlobalString(flags.GraphProviderFlag.Name),
		FromStoreNumber: ctx.GlobalUint64(flags.StartStoreNumFlag.Name),
	}
}
