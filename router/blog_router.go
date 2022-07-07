package router

import (
	"github.com/gin-gonic/gin"
	"zph/controller"
)

func initBlogRouter(group *gin.RouterGroup) {
	blogController := controller.NewBlogController()
	group.GET("/tags_categories", blogController.GetTagsAndCategories)
	group.POST("/blog", blogController.CreateNewBlog)
	group.GET("/blog/:blog_id", blogController.GetBlogById)
	group.GET("/blogs", blogController.GetBlogs)
	group.PUT("/blog/:blog_id", blogController.UpdateBlog)
	group.DELETE("/blog/:blog_id", blogController.DeleteBlogById)
	group.PUT("/blog/top/:blog_id", blogController.UpdateTopStatus)
	group.PUT("/blog/recommend/:blog_id", blogController.UpdateRecommendStatus)
	group.PUT("/blog/:blog_id/visibility", blogController.UpdateVisibilityStatus)
}

func initIndexBlogRouter(group *gin.RouterGroup) {
	c := controller.NewBlogController()

	group.GET("/blog/:blog_id", c.VisitorReadBlogController)

}
