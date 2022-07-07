package controller

import (
	"github.com/gin-gonic/gin"
	"time"
	"zph/constants"
	"zph/lib/cache"
	"zph/models/response"
	"zph/service"
	"zph/session"
)

type DashboardController struct {
	BaseController
	visitRecordService *service.VisitRecordService
	TagService *service.TagService
	categoryService *service.CategoryService
	blogService *service.BlogService
	cityVisitorService *service.CityVisitorService
}

func NewDashboardController() *DashboardController{
	return &DashboardController{
		visitRecordService: service.NewVisitRecordService(),
		TagService: service.NewTagService(),
		categoryService: service.NewCategoryService(),
		blogService: service.NewBlogService(),
		cityVisitorService: service.NewCityVisitorService(),
	}
}

func (controller *DashboardController) GetDashBoard(ctx *gin.Context) {
	userId := session.GetUserId(ctx)
	blogCount := controller.blogService.GetBlogCount(userId)

	category := response.Category{}
	categories := controller.categoryService.GetAllCategories(userId)
	categoryLegend := make([]string, 0)
	dashboardCatgories := make([]response.DashboardCategory, 0)
	for _, c := range categories {
		categoryLegend = append(categoryLegend, c.CategoryName)
		dashboardCatgories = append(dashboardCatgories, response.DashboardCategory{
			CategoryId: c.CategoryId,
			CategoryName: c.CategoryName,
			Value: controller.blogService.GetBlogCountByCategoryId(int(userId), c.CategoryId), // todo 查数量
		})
	}
	category.Legend = categoryLegend
	category.Series = dashboardCatgories

	today := time.Now().Format("20060102")
	todayPV := cache.GetTodayPV(userId, today)
	todayUV := cache.GetTodayUV(userId, today)

	var tag response.Tag
	dbTags := controller.TagService.GetAllTags(userId)
	for _, dbtag := range dbTags {
		tag.Legend = append(tag.Legend, dbtag.TagName)
		tag.Series = append(tag.Series, response.DashboardTag{
			TagId: dbtag.TagId,
			TagName: dbtag.TagName,
			Value: controller.blogService.GetBlogCountByTagId(dbtag.TagId),
		})
	}

	visitRecord := controller.visitRecordService.GetVisitRecords(userId)
	cityVisitor := controller.cityVisitorService.GetPVs(userId)

	resp := &response.DashboardResponse{
		BlogCount: blogCount,
		Category: category,
		CommentCount: 0,
		PV: todayPV,
		UV: todayUV,
		Tag: tag,
		VisitRecord: visitRecord,
		CityVisitor: cityVisitor,
	}
	ctx.Set(constants.OperationKey, constants.VisitDashboardPage.Value())
	controller.Success(ctx, resp)
}



