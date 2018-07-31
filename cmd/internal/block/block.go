package block

import "github.com/cpurta/eth-scanner/cmd/internal/transaction"

type RPCBlockResult struct {
	ID      int         `json:"id"`
	JSONRPC string      `json:"jsonrpc"`
	Result  BlockResult `json:"result"`
}

type BlockResult struct {
	Number           string                           `json:"number"`
	Hash             string                           `json:"hash"`
	ParentHash       string                           `json:"parentHash"`
	MixHash          string                           `json:"mixHash"`
	Nonce            string                           `json:"nonce"`
	Sha3Uncles       string                           `json:"sha3Uncles"`
	LogsBloom        string                           `json:"logsBloom"`
	TransactionsRoot string                           `json:"transactionsRoot"`
	StateRoot        string                           `json:"stateRoot"`
	Miner            string                           `json:"miner"`
	Difficulty       string                           `json:"difficulty"`
	TotalDifficulty  string                           `json:"totalDifficulty"`
	ExtraData        string                           `json:"extraData"`
	Size             string                           `json:"size"`
	GasLimit         string                           `json:"gasLimit"`
	GasUsed          string                           `json:"gasUsed"`
	Timestamp        string                           `json:"timestamp"`
	Transactions     []*transaction.TransactionResult `json:"transactions"`
	Uncles           []string                         `json:"uncles"`
}

type BlockRange struct {
	min int64
	max int64
}

func NewBlockRange(min int64, max int64) *BlockRange {
	return &BlockRange{
		min: min,
		max: max,
	}
}

func (blockrange *BlockRange) Min() int64 {
	return blockrange.min
}

func (blockrange *BlockRange) Max() int64 {
	return blockrange.max
}

func (blockRange *BlockRange) Len() int64 {
	if blockRange.max == 0 {
		return 0
	}

	return blockRange.max - blockRange.min
}
