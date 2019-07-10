package protocol

type ThirdUserLoginRequest struct {
	Platform    string `json:"platform"`    //三方平台/渠道
	AppID       string `json:"appId"`       //用户来自哪一个应用
	ChannelID   string `json:"channelId"`   //用户来自哪一个渠道
	Device      Device `json:"device"`      //设备信息
	Name        string `json:"name"`        //微信平台名
	OpenID      string `json:"openId"`      //微信平台OpenID
	AccessToken string `json:"accessToken"` //微信AccessToken
}

type LoginRequest struct {
	AppID     string `json:"appId"`     //用户来自哪一个应用
	ChannelID string `json:"channelId"` //用户来自哪一个渠道
	IMEI      string `json:"imei"`
	Device    Device `json:"device"`
}

type LoginResponse struct {
	Code     int          `json:"code"`
	Name     string       `json:"name"`
	Uid      int64        `json:"uid"`
	HeadUrl  string       `json:"headUrl"`
	FangKa   int64        `json:"fangka"`
	Sex      int          `json:"sex"` //[0]未知 [1]男 [2]女
	IP       string       `json:"ip"`
	Port     int          `json:"port"`
	PlayerIP string       `json:"playerIp"`
	Config   ClientConfig `json:"config"`
	Message  []string     `json:"message"`
	ClubList []ClubItem   `json:"clubList"`
	Debug    int          `json:"debug"`
}

type LoginToGameServerResponse struct {
	Uid      int64  `json:"acId"`
	Nickname string `json:"nickname"`
	HeadUrl  string `json:"headUrl"`
	Sex      int    `json:"sex"`
	FangKa   int    `json:"fangka"`
}

type LoginToGameServerRequest struct {
	Name    string `json:"name"`
	Uid     int64  `json:"uid"`
	HeadUrl string `json:"headUrl"`
	Sex     int    `json:"sex"` //[0]未知 [1]男 [2]女
	FangKa  int    `json:"fangka"`
	IP      string `json:"ip"`
}
