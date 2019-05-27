package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/liemei/crm-beego/crm/models"
)

type MainController struct {
	beego.Controller
}

func (self *MainController) Get() {
	v:=self.GetSession("User")
	u,ok:=v.(models.UserInfo)
	fmt.Println("登录用户：",u)
	if ok {
		self.Redirect("/ConsoleIndex/Index", 302)
		self.StopRun()
	}else {
		self.Redirect("/UserAccount/Login", 302)
		self.StopRun()
	}
}
