package routers

import (
	"github.com/uxff/taniago/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	beego.Router("/index", &controllers.IndexController{}, "*:Index")
	beego.Router("/picset/*", &controllers.IndexController{}, "*:Picset")
	beego.Router("/payment", &controllers.PaysapiController{}, "*:Payment")
	beego.Router("/notify", &controllers.PaysapiController{}, "*:Notify")
	beego.Router("/paymentstatus", &controllers.PaysapiController{}, "*:PaymentStatus")
	beego.Router("/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/signup", &controllers.LoginController{}, "get,post:Signup")
}
