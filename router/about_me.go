package router

import (
	"github.com/gin-gonic/gin"
	"zph/controller"
)

func initAboutMeRouter(group *gin.RouterGroup) {
	c := controller.NewAboutMeController()
	group.GET("/about_me", c.GetAboutMeInfo)
	group.PUT("/about_me", c.UpdateAboutMe)
}
