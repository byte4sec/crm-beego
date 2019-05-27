package main

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/liemei/crm-beego/crm/models"
	_ "github.com/liemei/crm-beego/crm/routers"
)
func init()  {
	sqlconn:=beego.AppConfig.String("sqlconn")
	mysqlclientMax, _ :=beego.AppConfig.Int("mysqlclientMax")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", sqlconn, mysqlclientMax)
	orm.RunSyncdb("default",false,true)
	beego.BConfig.Listen.EnableAdmin = true
	beego.BConfig.Listen.AdminAddr = "localhost"
	beego.BConfig.Listen.AdminPort = 8088
	//redisconn:=beego.AppConfig.String("redis.host")
	beego.BConfig.WebConfig.Session.SessionProvider="redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
	gob.Register(models.UserInfo{})
}
func main() {
	beego.Run()
}

