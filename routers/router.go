package routers

import (
	"github.com/astaxie/beego"
	"github.com/uxff/taniago/controllers"
)

func init() {
	beego.Router("/", &controllers.PicsetController{}, "get:Index")
	beego.Router("/picset/*", &controllers.PicsetController{}, "get:Picset")
	beego.Router("/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/signup", &controllers.LoginController{}, "get,post:Signup")
	beego.Router("/picset/*", &controllers.PicsetController{}, "get:Picset")
}
