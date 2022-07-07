package router

import (
	"github.com/gin-gonic/gin"
	"zph/controller"
)

func initMomentRouter(group *gin.RouterGroup) {
	c := controller.NewMomentController()
	group.POST("/moment", c.CreateMoment)
	group.GET("/moments", c.GetMoments)
	group.GET("/moment/:moment_id", c.GetMomentById)
	group.PUT("/moment/published/:moment_id", c.UpdatePublished)
	group.PUT("/moment/:moment_id", c.UpdateMoment)
	group.DELETE("/moment/:moment_id", c.DeleteMomentById)
}
