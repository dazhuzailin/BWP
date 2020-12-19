package models

type RpcJson struct {
	Id      int64  `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params []interface{} 	`json:"params"`
}
