package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/liemei/crm-beego/crm/models"
	"strings"
)

type BaseController struct {
	beego.Controller
	user *models.UserInfo
}

func (self *BaseController) Prepare()  {
	self.auth()
}
// 权限校验
func (self *BaseController) auth(){
	self.user = &models.UserInfo{}
	v:=self.GetSession("User")
	fmt.Println("当前用户信息：",v)
	u,ok:=v.(models.UserInfo)
	if ok {
		self.user = &u
	}else {
		self.redirect("/UserAccount/Login")
	}
}
// 是否POST提交
func (self *BaseController) isPost() bool {
	return self.Ctx.Request.Method == "POST"
}

//获取用户IP地址
func (self *BaseController) getClientIp() string {
	s := self.Ctx.Request.RemoteAddr
	l := strings.LastIndex(s, ":")
	return s[0:l]
}

// 重定向
func (self *BaseController) redirect(url string) {
	self.Redirect(url, 302)
	self.StopRun()
}

func (this *BaseController) SetJson(code int, data interface{}, Msg string) {
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


