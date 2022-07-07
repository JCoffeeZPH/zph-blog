package response

type LoginLogResponse struct {
	LoginLogId int `json:"id"`
	Operator string `json:"username"`
	IP string `json:"ip"`
	IPSource string `json:"ip_source"`
	OS string `json:"os"`
	Browser string `json:"browser"`
	Status bool `json:"status"`
	Description string `json:"description"`
	CreateTime string `json:"create_time"`
}

type ReturnLoginLogResponse struct {
	Total uint64 `json:"total"`
	Logs []LoginLogResponse `json:"logs"`
}