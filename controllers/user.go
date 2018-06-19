/*
    user controller
    todo: captcha of login, register
*/
package controllers

import (
	usermodel "github.com/uxff/taniago/models/user"
	"github.com/astaxie/beego"
)

type UserController struct {
	BaseController
}

func (c *UserController) NestPrepare() {
	if !c.IsLogin {
		c.Ctx.Redirect(302, c.LoginPath())
		return
	}
}

// func (c *UsersController) NestFinish() {}

func (c *UserController) Index() {
	beego.ReadFromRequest(&c.Controller)

	c.TplName = "users/index.tpl"
}


/*
    login action
    todo: captcha of login, register
*/
func (this *UserController) Login() {
	userinfo := this.GetSession("userinfo")
	// fmt.Println(userinfo)
	if userinfo != nil {
		this.Ctx.Redirect(302, "/index")
	}

	this.TplName = "login.html"

	email := this.GetString("email")
	if email != "" {
		password := this.GetString("password")
        nickname := this.GetString("nickname")
		ue := new(usermodel.UserEntity)
		ue.Email = email
		ue.Nickname = nickname
		ue.Password = password
		this.SetSession("userinfo", ue)
		this.Ctx.Redirect(302, "/index")
		return
	}
	return
}

/*
    logout action
*/
func (this *UserController) Logout() {
	this.DelSession("userinfo")
	this.Ctx.Redirect(302, "/login")
}

/*
    register action
*/



