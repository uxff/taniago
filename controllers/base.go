package controllers

import (
	"fmt"
	"github.com/uxff/taniago/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Index() {

	this.TplName = "index.html"
	return


}

func (this *MainController) Login() {
	userinfo := this.GetSession("userinfo")
	// fmt.Println(userinfo)
	if userinfo != nil {
		this.Ctx.Redirect(302, "/index")
	}
	this.TplName = "login.html"
	spacename := this.GetString("spacename")
	username := this.GetString("username")
	fmt.Println(username)
	if username != "" && spacename != "" {
		password := this.GetString("password")
		user := new(models.Space)
		user.Name = spacename
		user.UserName = username
		user.PassWord = password
		this.SetSession("userinfo", user)
		this.Ctx.Redirect(302, "/index")
		return
	}
	return
}
func (this *MainController) Logout() {
	this.DelSession("userinfo")
	this.Ctx.Redirect(302, "/login")
}

func (this *MainController) Captcha() {
	
}

