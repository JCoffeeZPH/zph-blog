package router

import (
	"github.com/gin-gonic/gin"
	"zph/controller"
)

func initLoginLogRouter(group *gin.RouterGroup) {
	c := controller.NewLoginLogController()

	group.GET("/login_logs", c.GetLoginLogs)
	group.DELETE("/login_log/:log_id", c.DeleteLoginLog)
}
