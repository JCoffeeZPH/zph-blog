package response

type OperationLogResponse struct {
	OperationLogId int `json:"id"`
	Operator string `json:"operator"`
	Method string `json:"method"`
	Uri string `json:"uri"`
	Description string `json:"description"`
	IP string `json:"ip"`
	IPSource string `json:"ip_source"`
	OS string `json:"os"`
	Browser string `json:"browser"`
	Times int64 `json:"times"`
	CreateTime string `json:"create_time"`
}

type LogResponse struct {
	Total uint64 `json:"total"`
	Logs []OperationLogResponse `json:"logs"`
}