package provider

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"goer/db/model"
	"goer/pkg/algoutil"
	"goer/pkg/errutil"
	"goer/protocol"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type wechat struct {
	appkey        string
	appId         string
	merId         string
	unifyOrderURL string
	callbackURL   string
}

var Wechat = &wechat{}

//请求 URL地址：https://api.mch.weixin.qq.com/pay/unifiedorder 需要填入的参数
// refs: https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_1

type UnifyOrderReq struct {
	Appid          string `xml:"appid"`
	Body           string `xml:"body"`             //商品描述
	MchID          string `xml:"mch_id"`           //微信支付分配的商户号
	NonceStr       string `xml:"nonce_str"`        //随机字符串
	NotifyURL      string `xml:"notify_url"`       //通知地址
	TradeType      string `xml:"trade_type"`       //支付类型 APP
	SpbillCreateIP string `xml:"spbill_create_ip"` //终端IP
	TotalFee       int    `xml:"total_fee"`        //总金额
	OutTradeNo     string `xml:"out_trade_no"`     //商户订单号
	Sign           string `xml:"sign"`             //签名
}

type UnifyOrderResp struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Appid      string `xml:"appid"`
	MchID      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	PrepayID   string `xml:"prepay_id"`
	TradeType  string `xml:"trade_type"`
	ErrCode    string `xml:"err_code"`
}

//微信支付计算签名的函数  md5的签名方式  还可以是用 HMAC-SHA256签名方式
func signCalculator(params map[string]interface{}, key string) (string, error) {
	if params == nil || key == "" {
		return "", errutil.ErrInvalidParameter
	}
	sortedKeys := make([]string, 0)
	for k := range params {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	buf := &bytes.Buffer{}
	for _, k := range sortedKeys {
		if params[k] == nil {
			continue
		}
		switch params[k].(type) {
		case string:
			if len(params[k].(string)) == 0 {
				continue
			}
		case int:
			if params[k].(int) == 0 {
				continue
			}
		}
		fmt.Fprintf(buf, "%s=%v&", k, params[k])
	}
	fmt.Fprintf(buf, "key=%v", key)
	md5Ctx := md5.New()
	signStr := buf.Bytes()
	md5Ctx.Write(signStr)
	cipherStr := md5Ctx.Sum(nil)
	sign := strings.ToUpper(hex.EncodeToString(cipherStr)) //将[]byte转16进制字符串
	return sign, nil
}

//HMAC-SHA256签名方式
func computeHmacSha256(msg []byte, key string) string {
	k := []byte(key)
	h := hmac.New(sha256.New, k)
	h.Write(msg)
	sha := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(sha))
}

func verify(m map[string]interface{}, key, sign string) bool {
	signed, _ := signCalculator(m, key)
	return sign == signed
}

func (wc *wechat) CreateOrderResponse(order *model.Order) (interface{}, error) {
	req := UnifyOrderReq{
		Appid:          wc.appId,             //微信开发平台app的appid
		Body:           order.ProductName,    //产品名
		MchID:          wc.merId,             //商户ID
		NonceStr:       algoutil.RandStr(32), //随机数
		NotifyURL:      wc.callbackURL,
		TradeType:      "APP",
		SpbillCreateIP: strings.Split(order.Remote, ":")[0],
		TotalFee:       order.ProductCount * 1,
		OutTradeNo:     order.OrderId,
	}
	m := make(map[string]interface{}, 0)

	m["appid"] = req.Appid
	m["body"] = req.Body
	m["mch_id"] = req.MchID
	m["notify_url"] = req.NotifyURL
	m["trade_type"] = req.TradeType
	m["spbill_create_ip"] = req.SpbillCreateIP
	m["total_fee"] = req.TotalFee
	m["out_trade_no"] = req.OutTradeNo
	m["nonce_str"] = req.NonceStr

	sign, err := signCalculator(m, wc.appkey)
	if err != nil {
		return nil, err
	}
	req.Sign = sign
	bytesReq, err := xml.Marshal(req)
	if err != nil {
		return nil, err
	}
	//unified order 接口需要http body中的根节点是<xml></xml>这种,所以这里需要replace一下
	strReq := strings.Replace(string(bytesReq), "UnifyOrderURL", "xml", -1)
	bytesReq = []byte(strReq)

	log.Debugf("prepay id request: %s", strReq)

	request, err := http.NewRequest("POST", wc.unifyOrderURL, bytes.NewReader(bytesReq))
	if err != nil {
		log.Errorf("create unify order failed: %s", err.Error())
		return nil, err
	}
	request.Header.Set("Accept", "application/xml")
	request.Header.Set("Content-Type", "application/xml;charset=utf-8")
	c := http.Client{}
	response, err := c.Do(request)
	if err != nil {
		log.Errorf("request unify order failed: %s", err.Error())
		return nil, err
	}
	println(response.StatusCode)

	xmlResp := &UnifyOrderResp{}
	if err := xml.NewDecoder(response.Body).Decode(xmlResp); err != nil {
		fmt.Println("unify order request prepay id failed " + xmlResp.ReturnMsg)
		return nil, err
	}

	const (
		fail = "FAIL"
	)
	log.Debugf("prepay id response: %+v", xmlResp)

	if xmlResp.ResultCode == fail {
		fmt.Println("unify order request prepay id failed: " + xmlResp.ReturnMsg)
		return nil, errutil.ErrRequestPrepayIDFailed
	}

	m = make(map[string]interface{}, 0)

	ret := protocol.CreateOrderWechatResponse{
		OrderId:   order.OrderId,
		PrepayID:  xmlResp.PrepayID,
		AppID:     xmlResp.Appid,
		PartnerId: xmlResp.MchID,
		NonceStr:  algoutil.RandStr(32),
		Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
		Extra:     order.Extra,
	}

	m["appid"] = ret.AppID
	m["partnerid"] = ret.PartnerId
	m["prepayid"] = ret.PrepayID
	m["noncestr"] = ret.NonceStr
	m["timestamp"] = ret.Timestamp
	m["package"] = "Sign=WXPay"
	ret.Sign, err = signCalculator(m, wc.appkey)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

const format = "20060102150405"

//这里是微信的服务器调用该接口来通知
func (wc *wechat) Notify(r *protocol.WechatOrderCallbackRequest) (*model.Trade, interface{}, error) {
	reqMap := make(map[string]interface{}, 0)

	reqMap["return_code"] = r.ReturnCode
	reqMap["return_msg"] = r.ReturnMsg
	reqMap["appid"] = r.Appid
	reqMap["mch_id"] = r.MchID
	reqMap["device_info"] = r.DeviceInfo
	reqMap["nonce_str"] = r.Nonce
	//sign
	reqMap["result_code"] = r.ResultCode
	reqMap["err_code"] = r.ErrCode
	reqMap["err_code_des"] = r.ErrCodeDes

	reqMap["openid"] = r.Openid
	reqMap["is_subscribe"] = r.IsSubscribe
	reqMap["trade_type"] = r.TradeType
	reqMap["bank_type"] = r.BankType
	reqMap["total_fee"] = r.TotalFee
	reqMap["fee_type"] = r.FeeType
	reqMap["cash_fee"] = r.CashFee
	reqMap["cash_fee_type"] = r.CashFeeType

	reqMap["coupon_fee"] = r.CouponFee
	reqMap["coupon_count"] = r.CouponCount
	reqMap["coupon_id_$n"] = r.CouponIDDollarN
	reqMap["coupon_fee_$n"] = r.CouponFeeDollarN

	reqMap["transaction_id"] = r.TransactionID
	reqMap["out_trade_no"] = r.OutTradeNo
	reqMap["attach"] = r.Attach
	reqMap["time_end"] = r.TimeEnd

	var resp protocol.WechatOrderCallbackResponse
	if verify(reqMap, wc.appkey, r.Sign) {
		resp.ReturnCode = "SUCCESS"
		resp.ReturnMsg = "OK"
	} else {
		resp.ReturnCode = "FAIL"
		resp.ReturnMsg = "failed to verify sign, please retry!"
	}

	formatStr := "<xml><return_code><![CDATA[%s]]></return_code><return_msg><![CDATA[%s]]></return_msg></xml>"
	cbResp := fmt.Sprintf(formatStr, resp.ReturnCode, resp.ReturnMsg)

	fmt.Println("response: ", cbResp)
	trade := &model.Trade{}
	trade.PayOrderId = r.TransactionID
	trade.OrderId = r.OutTradeNo

	createAt, err := time.Parse(format, r.TimeEnd)
	if err != nil {
		createAt = time.Now()
	}
	trade.PayCreatedAt = createAt.Unix()

	payAt, err := time.Parse(format, r.TimeEnd)
	if err != nil {
		payAt = time.Now()
	}
	trade.PayAt = payAt.Unix()
	trade.MerchantId = r.MchID
	trade.ConsumerId = r.Openid
	trade.Raw = r.Raw
	return trade, cbResp, nil
}

func (wc *wechat) Setup() error {
	log.Info("pay_provider: wechat setup")

	var (
		appId         = viper.GetString("wechat.appid")
		appKey        = viper.GetString("wechat.appsecret")
		merId         = viper.GetString("wechat.mer_id")
		unifyOrderURL = viper.GetString("wechat.unify_order_url")
		callbackURL   = viper.GetString("wechat.callback_url")
	)
	if appId == "" || appKey == "" || merId == "" || unifyOrderURL == "" || callbackURL == "" {
		log.Debugf("appID=%s appKey=%s merId=%s unifyOrderURL=%s callbackURL=%s",
			appId,
			appKey,
			merId,
			unifyOrderURL,
			callbackURL)
		return errors.New("the wechat's config is invalid.")
	}
	wc.appId = appId
	wc.appkey = appKey
	wc.merId = merId
	wc.unifyOrderURL = unifyOrderURL
	wc.callbackURL = callbackURL
	return nil
}
