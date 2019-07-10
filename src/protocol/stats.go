package protocol

type RetentionLite struct {
	Login int64  `json:"login"`
	Rate  string `json:"rate"`
}

type Retention struct {
	Date     int   `json:"date"`
	Register int64 `json:"register"`

	Retention_1  RetentionLite `json:"retention_1"`  //次日
	Retention_2  RetentionLite `json:"retention_2"`  //2日
	Retention_3  RetentionLite `json:"retention_3"`  //3日
	Retention_7  RetentionLite `json:"retention_7"`  //7日
	Retention_14 RetentionLite `json:"retention_14"` //14日
	Retention_30 RetentionLite `json:"retention_30"` //30日
}

type RetentionResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type CommonStatsItem struct {
	Date  int64 `json:"date"`
	Value int64 `json:"value"`
}

//房卡消耗
type CardConsume CommonStatsItem

//活跃用户
type ActivationUser CommonStatsItem
