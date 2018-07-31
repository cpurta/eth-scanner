package block

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/cpurta/eth-scanner/cmd/internal/rpc"
	"github.com/cpurta/eth-scanner/cmd/internal/transaction"
)

type BlockWorker struct {
	endpoint           string
	transactionChannel chan *transaction.TransactionResult
	blockRange         *BlockRange
	waitGroup          *sync.WaitGroup
	completed          int64
	done               bool
}

func NewBlockWorker(endpoint string, transactions chan *transaction.TransactionResult, wg *sync.WaitGroup) *BlockWorker {
	return &BlockWorker{
		endpoint:           endpoint,
		transactionChannel: transactions,
		blockRange:         NewBlockRange(0, 0),
		waitGroup:          wg,
		completed:          0,
		done:               false,
	}
}

func (worker *BlockWorker) Start() error {
	worker.waitGroup.Add(1)
	defer worker.waitGroup.Done()

	log.Println("Block worker starting to pull block info for blocks", worker.blockRange.Min(), "-", worker.blockRange.Max())
	for blockNumber := worker.blockRange.Min(); blockNumber < worker.blockRange.Max(); blockNumber++ {
		if worker.done {
			break
		}

		hexBlockNumber := fmt.Sprintf("0x%x", blockNumber)

		blockResult, err := worker.getBlock(hexBlockNumber)
		if err != nil {
			log.Println("Unable to retrieve block results for block", blockNumber, ":", err.Error())
			continue
		}

		worker.completed++
		worker.pushTransactions(blockResult.Result.Transactions)
	}

	return nil
}

func (worker *BlockWorker) Stop() {
	worker.done = true
}

func (worker *BlockWorker) SetBlockRange(blockRange *BlockRange) {
	worker.blockRange = blockRange
}

func (worker *BlockWorker) GetProgress() (int64, int64) {
	return worker.completed, worker.blockRange.Len()
}

func (worker *BlockWorker) getBlock(blockID string) (*RPCBlockResult, error) {
	var (
		rpcRequest *rpc.RPCRequest
		body       []byte
		err        error
	)

	rpcRequest = rpc.NewRPCRequest("eth_getBlockByNumber", []interface{}{blockID, true})

	rpcPayload, err := json.Marshal(rpcRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", worker.endpoint, strings.NewReader(string(rpcPayload)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error making request to rpc node")
		return nil, err
	}

	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body")
		return nil, err
	}

	rpcBlockRequest := &RPCBlockResult{}
	err = json.Unmarshal(body, rpcBlockRequest)

	return rpcBlockRequest, err
}

func (worker *BlockWorker) pushTransactions(transactions []*transaction.TransactionResult) {
	if worker.done {
		return
	}

	for _, txn := range transactions {
		worker.transactionChannel <- txn
	}
}
