package request

type DeleteLoginLogRequest struct {
	LogId int `uri:"log_id" binding:"required"`
}

type GetLoginLogsRequest struct {
	StartTime string `form:"start_time"`
	EndTime string `form:"end_time"`
	Page int `form:"page"`
	PerPage int `form:"per_page"`
}