package visit

import (
	"github.com/gin-gonic/gin"
	"zph/controller"
)

func InitVisitBlogRouter(group *gin.RouterGroup) {
	c := controller.NewBlogController()

	group.GET("/blog/:blog_id", c.VisitorReadBlogController)
}
