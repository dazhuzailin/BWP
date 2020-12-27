package main

import (
	"BWP/dbMysql"
	_ "BWP/routers"
	"github.com/astaxie/beego"
)

func main() {
	dbMysql.ConnectDB()
	beego.Run()
}

