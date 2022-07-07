package router

import (
	"github.com/gin-gonic/gin"
	"zph/controller"
)

func initSiteSettingRouter(group *gin.RouterGroup) {
	c := controller.NewSiteSettingController()
	group.GET("/site_settings", c.GetSettings)
	group.POST("/site_settings", c.UpdateSiteSettings)
}
