package controllers

import "astaxie/beego"

type LoginControllers struct {
	beego.Controller
}

func (l LoginControllers)Get() {
	l.TplName="login.html"
}

func (l LoginControllers)Post() {
	l.TplName="login.html"
}
