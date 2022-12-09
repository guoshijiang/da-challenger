package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	da_challenger "github.com/mantlenetworkio/da-challenger"
	"github.com/mantlenetworkio/da-challenger/flags"
	"github.com/urfave/cli"
	"os"
)

var (
	GitVersion = ""
	GitCommit  = ""
	GitDate    = ""
)

func main() {
	app := cli.NewApp()
	app.Flags = flags.Flags
	app.Version = fmt.Sprintf("%s-%s", GitVersion, params.VersionWithCommit(GitCommit, GitDate))
	app.Name = "da-challenger"
	app.Usage = "EigenDA Challenger Service"
	app.Description = "Service for eigneDa challenger " +
		"check eigenDa data store right or wrong"
	app.Action = da_challenger.Main(GitVersion)
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Application failed", "message", err)
	}
}
