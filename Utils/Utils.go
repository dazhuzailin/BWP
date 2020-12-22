package Utils

import (
	"beproject/entity"
	"encoding/json"
	"log"
	"time"
)

func RPCToJSON(method string, params ...interface{}) string {
	/*
	  1.准备一个json格式的数据（rpc通信的协议规范）
	  2.发送一个post请求，发送http链接到rpc服务节点
	*/
	//1.准备一个json数据（string）

	//获取节点的区块的总数
	RPCrequest := entity.RPCRequest{
		Id:      time.Now().Unix(),
		Method:  method,
		Jsonrpc: "2.0",
		Params:  params,
	}
	reqBytes, err := json.Marshal(&RPCrequest)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(reqBytes)
}
