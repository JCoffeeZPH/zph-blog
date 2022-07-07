package router

import (
	"github.com/gin-gonic/gin"
	"zph/controller"
)

func initDashboardRouter(group *gin.RouterGroup) {
	c := controller.NewDashboardController()
	group.GET("/dashboard", c.GetDashBoard)
}
