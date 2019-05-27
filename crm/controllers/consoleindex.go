package controllers

import (
	"fmt"
)

type ConsoleIndexController struct {
	BaseController
}

func (self *ConsoleIndexController) Index()  {
	fmt.Println("首页")
	self.Data["name"] = self.user.Name
	self.TplName = "index.tpl"
	self.Layout = "shard/Layout.tpl"
	self.LayoutSections = make(map[string]string)
	self.LayoutSections["header"] = "shard/header.html"
	self.LayoutSections["Sidebar"] = "shard/footer.html"
}