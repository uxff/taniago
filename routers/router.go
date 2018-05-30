package routers

import (
	"github.com/uxff/taniago/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	beego.Router("/index", &controllers.IndexController{}, "*:Index")
	beego.Router("/picset/*", &controllers.IndexController{}, "*:Picset")
	beego.Router("/main", &controllers.MainController{}, "*:Index")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
	beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/payment", &controllers.PaysapiController{}, "*:Payment")
	beego.Router("/notify", &controllers.PaysapiController{}, "*:Notify")
	beego.Router("/paymentstatus", &controllers.PaysapiController{}, "*:PaymentStatus")
}
