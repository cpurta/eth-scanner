package transaction

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type TransactionReporter struct {
	outputFile           *os.File
	filteredTransactions chan *TransactionResult
	waitGroup            *sync.WaitGroup
	done                 bool
}

func NewTransactionReporter(outputFile *os.File, transactions chan *TransactionResult, wg *sync.WaitGroup) *TransactionReporter {
	return &TransactionReporter{
		outputFile:           outputFile,
		filteredTransactions: transactions,
		waitGroup:            wg,
	}
}

func (reporter *TransactionReporter) Start() {
	defer reporter.waitGroup.Done()

	for !reporter.done {
		select {
		case transaction := <-reporter.filteredTransactions:
			transactionRow := fmt.Sprintf("%s,%s,%s,%s\n", transaction.BlockHash, transaction.To, transaction.From, transaction.Input)
			if _, err := reporter.outputFile.WriteString(transactionRow); err != nil {
				log.Println("Error writing transaction to file:", err.Error())
				reporter.done = true
			}
		default:
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func (reporter *TransactionReporter) Stop() {
	reporter.done = true
}
