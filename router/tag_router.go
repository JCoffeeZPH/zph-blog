package router

import (
	"github.com/gin-gonic/gin"
	"zph/controller"
)

func initTagRouter(group *gin.RouterGroup) {
	tagController := controller.NewTagController()
	group.GET("/tags", tagController.GetTags)
	group.POST("/tag", tagController.CreateTag)
	group.PUT("/tag", tagController.UpdateTag)
	group.DELETE("/tag/:tag_id", tagController.DeleteTagById)
}
