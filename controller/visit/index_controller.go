package visit

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
	"zph/constants"
	"zph/controller"
	api_error "zph/error"
	"zph/models/db"
	"zph/models/request"
	"zph/models/response"
	"zph/service"
	"zph/utils"
)

type IndexController struct {
	controller.BaseController
	tagService         *service.TagService
	categoryService    *service.CategoryService
	siteSettingService *service.SiteSettingService
	blogService        *service.BlogService
	userService        service.UserService
}

func NewIndexController() *IndexController {
	return &IndexController{
		tagService:         service.NewTagService(),
		categoryService:    service.NewCategoryService(),
		siteSettingService: service.NewSiteSettingService(),
		blogService:        service.NewBlogService(),
		userService:        service.NewUserService(),
	}
}

func (controller *IndexController) GetIndexDetail(ctx *gin.Context) {
	v, _ := ctx.Get(constants.VisitorBlogOwnerId)
	userId := v.(uint64)
	badges := controller.siteSettingService.GetFooterSettings(int(userId))
	dataCards := controller.siteSettingService.GetDataCardSettings(int(userId))
	basicSettings := controller.siteSettingService.GetBasicSettings(int(userId))
	categoryList := controller.categoryService.GetAllCategories(userId)
	tagList := controller.tagService.GetAllTags(userId)
	latestBlogs, _ := controller.blogService.GetBlogs(int(userId), 0, 3, 0, "")
	randomBlogList := controller.blogService.GetRandomBlogList(userId)
	resp := controller.convertIndexResponse(badges, dataCards, basicSettings, categoryList, tagList, latestBlogs)
	resp.RandomBlogList = randomBlogList
	controller.Success(ctx, resp)
}

func (controller *IndexController) VisitGetBlog(ctx *gin.Context) {
	req := &request.VisitGetBlogRequest{}
	err := ctx.ShouldBindQuery(req)
	if err != nil {
		panic(api_error.ParamError)
	}
	var offset int
	if req.Page == 0 {
		offset = 0
	} else {
		offset = (req.Page - 1) * constants.VisitPerPage
	}
	v, _ := ctx.Get(constants.VisitorBlogOwnerId)
	userId := v.(uint64)

	resp, total := controller.blogService.GetBlogs(int(userId), offset, constants.VisitPerPage, 0, "")

	blogResp := &response.VisitBlogResponse{
		Blogs:     controller.convertVisitGetBlogResponse(resp),
		TotalPage: total / 5,
	}
	controller.Success(ctx, blogResp)
}

func (controller *IndexController) convertIndexResponse(badges []db.SiteSetting, dataCards, basicSettings []db.SiteSetting,
	categoryList []db.Category, tagList []db.Tag, latestBlogs []response.BlogResponse) response.IndexResponse {
	resp := response.IndexResponse{}
	for _, badgeSetting := range badges {
		value := badgeSetting.Value
		var badge response.BadgeSetting
		json.Unmarshal([]byte(value), &badge)
		resp.Badges = append(resp.Badges, badge)
	}
	for _, dataCardSetting := range dataCards {
		switch dataCardSetting.NameEn {
		case "avatar":
			resp.DataCardSetting.Avatar = dataCardSetting.Value
		case "bilibili":
			resp.DataCardSetting.BiliBili = dataCardSetting.Value
		case "email":
			resp.DataCardSetting.Email = dataCardSetting.Value
		case "favorite":
			var favorite response.Favorite
			json.Unmarshal([]byte(dataCardSetting.Value), &favorite)
			resp.DataCardSetting.Favorites = append(resp.DataCardSetting.Favorites, favorite)
		case "github":
			resp.DataCardSetting.Github = dataCardSetting.Value
		case "name":
			resp.DataCardSetting.Name = dataCardSetting.Value
		case "netease":
			resp.DataCardSetting.Netease = dataCardSetting.Value
		case "rollText":
			texts := strings.Split(dataCardSetting.Value, ";")
			resp.DataCardSetting.RollText = texts
		default:
			resp.DataCardSetting.QQ = dataCardSetting.Value
		}
	}

	for _, basicSetting := range basicSettings {
		value := basicSetting.Value
		switch basicSetting.NameEn {
		case "beian":
			resp.BasicSetting.BeiAn = value
		case "blogName":
			resp.BasicSetting.BlogName = value
		case "commentAdminFlag":
			resp.BasicSetting.CommentAdminFlag = value
		case "copyright":
			resp.BasicSetting.Copyright = value
		case "footImgTitle":
			resp.BasicSetting.FooterImgTitle = value
		case "footerImgUrl":
			resp.BasicSetting.FooterImgUrl = value
		case "reward":
			resp.BasicSetting.Reward = value
		default:
			resp.BasicSetting.WebTitleSuffix = value
		}
	}

	for _, category := range categoryList {
		resp.CategoryList = append(resp.CategoryList, response.CategoryResponse{
			CategoryId:   category.CategoryId,
			CategoryName: category.CategoryName,
		})
	}

	for _, tag := range tagList {
		resp.TagList = append(resp.TagList, response.TagResponse{
			TagId:   tag.TagId,
			TagName: tag.TagName,
			Color:   tag.Color,
		})
	}

	for _, latestBlog := range latestBlogs {
		resp.LatestBlogList = append(resp.LatestBlogList, response.LatestBlog{
			Id:       latestBlog.BlogId,
			Privacy:  latestBlog.IsPublished,
			Password: latestBlog.Password,
			Title:    latestBlog.Title,
		})
	}
	return resp
}

func (controller *IndexController) convertVisitGetBlogResponse(blogs []response.BlogResponse) []response.GetVisitBlogResponse {
	resp := make([]response.GetVisitBlogResponse, 0)
	for _, blog := range blogs {
		resp = append(resp, response.GetVisitBlogResponse{
			Id:          blog.BlogId,
			Category:    blog.Category,
			CreateTime:  blog.CreateTime,
			Description: utils.MarkdownToHtml(blog.Description),
			Password:    blog.Password,
			Privacy:     utils.GetPrivacy(blog.IsPublished),
			ReadTime:    blog.ReadTime,
			Tags:        blog.Tags,
			Title:       blog.Title,
			Top:         blog.IsTop,
			Views:       blog.ViewCount,
			Words:       blog.WordCount,
			UserId:      blog.UserId,
		})
	}
	return resp
}
