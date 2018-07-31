package rpc

type RPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int64         `json:"id"`
}

func NewRPCRequest(method string, params []interface{}) *RPCRequest {
	return &RPCRequest{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      int64(1),
	}
}
