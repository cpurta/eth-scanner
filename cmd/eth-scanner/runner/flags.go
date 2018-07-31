package runner

import "github.com/urfave/cli"

func (runner *EthereumTransactionScannerRunner) getStartFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "ethereum-host",
			Value:       "http://localhost:8545",
			EnvVar:      "ETHEREUM_HOST",
			Usage:       "host endpoint of the ethereum node",
			Destination: &runner.endpoint,
		},
		cli.IntFlag{
			Name:        "block-workers",
			Value:       1,
			EnvVar:      "BLOCK_WORKERS",
			Usage:       "number of routines that will pull block information",
			Destination: &runner.blockWorkerNum,
		},
		cli.Int64Flag{
			Name:        "start-block",
			Value:       1,
			EnvVar:      "START_BLOCK",
			Usage:       "block number of the first block to begin scanning",
			Destination: &runner.startBlock,
		},
		cli.Int64Flag{
			Name:        "end-block",
			Value:       6000000,
			EnvVar:      "END_BLOCK",
			Usage:       "block number of the first block to begin scanning",
			Destination: &runner.endBlock,
		},
		cli.StringFlag{
			Name:        "filter-address",
			EnvVar:      "FILTER_ADDRESS",
			Usage:       "filters all transactions to those only containing the specified address",
			Destination: &runner.filterAddress,
		},
	}
}
