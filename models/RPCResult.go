package entity

type RPCResult struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data Resultdata `json:"data"`
}
