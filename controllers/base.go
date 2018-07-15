package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/ikeikeikeike/gopkg/convert"
	"github.com/uxff/beego-samples/auth/models"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

// 初始化captcha
var TheStore cache.Cache
var TheCaptcha *captcha.Captcha

func init() {
	TheStore = cache.NewMemoryCache()
	TheCaptcha = captcha.NewWithFilter("/captcha/", TheStore) //一定要写在构造函数里面，要不然第一次打开页面有可能是X
	TheCaptcha.StdHeight = 40
	TheCaptcha.StdWidth = 100
	TheCaptcha.ChallengeNums = 4
}

type BaseController struct {
	beego.Controller

	Userinfo   *models.User
	IsLogin    bool

	theStore   cache.Cache
	theCaptcha *captcha.Captcha
}

type NestPreparer interface {
	NestPrepare()
}

type NestFinisher interface {
	NestFinish()
}

// every request will call this
func (c *BaseController) Prepare() {
	c.SetParams()

	c.IsLogin = c.GetSession("userinfo") != nil
	if c.IsLogin {
		c.Userinfo = c.GetLogin()
	}

	c.Data["appname"] = beego.AppConfig.String("appname")

	c.Data["IsLogin"] = c.IsLogin
	c.Data["Userinfo"] = c.Userinfo

	c.Data["HeadStyles"] = []string{}
	c.Data["HeadScripts"] = []string{}

	c.Layout = "base.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["BaseHeader"] = "header.tpl"
	c.LayoutSections["BaseFooter"] = "footer.tpl"

	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (c *BaseController) Finish() {
	if app, ok := c.AppController.(NestFinisher); ok {
		app.NestFinish()
	}
}

func (c *BaseController) GetLogin() *models.User {
	u := &models.User{Id: c.GetSession("userinfo").(int64)}
	u.Read()
	return u
}

func (c *BaseController) DelLogin() {
	c.DelSession("userinfo")
}

func (c *BaseController) SetLogin(user *models.User) {
	c.SetSession("userinfo", user.Id)
}

func (c *BaseController) LoginPath() string {
	return c.URLFor("LoginController.Login")
}

func (c *BaseController) SetParams() {
	c.Data["Params"] = make(map[string]string)
	for k, v := range c.Input() {
		c.Data["Params"].(map[string]string)[k] = v[0]
	}
}

func (c *BaseController) BuildRequestUrl(uri string) string {
	if uri == "" {
		uri = c.Ctx.Input.URI()
	}
	return fmt.Sprintf("%s:%s%s",
		c.Ctx.Input.Site(), convert.ToStr(c.Ctx.Input.Port()), uri)
}
