package routers

import (
	"beproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/LoginPage",&controllers.ControllerLogin{})
}
