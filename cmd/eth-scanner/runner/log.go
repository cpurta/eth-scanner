package runner

import "log"

// TODO: create a csv writer that will output all of the address locations within
// a transaction
func (runner *EthereumTransactionScannerRunner) logTransactions() {
	defer runner.waitGroup.Done()
	for {
		transaction := <-runner.filteredTransactions
		log.Printf("Transaction found: %+v\n", transaction)
	}
}
