package routers

import (
	"github.com/liemei/crm-beego/crm/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.AutoRouter(&controllers.UserAccountController{})
    beego.AutoRouter(&controllers.ConsoleIndexController{})
}
