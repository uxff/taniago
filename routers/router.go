package routers

import (

	"github.com/astaxie/beego"

	"github.com/uxff/taniago/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	beego.Router("/index", &controllers.IndexController{}, "*:Index")
	beego.Router("/picset/*", &controllers.PicsetController{}, "*:Picset")
	beego.Router("/clearcache/*", &controllers.PicsetController{}, "*:ClearCache")
	beego.Router("/payment", &controllers.PaysapiController{}, "*:Payment")
	beego.Router("/notify", &controllers.PaysapiController{}, "*:Notify")
	beego.Router("/paymentstatus", &controllers.PaysapiController{}, "*:PaymentStatus")
	beego.Router("/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/signup", &controllers.LoginController{}, "get,post:Signup")
}
