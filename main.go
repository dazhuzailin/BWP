package main

import (
	_ "BWP/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("你好")
	beego.Run()
}

