package main

import (
	"log"
	"os"

	"github.com/cpurta/eth-scanner/cmd/eth-scanner/runner"
	"github.com/urfave/cli"
)

var ()

func main() {
	sigKillChan := make(chan os.Signal, 1)

	app := cli.App{
		Name:    "eth-scanner",
		Usage:   "Scan blocks on the Ethereum block chain and filter transactions",
		Version: "0.0.1",
		Commands: []cli.Command{
			runner.NewCommand(sigKillChan),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("unexpected error: ", err)
	}

}
