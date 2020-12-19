package main

import (
	"BeegoRpc/db_mysql"
	"BeegoRpc/models"
	_ "BeegoRpc/routers"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	db_mysql.ConnectDB();

	beego.SetStaticPath("/views","./views")
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")

	JsonData :=PraPareJson("getblockhash",1)
	ClienCount(JsonData)

	beego.Run()
}

func PraPareJson(method string,parmas ...interface{})string{
	Rpc :=models.RpcJson{
		Id:      time.Now().UnixNano(),
		Jsonrpc: "5.2.0",
		Method:  method,
		Params:  parmas,
	}
	data,err:=json.Marshal(&Rpc)
	if err !=nil{
		log.Fatal(err.Error())
	}
	return string(data)

}
const RPCUSER  ="user" //rpc服务的用户名
const RPCPASSWORD  ="pwd" //rpc服务的用户名
const RPCPORT  ="http://127.0.0.1:8332" //rpc服务的用户名

func ClienCount(Data string){
	client :=http.Client{};
	request, err :=http.NewRequest("POST",RPCPORT,strings.NewReader(Data))
	if err !=nil{
		log.Fatal(err.Error())
	}
	request.Header.Add("Encoding","UTF-8");
	request.Header.Add("Content-Type","application/json")
	request.SetBasicAuth(RPCUSER,RPCPASSWORD);
	resp,err :=client.Do(request);
	if err !=nil{
		log.Fatal(err.Error())
	}
	Code :=resp.StatusCode;
	if Code ==200{
		bytes,err :=ioutil.ReadAll(resp.Body)
		if err !=nil{
			log.Fatal(err.Error())
		}
		var f models.RPCResult
		err = json.Unmarshal(bytes,&f);
		if err !=nil{
			log.Fatal(err.Error())
		}
		fmt.Println(f.Result)
	}else {
		fmt.Println("请求失败")
	}
}


