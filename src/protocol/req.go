package protocol

import "goer/pkg/constant"

type ReJoinDeskRequest struct {
	DeskNo string `json:"deskId"`
}

type ReJoinDeskResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type ReEnterDeskRequest struct {
	DeskNo string `json:"deskId"`
}

type ReEnterDeskResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type JoinDeskRequest struct {
	Version string `json:"version"`
	DeskNo  string `json:"deskId"`
}

type TableInfo struct {
	DeskNo    string              `json:"deskId"`
	CreatedAt int64               `json:"createdAt"`
	Creator   int64               `json:"creator"`
	Title     string              `json:"title"`
	Desc      string              `json:"desc"`
	Status    constant.DeskStatus `json:"status"`
	Round     uint32              `json:"round"`
	Mode      int                 `json:"mode"`
}

type JoinDeskResponse struct {
	Code      int       `json:"code"`
	Error     string    `json:"error"`
	TableInfo TableInfo `json:"tableInfo"`
}

//选择执行的动作
type OpChoosed struct {
	Type   int
	TileID int
}

type OpChooseRequest struct {
	OpType int `json:"optype"`
	Index  int `json:"index"`
}

type CheckOrderRequest struct {
	OrderID string `json:"orderid"`
}

type CheckOrderResponse struct {
	Code   int    `json:"code"`
	Error  string `json:"error"`
	FangKa int    `json:"fangka"`
}

type DissolveStatusItem struct {
	DeskPos int    `json:"deskPos"`
	Status  string `json:"status"`
}

type DissolveResponse struct {
	DissolveUid    int64                `json:"dissolveUid"`
	DissolveStatus []DissolveStatusItem `json:"dissolveStatus"`
	ResetTime      int32                `json:"resetTime"`
}

type DissolveStatusRequest struct {
	Result bool `json:"result"`
}

type DissolveResult struct {
	DeskPos int `json:"deskPos"`
}

type DissolveStatusResponse struct {
	DissolveStatus []DissolveStatusItem `json:"dissolveStatus"`
	RestTime       int32                `json:"restTime"`
}

type PlayerOfflineStatus struct {
	Uid     int64 `json:"uid"`
	Offline bool  `json:"offline"`
}

type CoinChangeInformation struct {
	Coin int64 `json:"coin"`
}
