package transaction

import (
	"strings"
	"sync"
	"time"
)

type TransactionWorker struct {
	incomingTransactions chan *TransactionResult
	filteredTransactions chan *TransactionResult
	filterAddress        string
	done                 bool
	waitGroup            *sync.WaitGroup
}

func NewTransactionWorker(incoming, outgoing chan *TransactionResult, filterAddress string, wg *sync.WaitGroup) *TransactionWorker {
	return &TransactionWorker{
		incomingTransactions: incoming,
		filteredTransactions: outgoing,
		filterAddress:        filterAddress,
		waitGroup:            wg,
		done:                 false,
	}
}

func (worker *TransactionWorker) Start() {
	defer worker.waitGroup.Done()
	for !worker.done {
		select {
		case transaction := <-worker.incomingTransactions:
			if worker.containsFilterAddress(transaction) {
				worker.filteredTransactions <- transaction
			}
		default:
			time.Sleep(time.Millisecond * 100)
		}

	}
}

func (worker *TransactionWorker) Stop() {
	worker.done = true
}

func (worker *TransactionWorker) containsFilterAddress(transaction *TransactionResult) bool {
	if transaction.To == worker.filterAddress || transaction.From == worker.filterAddress || strings.Contains(transaction.Input, worker.filterAddress[2:]) {
		return true
	}

	return false
}
