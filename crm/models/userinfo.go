package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/liemei/crm-beego/crm/utils"
	"time"
)

/**角色*/
type Role int
const (
	// 销售业务员
	Role_Sale Role = iota
	// 销售主管
	Role_SaleSuper
	// 系统管理员
	Role_SystemAdmin
)
/**用户基本信息*/
type UserInfo struct{
	Id int
	Name string `orm:"size(100)"`
	Role int
	HeadImage string `orm:"size(256)"`
	CreateTime time.Time
}
/**用户账户信息*/
type UserAccount struct {
	Id int
	UserName string `orm:"size(32)"`
	Password string `orm:"size(126)"`
	Userid int
}
func init(){
	orm.RegisterModel(new(UserInfo),new(UserAccount))
}


func RegisterUserInfo(ua *UserAccount,u *UserInfo) int{
	ua.Password = utils.Md5String(ua.Password+"&liushuaichaocrm!@#")
	o:=orm.NewOrm()
	err:=o.Begin()
	uid,err_u:=o.Insert(&u)
	ua.Userid = int(uid)
	acid,err_ac:=o.Insert(&ua)
	if err_u!=nil || err_ac!=nil {
		err = o.Rollback()
	}else {
		err = o.Commit()
	}
	if err == nil {
		//fmt.Println("RegisterUserInfo",err)
	}
	return int(acid)
}

func UserLogin(username string ,password string) UserInfo{
	password = utils.Md5String(password+"&liushuaichaocrm!@#")
	fmt.Println("密码：",password)
	var u UserInfo
	o:=orm.NewOrm()
	uc:= UserAccount{UserName:username,Password:password}
	err:=o.Read(&uc,"UserName","Password")
	fmt.Println(uc,"---useraccount")
	if err == orm.ErrNoRows {
		return u
	}
	u = UserInfo{Id:uc.Userid}
	err=o.Read(&u,"Id")
	fmt.Println(u,"----user")
	return u
}

