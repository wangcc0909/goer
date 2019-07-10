package protocol

type HistoryLite struct {
	Id           int64  `json:"id"`
	DeskId       int64  `json:"desk_id"`
	Mode         int    `json:"mode"`
	BeginAt      int64  `json:"begin_at"`
	BeginAtStr   string `json:"begin_at_str"`
	EndAt        int64  `json:"end_at"`
	PlayName0    string `json:"play_name0"`
	PlayName1    string `json:"play_name1"`
	PlayName2    string `json:"play_name2"`
	PlayName3    string `json:"play_name3"`
	ScoreChange0 int    `json:"score_change0"`
	ScoreChange1 int    `json:"score_change1"`
	ScoreChange2 int    `json:"score_change2"`
	ScoreChange3 int    `json:"score_change3"`
}

type History struct {
	HistoryLite
	Snapshot string `json:"snapshot"`
}

type HistoryLiteListResponse struct {
	Code  int           `json:"code"`
	Total int64         `json:"total"` //总数量
	Data  []HistoryLite `json:"data"`
}

type HistoryByIDResponse struct {
	Code int      `json:"code"`
	Data *History `json:"data"`
}
