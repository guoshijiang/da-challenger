// Package main implements a grpc service that can be used with any service.
//
// Definitions can be found under the proto folders for each service.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Layr-Labs/datalayr/common/graphView"
	"github.com/Layr-Labs/datalayr/common/logging"
	"github.com/Layr-Labs/datalayr/middleware/rollup-example/utils/contracts"
	"github.com/mantlenetworkio/da-challenger"
	"github.com/mantlenetworkio/da-challenger/flags"

	"github.com/urfave/cli"
)

var (
	Version   = ""
	GitCommit = ""
	GitDate   = ""
)

func main() {

	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Version = fmt.Sprintf("%s-%s-%s", Version, GitCommit, GitDate)
	app.Name = "sequencer"
	app.Usage = "DataLayr Rollup Sequencer"
	app.Description = "Models interaction between sequencer and Datalayr"

	app.Action = runner
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed", "message", err)
	}
}

func runner(ctx *cli.Context) error {

	settings := challenger.NewSettings(ctx)

	logger, err := logging.GetLogger(settings.LoggingConfig)
	if err != nil {
		return err
	}

	// Make ChainClient and GraphClient
	chainClient := contracts.MakeChainClient(settings.ChainSettings, logger)
	graphClient := graphView.NewGraphClient(settings.GraphEndpoint, logger)

	// Make Sequencer
	c := challenger.NewChallenger(chainClient, graphClient, logger, settings, settings.FromStoreNumber)

	c.FraudCheckLoop()

	select {}
}
