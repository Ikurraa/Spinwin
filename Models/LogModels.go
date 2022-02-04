package Models

type Log struct {
	User_id     int    `json:"userid"`
	Last_update string `json:"lastupdate"`
}

func (log *Log) TableName() string {
	return "tbl_log"
}

type ViewLog struct {
	User_id     int    `json:"userid"`
	Username    string `json:"username"`
	Last_update string `json:"lastupdate"`
	Update_at   string `json:"updateat"`
}

func (log *ViewLog) TableName() string {
	return "v_ticketlog"
}
