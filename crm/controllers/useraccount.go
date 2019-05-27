package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/liemei/crm-beego/crm/models"
)

type UserAccountController struct {
	beego.Controller
}

func (self *UserAccountController) Login()  {
	if self.isPost() {
		username:=self.GetString("username")
		password:=self.GetString("password")
		var user = models.UserLogin(username,password)
		fmt.Println("登录用户：",user)
		if user.Id>0 {
			self.SetSession("User",user)
			self.setJson(1,user,"登录成功")
		}else {
			self.setJson(0,user,"登录失败")
		}
	}else {
		self.TplName="login.tpl"
	}
}


func (self *UserAccountController) isPost() bool {
	return self.Ctx.Request.Method == "POST"
}
func (this *UserAccountController) setJson(code int, data interface{}, Msg string) {
	if code == 0 {
		if Msg == "" {
			Msg = "sucess"
		}
		this.Data["json"] = map[string]interface{}{"code": code, "msg": Msg, "data": data}
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"code": code, "msg": Msg, "data": data}
		this.ServeJSON()
	}

}