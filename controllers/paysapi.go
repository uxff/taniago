package controllers

import (
	"github.com/astaxie/beego"
	"github.com/uxff/taniago/models/paysapi"
	"github.com/astaxie/beego/logs"
)

type PaysapiController struct {
	beego.Controller
}

/*
	initiat or create a payment
*/
func (this *PaysapiController) Payment() {
	this.TplName = "paysapi/payment.html"

	//this.GetString()

	paysapi.SetNotifyUrl("http://localhost"+beego.URLFor("PaysapiController.Notify"))
	paysapi.SetReturnUrl("http://localhost"+beego.URLFor("PaysapiController.PaymentStatus"))

	logs.Info("uid=%s token=%v", beego.AppConfig.DefaultString("paysapi_uid", ""), beego.AppConfig.DefaultString("paysapi_token", ""))
	paysapi.SetPaysapi(beego.AppConfig.DefaultString("paysapi_uid", ""), beego.AppConfig.DefaultString("paysapi_token", ""))

	pres, err := paysapi.SimplePayment("100321", 8, paysapi.ChanWeixin, "uid1001", "")
	if err != nil {
		logs.Error("go payment error:%v", err)
		return
	}

	//pres.Data.Qrcode
	this.Data["qrcode"] = pres.Data.Qrcode
	this.Data["msg"] = pres.Msg

	//this.Redirect()
}

/*
	recv notify from paysapi.com
	will call by paysapi.com
*/
func (this* PaysapiController) Notify() {

}

/*
	show a payment status
*/
func (this* PaysapiController) PaymentStatus() {

}
