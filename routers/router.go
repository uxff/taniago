package routers

import (
	"github.com/astaxie/beego"
	"github.com/uxff/taniago/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/picset/*", &controllers.PicsetController{}, "get:Picset")
	beego.Router("/picset", &controllers.PicsetController{}, "delete:ClearCache")
	beego.Router("/user", &controllers.UsersController{}, "get,post:Index")
	beego.Router("/login", &controllers.UsersController{}, "get,post:Login")
	beego.Router("/logout", &controllers.UsersController{}, "get:Logout")
	beego.Router("/signup", &controllers.UsersController{}, "get,post:Signup")
	beego.Router("/picset/*", &controllers.PicsetController{}, "get:Picset")
}
