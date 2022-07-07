package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	client2 "zph/lib/client"
	"zph/middleware"
	"zph/router/visit"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.Cors(), gin.Recovery())
	r.GET("/", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
	r.GET("/test", func(context *gin.Context) {
		client := client2.NewIPClient()
		res, _ := client.GetIPAndAttribution()
		context.JSON(http.StatusOK, gin.H{"res": res})
	})

	apiGroup := r.Group("/blog/api/v1", middleware.GlobalPanicHandling())

	adminLoginGroup := apiGroup.Group("/admin")
	initLoginRouter(adminLoginGroup)

	adminRouter := apiGroup.Group("/admin", middleware.AuthToken(), middleware.OperationLog())
	initDashboardRouter(adminRouter)
	initTagRouter(adminRouter)
	initCategoryRouter(adminRouter)
	initBlogRouter(adminRouter)
	initMomentRouter(adminRouter)
	initSiteSettingRouter(adminRouter)
	initFriendChainRouter(adminRouter)
	initAboutMeRouter(adminRouter)
	initAdminCommentRouter(adminRouter)

	logGroup := apiGroup.Group("/admin", middleware.AuthToken())
	initOperationLogRouter(logGroup)
	initLoginLogRouter(logGroup)

	//访问
	visitGroup := apiGroup.Group("/:username/visit", middleware.VisitorVisitBlogPage())
	visit.InitIndexRouter(visitGroup)
	initIndexBlogRouter(visitGroup)

	//g := visitGroup.Group("/", middleware.VisitorVisitBlogPage())
	//initIndexBlogRouter(g)

	//visitBlogGroup := apiGroup.Group("/visit", middleware.ReadBlogMiddleware())
	//visit.InitVisitBlogRouter(visitBlogGroup)

	// todo 访问日志
	//visitApiGroup := apiGroup.Group("/visit")

	// visit comment
	initVisitCommentRouter(apiGroup)

	return r
}
