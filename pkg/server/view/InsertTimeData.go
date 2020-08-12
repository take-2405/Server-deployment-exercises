package view

//timeDataResponse 保存時間のレスポンス
type InsertTimeDataResponse struct {
	TimeID string `json:"timeID"`
	Status string `json:"status"`
	Timedata TimeData `json:"timedata"`
}

type TimeData struct{
	Year int `json:"year"`
	Month int `json:"month"`
	Day int `json:"data"`
	Hour int `json:"hour"`
	Minute int `json:"minute"`
}