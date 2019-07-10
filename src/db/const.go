package db

const (
	defaultMaxConns = 10
)

//User 表中role字段的取值
const (
	RoleTypeAdmin = 1 //管理员账号
	RoleTypeThird = 2 //三方平台账号
)

const (
	UserOffline = 1 //离线
	UserOnline  = 2 //在线
)

const (
	StatusNormal  = 1 //正常
	StatusDeleted = 2 //删除
	StatusFreezed = 3 //账号冻结
	StatusBound   = 4 //绑定
)

//订单状态
const (
	OrderStatusCreated  = 1 //创建
	OrderStatusPayed    = 2 //完成
	OrderStatusNotified = 3 //已确认
)

const (
	OrderTypeUnknown      = iota
	OrderTypeBuyToken     //购买令牌
	OrderTypeConsumeToken //消费代币(eg:使用令牌购买游戏中的道具,比如房卡)
	OrderTypeConsume3rd   //第三方支付平台消费(eg:直接使用alipay,wechat购买游戏中的道具)
	OrderTypeTest         //支付测试
)

const (
	dayInSecond = 24 * 60 * 60

	day1  = dayInSecond
	day2  = day1 * 2
	day3  = day1 * 3
	day7  = day1 * 7
	day14 = day1 * 14
	day30 = day1 * 30
)
