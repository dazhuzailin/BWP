package main

import (
	"BWP/dbMysql"
	_ "BWP/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("Hello World")
	dbMysql.ConnectDB()
	beego.Run()
}

