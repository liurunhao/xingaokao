package routers

import (
	"xingaokao/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{},"get:ShowLogin;post:HandleLogin")
	beego.Router("/register", &controllers.MainController{},"get:ShowRegister;post:HandleRegister")
	beego.Router("/logout", &controllers.MainController{},"post:Logout")
	beego.Router("/gw", &controllers.MainController{}, "GET:GuoWu")
	beego.Router("/jx", &controllers.MainController{}, "GET:JiangXi")
	beego.Router("/sh", &controllers.MainController{}, "GET:ShangHai")
	beego.Router("/zj", &controllers.MainController{}, "GET:ZheJiang")
	beego.Router("/qw", &controllers.MainController{}, "GET:QuWei")
	beego.Router("/location/:id([\\d]+)/colleges", &controllers.MainController{}, "GET:GetUniversities")
	beego.Router("/colleges/:id([\\d]+)/professionals", &controllers.MainController{}, "GET:GetProfessionls")
	beego.Router("/test", &controllers.MainController{}, "POST:Test")
}
