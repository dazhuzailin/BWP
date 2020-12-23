package controllers

import (
	"BWP/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LonginControllers struct {
	beego.Controller
}

func (l *LonginControllers) Get() {
	l.TplName = "LoginPage.html"
}

func (l *LonginControllers) Post() {
	var user models.User

	err := l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("1")
		return
	}
	fmt.Println(user.Name,user.Password)
	l.Ctx.WriteString("2")
	 _ , err =user.Queryuser()
	if err != nil {
		l.Ctx.WriteString("cw")
	}
	l.Ctx.WriteString("dengluchenggong")
}