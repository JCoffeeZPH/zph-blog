package router

import (
	"github.com/gin-gonic/gin"
	"zph/controller"
)

func initOperationLogRouter(group *gin.RouterGroup) {
	c := controller.NewOperationLogController()

	group.GET("/operation_logs", c.GetOperationLogs)
	group.DELETE("/operation_log/:log_id", c.DeleteOperationLog)
}
