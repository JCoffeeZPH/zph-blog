package router

import (
	"github.com/gin-gonic/gin"
	"zph/controller"
)

func initCategoryRouter(group *gin.RouterGroup) {
	categoryController := controller.NewCategoryController()
	group.GET("/categories", categoryController.GetCategories)
	group.PUT("/category/:category_id", categoryController.UpdateCategory)
	group.POST("/category", categoryController.CreateNewCategory)
	group.DELETE("/category/:category_id", categoryController.DeleteCategoryById)
}
