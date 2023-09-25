package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/uxff/taniago/lib/tiksaver"
)

type TiksaverController struct {
	BaseController
}

// tiktok saver index
func (this *TiksaverController) Index() {

	this.Data["appname"] = beego.AppConfig.String("appname")
	this.Data["outputfile"] = ""
	this.Data["link"] = ""
	this.Data["errmsg"] = ""
	this.TplName = "tiksaver/index.tpl"

}
func (this *TiksaverController) Download() {

	this.Data["appname"] = beego.AppConfig.String("appname")
	link := this.GetString("link")
	// is := this.GetString("link")
	if link == "" {
		logs.Warn("no link param")
		this.Data["errmsg"] = "no link param"
		this.Redirect(this.URLFor("TiksaverController.Index"), 303)
		return
	}

	outputfile, err := tiksaver.DownloadTiktokTikwm(link)

	this.Data["outputfile"] = outputfile
	this.Data["link"] = link
	this.Data["errmsg"] = ""
	if err != nil {
		this.Data["errmsg"] = fmt.Sprintf("%+v", err)
		logs.Warn("download %v error:%+v", link, err)
	}

	logs.Debug("download %v to %v successfully", link, outputfile)

	// this.Redirect(this.URLFor("TiksaverController.Index"), 303)
	this.TplName = "tiksaver/index.tpl"
}
