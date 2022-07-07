package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"zph/constants"
	"zph/lib/client"
	"zph/models/db"
	"zph/service"
	"zph/session"
	"zph/utils"
)

func OperationLog() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		saveOperationLog(ctx, start)
	}
}

func saveOperationLog(ctx *gin.Context, start time.Time) {
	operation := ctx.GetString(constants.OperationKey)

	ipClient := client.NewIPClient()
	browserClient := client.NewBrowserClient()
	ipDetails, err := ipClient.GetIPAndAttribution()
	if err != nil {
		return
	}

	userAgent := ctx.Request.UserAgent()
	browserDetail, e := browserClient.GetBrowserDetails(userAgent)
	if e != nil {
		return
	}
	os := "postman"
	browser := "postman"
	if browserDetail != nil {
		os = browserDetail.Data[constants.OS][:strings.Index(browserDetail.Data[constants.OS], "(")]
		browser = browserDetail.Data[constants.Browser]
	}

	times := time.Since(start).Milliseconds()
	username := session.GetUsername(ctx)
	uri := ctx.Request.RequestURI
	method := ctx.Request.Method

	createTime := utils.NowTime()
	ip := ipDetails.IP
	ipSource := ipDetails.IpSource
	operationLog := &db.OperationLog{
		Username:    username,
		Uri:         uri,
		Method:      method,
		Description: operation,
		Ip:          ip,
		IpSource:    ipSource,
		OS:          os,
		Browser:     browser,
		Times:       times,
		CreateTime:  createTime,
		UserAgent:   userAgent,
		UserId:      session.GetUserId(ctx),
	}

	s := service.NewOperationLogService()
	s.CreateNewOperationLog(operationLog)
}
