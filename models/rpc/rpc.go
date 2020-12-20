package rpc

//norm:标椎、规范

//RPC通信标椎协议Request实体
type RPCNorm struct {
	Id      int64         `json:"id"`
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type RPCResult struct {
	Code int
	Msg  string
	Data *Result
}


//RPC通信标椎协议Response实体
type Result struct {
	Id     int64       `json:"id"`
	Error  error       `json:"error"`
	Result interface{} `json:"result"`
}
