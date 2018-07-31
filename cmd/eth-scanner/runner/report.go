package runner

import (
	"log"
	"time"
)

func (runner *EthereumTransactionScannerRunner) reportProgress() {
	defer runner.waitGroup.Done()
	for !runner.done {
		retrieved, total := runner.blockWorkerManager.BlocksRetreived()
		log.Printf("Current block progress: %d / %d (%.4f%%)", retrieved, total, float64(retrieved)/float64(total)*100)
		time.Sleep(time.Minute * 1)
	}
}
