package flags

import (
	"github.com/Layr-Labs/datalayr/common/logging"
	"github.com/urfave/cli"
)

const envVarPrefix = "CHALLENGER"

func prefixEnvVar(suffix string) string {
	return envVarPrefix + "_" + suffix
}

// TODO: Some of these flags should be integers?

var (
	/* Required Flags */

	RetrieverEndpointFlag = cli.StringFlag{
		Name:     "retriever",
		Usage:    "Endpoint at which retriever is available",
		Required: true,
		EnvVar:   prefixEnvVar("RETRIEVER"),
	}

	ChainIdFlag = cli.Uint64Flag{
		Name:     "chain-id",
		Usage:    "Chain id for ethereum chain",
		Required: true,
		EnvVar:   prefixEnvVar("CHAIN_ID"),
	}
	ChainProviderFlag = cli.StringFlag{
		Name:     "chain-provider",
		Usage:    "Ethereum chain rpc",
		Required: true,
		EnvVar:   prefixEnvVar("CHAIN_PROVIDER"),
	}
	GraphProviderFlag = cli.StringFlag{
		Name:     "graph-provider",
		Usage:    "Graphql endpoint for graph node",
		Required: true,
		EnvVar:   prefixEnvVar("GRAPH_PROVIDER"),
	}
	PrivateFlag = cli.StringFlag{
		Name:     "private",
		Usage:    "Ethereum private key for node operator",
		Required: true,
		EnvVar:   prefixEnvVar("PRIVATE"),
	}
	RollupAddressFlag = cli.StringFlag{
		Name:     "rollup-address",
		Usage:    "Address of the datalayr repository contract",
		Required: true,
		EnvVar:   prefixEnvVar("ROLLUP_ADDRESS"),
	}
	G1PathFlag = cli.StringFlag{
		Name:     "g1-path",
		Usage:    "Path to G1 SRS",
		Required: true,
		EnvVar:   prefixEnvVar("G1_PATH"),
	}
	G2PathFlag = cli.StringFlag{
		Name:     "g2-path",
		Usage:    "Path to G2 SRS",
		Required: true,
		EnvVar:   prefixEnvVar("G2_PATH"),
	}
	SrsTablePathFlag = cli.StringFlag{
		Name:     "srs-table-path",
		Usage:    "Path to SRS Table directory",
		Required: true,
		EnvVar:   prefixEnvVar("SRS_TABLE_PATH"),
	}
	OrderFlag = cli.StringFlag{
		Name:     "order",
		Usage:    "Order of the SRS",
		Required: true,
		EnvVar:   prefixEnvVar("ORDER"),
	}
	KzgWorkersFlag = cli.IntFlag{
		Name:     "kzg-num-workers",
		Usage:    "Order of the SRS",
		Required: false,
		Value:    4,
		EnvVar:   prefixEnvVar("KZG_NUM_WORKERS"),
	}
	StartStoreNumFlag = cli.Uint64Flag{
		Name:     "starting-store-numer",
		Usage:    "Store number from which challenger should pull",
		Required: false,
		Value:    4,
		EnvVar:   prefixEnvVar("STARTING_STORE_NUMER"),
	}
)

var requiredFlags = []cli.Flag{
	RetrieverEndpointFlag,
	ChainIdFlag,
	ChainProviderFlag,
	GraphProviderFlag,
	PrivateFlag,
	RollupAddressFlag,
	G1PathFlag,
	G2PathFlag,
	SrsTablePathFlag,
	OrderFlag,
}

var optionalFlags = []cli.Flag{
	KzgWorkersFlag,
}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
	Flags = append(Flags, logging.CLIFlags(envVarPrefix)...)
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag
