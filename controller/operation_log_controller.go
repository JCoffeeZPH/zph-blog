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

type OperationLogController struct {
	BaseController
	logService *service.OperationLogService
}

func NewOperationLogController() *OperationLogController {
	return &OperationLogController{
		logService: service.NewOperationLogService(),
	}
}

func (controller *OperationLogController) GetOperationLogs(c *gin.Context) {
	req, err := controller.parseGetOperationLogsParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	offset := (req.Page - 1)* req.PerPage
	limit := req.PerPage
	userId := session.GetUserId(c)
	resp, count := controller.logService.GetOperationLogs(req.StartTime, req.EndTime, offset, limit, int(userId))
	controller.Success(c, &response.LogResponse{
		Total: count,
		Logs: resp,
	})
}

func (controller *OperationLogController) DeleteOperationLog(c *gin.Context) {
	req := &request.DeleteOperationLogRequest{}
	err := c.ShouldBindUri(req)
	if err != nil {
		panic(api_error.ParamError)
	}
	controller.logService.DeleteOperationLog(req.LogId)
	controller.NoContent(c)
}

func (controller *OperationLogController) parseGetOperationLogsParams(c *gin.Context)(*request.GetOperationLogsRequest, error)  {
	req := &request.GetOperationLogsRequest{}
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