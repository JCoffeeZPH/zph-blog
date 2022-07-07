package router

import (
	"github.com/gin-gonic/gin"
	"zph/controller"
)

func initVisitCommentRouter(group *gin.RouterGroup) {
	c := controller.NewCommentController()

	group.GET("/comments", c.GetComments)
	group.POST("/comment", c.CreateNewComment)
}

func initAdminCommentRouter(group *gin.RouterGroup) {
	c := controller.NewCommentController()

	group.GET("/comments", c.AdminGetComments)
	group.PUT("/comment/:comment_id/notice", c.AdminUpdateNotice)
	group.PUT("/comment/:comment_id/published", c.AdminUpdatePublished)
	group.PUT("/comment/:comment_id", c.AdminUpdateComment)
	group.DELETE("/comment/:comment_id", c.DeleteCommentById)
}
