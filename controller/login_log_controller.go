package controller

import (
	"github.com/gin-gonic/gin"
	"zph/constants"
	api_error "zph/error"
	"zph/models/request"
	"zph/models/response"
	"zph/service"
	"zph/session"
)

type LoginLogController struct {
	BaseController
	loginLogService service.LoginLogService
}

func NewLoginLogController() *LoginLogController {
	return &LoginLogController{
		loginLogService: service.NewLoginLogService(),
	}
}

func (controller *LoginLogController) GetLoginLogs(c *gin.Context) {
	req, err := controller.parseGetLoginLogsParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	offset := (req.Page - 1)* req.PerPage
	limit := req.PerPage
	userId := session.GetUserId(c)
	resp, count := controller.loginLogService.GetLoginLogs(req.StartTime, req.EndTime, offset, limit, int(userId))
	controller.Success(c, &response.ReturnLoginLogResponse{
		Total: count,
		Logs: resp,
	})
}

func (controller *LoginLogController) DeleteLoginLog(c *gin.Context) {
	req := &request.DeleteLoginLogRequest{}
	err := c.ShouldBindUri(req)
	if err != nil {
		panic(api_error.ParamError)
	}
	controller.loginLogService.DeleteLoginLog(req.LogId)
	controller.NoContent(c)
}

func (controller *LoginLogController) parseGetLoginLogsParams(c *gin.Context)(*request.GetLoginLogsRequest, error)  {
	req := &request.GetLoginLogsRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		return nil, err
	}
	if req.Page == 0 {
		req.Page = constants.DefaultPage
	}

	if req.PerPage == 0 {
		req.PerPage = constants.DefaultPerPage
	}
	return req, nil
}