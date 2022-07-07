package router

import (
	"github.com/gin-gonic/gin"
	"zph/controller"
)

func initLoginRouter(group *gin.RouterGroup) {
	c := controller.NewUserController()
	group.POST("/login", c.LoginController)
}
