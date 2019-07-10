package protocol

import "fmt"

type OrderInfo struct {
	OrderId      string `json:"orderId"`       //订单号
	Uid          string `json:"uid"`           //接受者ID
	AppId        string `json:"appid"`         //应用id
	ServerName   string `json:"server_name"`   //区服名
	RoleID       string `json:"role_id"`       //角色ID
	Extra        string `json:"extra"`         //额外信息
	Imei         string `json:"imei"`          //imei
	ProductName  string `json:"product_name"`  //商品名
	PayBy        string `json:"pay_by"`        //收支渠道:alipay,wechat ...
	ProductCount int    `json:"product_count"` //商品数量
	Money        int    `json:"money"`         //标价
	RealMoney    int    `json:"real_money"`    //实际售价
	Status       int    `json:"status"`        //订单状态  1-创建 2-完成 3-游戏服务器已确认
	CreatedAt    int64  `json:"created_at"`    //发放时间
}

type DailyStats struct {
	Score     int      `json:"score"`      //战绩
	AsCreator int64    `json:"as_creator"` //开放次数
	Win       int      `json:"win"`        //赢的次数
	DeskNos   []string `json:"desks"`      //所参加的桌号
}

type UserStatsInfo struct {
	ID            int64  `json:"id"`
	Uid           int64  `json:"uid"`
	Name          string `json:"name"`
	RegisterAt    int64  `json:"register_at"`
	RegisterIP    string `json:"register_ip"`
	LatestLoginAt int64  `json:"latest_login_at"`
	LatestLoginIP string `json:"latest_login_ip"`
	TotalMatch    int64  `json:"total_match"` //总对局数
	RemainCard    int64  `json:"remain_card"` //剩余房卡数

	StatsAt []int64               //统计时间
	Stats   map[int64]*DailyStats //时间对应的数据
}

type Device struct {
	IMEI   string `json:"imei"`   //设备的IMEI号
	OS     string `json:"os"`     //os版本号
	Model  string `json:"model"`  //硬件型号
	IP     string `json:"ip"`     //内网IP
	Remote string `json:"remote"` //外网IP
}

type CommonResponse struct {
	Code int         `json:"code"` //状态码
	Data interface{} `json:"data"` //整数状态
}

type StringResponse struct {
	Code int    `json:"code"` //状态码
	Data string `json:"data"` //字符串数据
}

var SuccessResponse = StringResponse{0, "success"}

const (
	RegTypeThird = 5
)

var EmptyMessage = &None{}
var SuccessMessage = &StringMessage{Message: "success"}

type None struct{}

type StringMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

//听牌信息
type Ting struct {
	Index int   `json:"index"`
	Hu    []int `json:"hu"`
}

//所有被听的牌
type Tings []Ting

//所有可执行的操作
type Ops []Op

//提示
type Hint struct {
	Ops   Ops   `json:"ops"`
	Tings Tings `json:"tings"`
	Uid   int64 `json:"uid"`
}

func (h *Hint) String() string {
	return fmt.Sprintf("UID=%d, ops=%+v, Tings=%+v", h.Uid, h.Ops, h.Tings)
}

type ClientConfig struct {
	Version     string `json:"version"`
	Android     string `json:"android"`
	IOS         string `json:"ios"`
	Heartbeat   int    `json:"heartbeat"`
	ForceUpdate bool   `json:"forceUpdate"`
	Title       string `json:"title"` //分享标题
	Desc        string `json:"desc"`  //分享描述
	Daili1      string `json:"daili1"`
	Daili2      string `json:"daili2"`
	Kefu1       string `json:"kefu1"`
	AppId       string `json:"appId"`
	AppKey      string `json:"appKey"`
}
