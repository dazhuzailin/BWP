package models

type Resultdata struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
	Id     int         `json:"id"`
}
