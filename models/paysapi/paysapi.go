/*
	参考：https://www.paysapi.com/docpay
	流程：
		1、调用发起付款接口，获得支付二维码
		2、用户端支付
		3、paysapi服务器回调通知给商户
		4、商户可以查询订单

*/
package paysapi

import (
	"crypto/md5"
	"fmt"
	"encoding/json"
	"encoding/hex"

	"github.com/uxff/taniago/utils"
	"github.com/astaxie/beego/logs"
)


const (
	ApiPaymentJson = "https://pay.bbbapi.com/?format=json"
	ApiPaymentPage = "https://pay.bbbapi.com/"
	ApiQueryPayment = "https://api.bbbapi.com/get_order_staus_by_id"
)

type ChanType = int
const (
	ChanAlipay ChanType = 1
	ChanWeixin ChanType = 2
)

type OrderStatus = int
const (
	OrderStatusNone OrderStatus = iota
	OrderStatusPaySuccess OrderStatus = iota
	OrderStatusPayFail OrderStatus = iota
)

var (
	paysapiUid string
	paysapiToken string
	notifyUrl string
	returnUrl string
)

type PaymentReq struct {
	Uid       string
	Price     string // 单位：分
	IsType    ChanType	// 1:alipay 2:weipay
	NotifyUrl string
	ReturnUrl string
	OrderId   string
	OrderUid  string
	GoodsName string
	key       string
}

func (r *PaymentReq) ToString() string {
	return fmt.Sprintf("uid=%s&price=%s&istype=%d&notify_url=%s&return_url=%s&orderid=%s&orderuid=%s&goodsname=%s&key=%s",
		r.Uid, r.Price, r.IsType, r.NotifyUrl, r.ReturnUrl, r.OrderId, r.OrderUid, r.GoodsName, r.MakeSign())
}

func (r *PaymentReq) MakeSign() string {
	return Md5(fmt.Sprintf("%s%d%s%s%s%s%s%s%s", r.GoodsName, r.IsType, r.NotifyUrl, r.OrderId, r.OrderUid, r.Price, r.ReturnUrl, paysapiToken, paysapiUid))
}

type PaymentRes struct {
	Msg string `json:"msg"`
	//Data struct {
	//	Qrcode string `json:"qrcode"`
	//	Istype ChanType `json:"istype"`
	//	Realprice float32 `json:"realprice"`
	//} `json:"data"`
	Data interface{} `json:"data"`
	Code int `json:"code"`
	Url string `json:"url"`
}


type PaymentNotifyReq struct {
	PaysapiId string
	OrderId   string
	Price     string
	Realprice string
	OrderUid  string
	key       string
}

func (r *PaymentNotifyReq) MakeSign() string {
	return Md5(fmt.Sprintf("%s%s%s%s%s%s", r.OrderId, r.OrderUid, r.PaysapiId, r.Price, r.Realprice, paysapiToken))
}

func (r *PaymentNotifyReq) IsValid() bool {
	return r.key == r.MakeSign()
}

func (r *PaymentNotifyReq) FromString(str string) {

}

func RequestPayment(req *PaymentReq) (*PaymentRes, error) {

	logs.Info("before request parment req.Tostring=%s", req.ToString())
	resBody, err := utils.HttpPostBody(ApiPaymentJson, map[string]string{"Content-Type":"application/x-www-form-urlencoded"}, []byte(req.ToString()))
	if err != nil {
		return nil, fmt.Errorf("when RequestPayment orderId:%s error:%v req=%s", req.OrderId, err, req.ToString())
	}

	res := &PaymentRes{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, fmt.Errorf("when RequestPayment orderId:%s error:%v resBody=%s", req.OrderId, err, resBody)
	}

	return res, nil
}

func SimplePayment(orderId string, price int, payChan ChanType, orderUid, goodsName string) (*PaymentRes, error) {
	req := &PaymentReq{
		OrderId:orderId,
		Price:fmt.Sprintf("%.2f", float32(price)/100.0),
		IsType:payChan,
		OrderUid:orderUid,
		GoodsName:goodsName,
		ReturnUrl:returnUrl,
		NotifyUrl:notifyUrl,
		Uid:paysapiUid,
	}

	return RequestPayment(req)
}

type QueryPaymentReq struct {
	Uid string
	OrderId string
	r string
	key string
}

func (r *QueryPaymentReq) MakeSign() string {
	return Md5(fmt.Sprintf("%s%s%s%s", r.Uid, r.OrderId, r.r, paysapiToken))
}

func (r *QueryPaymentReq) ToString() string {
	return fmt.Sprintf("uid=%s&orderid=%s&r=%s&key=%s",
		r.Uid, r.OrderId, r.r, r.MakeSign())
}

type PaymentQueryRes struct {
	Msg string `json:"msg"`
	Data struct {
		OrderId string `json:"orderid"`
		Status int `json:"status"`
	} `json:"data"`
	Code int `json:"code"`
	Url string `json:"url"`
}

func QueryPayment(req *QueryPaymentReq) (OrderStatus, error) {

	resBody, err := utils.HttpPostBody(ApiQueryPayment, nil, []byte(req.ToString()))
	if err != nil {
		return OrderStatusNone, fmt.Errorf("when QueryPayment error:%v orderId:%s", err, req.OrderId)
	}

	res := &PaymentQueryRes{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return OrderStatusNone, fmt.Errorf("when QueryPayment error:%v orderId:%s", err, req.OrderId)
	}

	if res.Data.Status != OrderStatusPaySuccess {
		return res.Data.Status, fmt.Errorf("%s", res.Msg)
	}

	return res.Data.Status, nil
}

func SimpleQueryPayment(orderId string) (OrderStatus, error) {
	req := &QueryPaymentReq{
		OrderId:orderId,
		Uid:paysapiUid,
		r:utils.NewUUID(),
	}

	return QueryPayment(req)
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func SetPaysapi(uid, token string) {
	paysapiUid = uid
	paysapiToken = token
}

func SetNotifyUrl(url string) {
	notifyUrl = url
}
func SetReturnUrl(url string) {
	returnUrl = url
}
