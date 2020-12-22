package main

import (
	"beproject/Utils"
	"beproject/constants"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("hello,world")
    //beego.Run()
	bodyJson := Utils.RPCToJSON("getblockhash", 0)

	//2.发送一个post请求
	client := &http.Client{}
	request, err := http.NewRequest("POST", constants.RPCURL, strings.NewReader(string((bodyJson))))
	client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//请求头设置
	request.Header.Add("Encoding", "UTF-8")
	request.Header.Add("Content", "application/json")
	request.SetBasicAuth(constants.USER, constants.PASSWORLD)
	//autStr :=constants.USER+":"+constants.PASSWORLD
	//request.Header.Add("Authrization","Basic "+Utils.Base64(autStr))
	repose, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	code := repose.StatusCode
	if code == 200 {
		fmt.Println("请求成功")
	} else {
		fmt.Println("请求失败")
	}
	respByte,_ := ioutil.ReadAll(repose.Body)
	fmt.Println(string(respByte))
}
