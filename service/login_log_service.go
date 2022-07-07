package service

import (
	"github.com/gin-gonic/gin"
	"strings"
	"zph/constants"
	"zph/dao"
	"zph/lib/client"
	"zph/lib/common"
	"zph/models/db"
	"zph/models/response"
	"zph/session"
	"zph/utils"
)

type LoginLogService struct {
	loginLogDao   dao.LoginLogDao
	ipClient      client.IPClient
	browserClient client.BrowserClient
}

func NewLoginLogService() LoginLogService {
	return LoginLogService{
		loginLogDao:   dao.NewLoginLogDao(),
		ipClient:      client.NewIPClient(),
		browserClient: client.NewBrowserClient(),
	}
}

func (service *LoginLogService) CreateNewLoginLog(ctx *gin.Context, status bool) {
	loginLog := service.parseParamsToDdModel(ctx, status)
	service.loginLogDao.CreateLoginLog(loginLog)
}

func (service *LoginLogService) DeleteLoginLog(logId int) {
	service.loginLogDao.DeleteLoginLogByLogId(logId)
}

func (service *LoginLogService) GetLoginLogs(startTime, endTime string, offset, limit, userId int) ([]response.LoginLogResponse, uint64) {

	stime := utils.StrToUnix(startTime)
	etime := utils.StrToUnix(endTime)
	count := service.loginLogDao.GetLoginLogCount(stime, etime, userId)
	dbLogs := service.loginLogDao.GetLoginLogs(stime, etime, offset, limit, userId)

	return service.newLoginLogResponse(dbLogs), count
}

func (service *LoginLogService) newLoginLogResponse(logs []db.LoginLog) []response.LoginLogResponse {
	resp := make([]response.LoginLogResponse, 0)
	for _, loginLog := range logs {
		resp = append(resp, response.LoginLogResponse{
			LoginLogId:  loginLog.LoginLogId,
			Operator:    loginLog.Username,
			IP:          loginLog.IP,
			IPSource:    loginLog.IPSource,
			OS:          loginLog.OS,
			Browser:     loginLog.Browser,
			Status:      utils.IntToBool(loginLog.Status),
			Description: loginLog.Description,
			CreateTime:  utils.TimeFormat(loginLog.CreateTime),
		})
	}
	return resp
}

func (service *LoginLogService) parseParamsToDdModel(ctx *gin.Context, status bool) *db.LoginLog {
	userId := session.GetUserId(ctx)
	username := session.GetUsername(ctx)
	ipDetails, err := service.ipClient.GetIPAndAttribution()
	if err != nil {
		panic(common.ServiceError{Err: err, API: "parseParamsToDdModel"})
	}
	userAgent := ctx.Request.UserAgent()
	browser, err := service.browserClient.GetBrowserDetails(userAgent)
	if err != nil {
		panic(common.ServiceError{Err: err, API: "parseParamsToDdModel"})
	}
	os := "postman"
	browserType := "postman"
	if browser != nil {
		os = browser.Data[constants.OS][:strings.Index(browser.Data[constants.OS], "(")]
		browserType = browser.Data[constants.Browser]
	}
	description := constants.LoginSuccess
	if !status {
		description = constants.LoginDefeat
	}

	return &db.LoginLog{
		Username:    username,
		IP:          ipDetails.IP,
		IPSource:    ipDetails.IpSource,
		OS:          os,
		Browser:     browserType,
		Status:      utils.BoolToInt(status),
		Description: description,
		CreateTime:  utils.NowTime(),
		UserAgent:   userAgent,
		UserId:      int(userId),
	}
}
