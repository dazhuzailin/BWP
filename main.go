package main

import (
	_ "BWP/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("Hello World")
	beego.Run()
}

