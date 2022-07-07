package service

import (
	"zph/dao"
	"zph/models/db"
	"zph/models/response"
	"zph/utils"
)

type OperationLogService struct {
	operationLogDao dao.OperationLogDao
}

func NewOperationLogService() *OperationLogService {
	return &OperationLogService{
		operationLogDao: dao.NewOperationLogDao(),
	}
}

func (service *OperationLogService) CreateNewOperationLog(operationLog *db.OperationLog) {
	service.operationLogDao.CreateOperationLog(operationLog)
}

func (service *OperationLogService) DeleteOperationLog(logId int) {
	service.operationLogDao.DeleteLogByLogId(logId)
}

func (service *OperationLogService) GetOperationLogs(startTime, endTime string, offset, limit, userId int) ([]response.OperationLogResponse, uint64){

	stime := utils.StrToUnix(startTime)
	etime := utils.StrToUnix(endTime)
	count := service.operationLogDao.GetOperationLogCount(stime, etime, userId)
	dbLogs := service.operationLogDao.GetOperations(stime, etime , offset, limit, userId)
	return service.newOperationResponse(dbLogs), count
}

func (service *OperationLogService) newOperationResponse(logs []db.OperationLog) []response.OperationLogResponse {
	resp := make([]response.OperationLogResponse,0)
	for _, operationLog := range logs {
		resp = append(resp, response.OperationLogResponse{
			OperationLogId: operationLog.OperationLogId,
			Operator: operationLog.Username,
			Method: operationLog.Method,
			Uri: operationLog.Uri,
			Description: operationLog.Description,
			IP: operationLog.Ip,
			IPSource: operationLog.IpSource,
			OS: operationLog.OS,
			Browser: operationLog.Browser,
			Times: operationLog.Times,
			CreateTime: utils.TimeFormat(operationLog.CreateTime),
		})
	}
	return resp
}
