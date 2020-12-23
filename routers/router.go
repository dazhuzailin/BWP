package routers

import (
	"BWP/controllers"
	"github.com/astaxie/beego"
)
func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login",&controllers.LonginControllers{})
}
