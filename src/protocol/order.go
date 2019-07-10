package protocol

type CreateOrderRequest struct {
	AppID          string `json:"appId"`     //来自哪个应用的订单
	ChannelID      string `json:"channelId"` //来自哪个渠道的订单
	Platform       string `json:"platform"`  //支付平台
	ProductionName string `json:"name"`
	ProductCount   int    `json:"count"`  //房卡数量
	Extra          string `json:"extra"`  //额外信息
	Device         Device `json:"device"` //设备信息
	Uid            int64  `json:"uid"`    //Token
}

type OrderListRequest struct {
	Offset    int    `json:"offset"`
	Count     int    `json:"count"`
	Status    uint8  `json:"status"`
	Start     int64  `json:"start"` //时间起点
	End       int64  `json:"end"`   //时间终点
	PayBy     string `json:"pay_by"`
	Uid       string `json:"uid"`        //用户id
	OrderID   string `json:"order_id"`   //订单号
	AppID     string `json:"appid"`      //来自哪个应用的订单
	ChannelID string `json:"channel_id"` //来自哪个渠道的订单
}

type OrderListResponse struct {
	Code  int         `json:"code"`
	Data  []OrderInfo `json:"data"`
	Total int         `json:"total"`
}

type CreateOrderWechatResponse struct {
	AppID     string `json:"appid"`
	PartnerId string `json:"partnerid"`
	OrderId   string `json:"orderid"`
	PrepayID  string `json:"prepayid"`
	NonceStr  string `json:"noncestr"`
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
	Extra     string `json:"extra"`
}

type WechatOrderCallbackRequest struct {
	DeviceInfo       string `xml:"device_info,omitempty"`
	ErrCode          string `xml:"err_code,omitempty"`
	ErrCodeDes       string `xml:"err_code_des,omitempty"`
	Attach           string `xml:"attach,omitempty"`        //商家数据包
	CashFeeType      string `xml:"cash_fee_type,omitempty"` //现金支付货币种类
	CouponFee        int    `xml:"coupon_fee,omitempty"`    //代金券金额
	CouponCount      int    `xml:"coupon_count,omitempty"`  //代金券数量
	CouponIDDollarN  string `xml:"coupon_id_$n,omitempty"`  //代金券或立减优惠ID,$n为下标，从0开始编号
	CouponFeeDollarN string `xml:"coupon_fee_$n,omitempty"` //单个代金券或立减优惠支付金额,$n为下标，从0开始编号

	ReturnCode    string `xml:"return_code"`
	ReturnMsg     string `xml:"return_msg"`
	Appid         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	Nonce         string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	ResultCode    string `xml:"result_code"`
	Openid        string `xml:"openid"`
	IsSubscribe   string `xml:"is_subscribe"`
	TradeType     string `xml:"trade_type"`
	BankType      string `xml:"bank_type"`
	TotalFee      int    `xml:"total_fee"`
	FeeType       string `xml:"fee_type"`       //货币类型
	CashFee       int    `xml:"cash_fee"`       //现金支付金额
	TransactionID string `xml:"transaction_id"` //微信支付订单号
	OutTradeNo    string `xml:"out_trade_no"`
	TimeEnd       string `xml:"time_end"`

	Raw string
}

type WechatOrderCallbackResponse struct {
	ReturnCode string `xml:"return_code,cdata"`
	ReturnMsg  string `xml:"return_msg,cdata"`
}

type RechargeRequest struct {
	Count int64 `json:"count"`
	Uid   int64 `json:"uid"`
}
