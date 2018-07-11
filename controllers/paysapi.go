package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/uxff/taniago/models/paysapi"
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
	appDomain := beego.AppConfig.String("appdomain")

	paysapi.SetNotifyUrl("http://"+appDomain+beego.URLFor("PaysapiController.Notify"))
	paysapi.SetReturnUrl("http://"+appDomain+beego.URLFor("PaysapiController.PaymentStatus"))

	logs.Info("uid=%s token=%v", beego.AppConfig.DefaultString("paysapi_uid", ""), beego.AppConfig.DefaultString("paysapi_token", ""))
	paysapi.SetPaysapi(beego.AppConfig.DefaultString("paysapi_uid", ""), beego.AppConfig.DefaultString("paysapi_token", ""))

	// alipay is notify success
	pres, err := paysapi.SimplePayment(fmt.Sprintf("test-%d", time.Now().Unix()), 8, paysapi.ChanWeixin, "uid1001", "")
	if err != nil {
		logs.Error("go payment error:%v", err)
		return
	}

	//pres.Data.Qrcode
	this.Data["qrcode"] = pres.Data.(map[string]interface{})["qrcode"]
	this.Data["msg"] = pres.Msg

	//this.Redirect()
}

/*
	recv notify from paysapi.com
	will call by paysapi.com
*/
func (this* PaysapiController) Notify() {
	//body, err := ioutil.ReadAll(this.Ctx.Request.Body)
	//if err != nil {
	//	logs.Error("read notify body error:%v body:%s", err, body)
	//}
	//logs.Info("i recved payment from request.body:%s", body)
	//
	//body = this.Ctx.Input.CopyBody(1<<20)
	//logs.Info("i recved payment from beego:%s input.data=%v", body, this.Ctx.Input.Data())

	logs.Info("i recved payment param: paysapi_id:%v", this.GetString("paysapi_id"))
	logs.Info("i recved payment param: key:%v", this.GetString("key"))
	logs.Info("i recved payment param: realprice:%v", this.GetString("realprice"))
	logs.Info("i recved payment param: orderid:%v", this.GetString("orderid"))

	logs.Info("i recved payment param: price:%v", this.GetString("price"))
	logs.Info("i recved payment param: orderuid:%v", this.GetString("orderuid"))


	//this.Data["data"] = body
	this.TplName = "paysapi/notify.html"
}

/*
	show a payment status
*/
func (this* PaysapiController) PaymentStatus() {
	this.TplName = "paysapi/paystatus.html"

	logs.Info("in return_url, orderid=%v", this.GetString("orderid"))

}
