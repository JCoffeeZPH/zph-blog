package visit

import (
	"github.com/gin-gonic/gin"
	"zph/controller/visit"
	"zph/middleware"
)

func InitIndexRouter(group *gin.RouterGroup) {
	c := visit.NewIndexController()

	group.GET("/site", c.GetIndexDetail)
	group.GET("/blogs", middleware.VisitorVisitBlogPage(), c.VisitGetBlog)
}
