package routers

import (
	"github.com/uxff/taniago/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	beego.Router("/index", &controllers.IndexController{}, "*:Index")
	beego.Router("/picset/*", &controllers.IndexController{}, "*:Picset")
	beego.Router("/main", &controllers.BaseController{}, "*:Index")
	beego.Router("/login", &controllers.BaseController{}, "*:Login")
	beego.Router("/logout", &controllers.BaseController{}, "*:Logout")
	beego.Router("/payment", &controllers.PaysapiController{}, "*:Payment")
	beego.Router("/notify", &controllers.PaysapiController{}, "*:Notify")
	beego.Router("/paymentstatus", &controllers.PaysapiController{}, "*:PaymentStatus")
}
