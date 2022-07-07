package controller

import (
	"github.com/gin-gonic/gin"
	"zph/constants"
	api_error "zph/error"
	"zph/models/request"
	"zph/models/response"
	"zph/service"
	"zph/session"
	"zph/utils"
)

type BlogController struct {
	BaseController
	tagService      *service.TagService
	categoryService *service.CategoryService
	blogService     *service.BlogService
}

func NewBlogController() *BlogController {
	return &BlogController{
		tagService:      service.NewTagService(),
		categoryService: service.NewCategoryService(),
		blogService:     service.NewBlogService(),
	}
}

func (controller *BlogController) GetTagsAndCategories(c *gin.Context) {
	userId := session.GetUserId(c)
	categories := controller.categoryService.GetAllCategories(userId)
	tags := controller.tagService.GetAllTags(userId)

	categoriesResp := make([]response.CategoryResponse, 0)
	for _, category := range categories {
		categoriesResp = append(categoriesResp, response.NewGetCategoryResponse(category))
	}

	tagsResp := make([]response.TagResponse, 0)
	for _, tag := range tags {
		tagsResp = append(tagsResp, response.NewTagResponse(tag))
	}

	resp := &response.GetTagsAndCategoriesResponse{
		Categories: categoriesResp,
		Tags:       tagsResp,
	}
	c.Set(constants.OperationKey, constants.GetTagsAndCategories.Value())
	controller.Success(c, resp)
}

func (controller *BlogController) CreateNewBlog(c *gin.Context) {
	req, err := controller.parseCreateBlogParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(c)
	controller.blogService.CreateNewBlog(userId, req)
	c.Set(constants.OperationKey, constants.CreateNewBlog.Value())
	controller.NoContent(c)
}

func (controller *BlogController) parseCreateBlogParams(c *gin.Context) (*request.BlogRequest, error) {
	req := &request.BlogRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (controller *BlogController) GetBlogById(c *gin.Context) {
	req := &request.GetBlogByIdRequest{}
	err := c.ShouldBindUri(req)
	if err != nil {
		panic(api_error.ParamError)
	}
	resp := controller.blogService.GetBlogById(req.BlogId)
	c.Set(constants.OperationKey, constants.GetBlogById.Value())
	controller.Success(c, resp)
}

func (controller *BlogController) GetBlogs(c *gin.Context) {
	req, err := controller.parseGetBlogsParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	offset := (req.Page - 1) * 10
	limit := req.PerPage

	userId := session.GetUserId(c)
	blogs, total := controller.blogService.GetBlogs(int(userId), offset, limit, req.CategoryId, req.Title)
	categories := controller.categoryService.GetAllCategories(userId)
	categoriesResp := make([]response.CategoryResponse, 0)
	for _, category := range categories {
		categoriesResp = append(categoriesResp, response.CategoryResponse{
			CategoryId:   category.CategoryId,
			CategoryName: category.CategoryName,
		})
	}
	resp := &response.BlogsResponse{
		Total:      total,
		Blogs:      blogs,
		Categories: categoriesResp,
	}
	c.Set(constants.OperationKey, constants.GetBlogs.Value())
	controller.Success(c, resp)
}

func (controller *BlogController) UpdateBlog(c *gin.Context) {
	req, err := controller.parseUpdateBlogParams(c)
	if err != nil || req.BlogId == 0 {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(c)
	controller.blogService.UpdateBlog(req, userId)
	c.Set(constants.OperationKey, constants.UpdateBlog.Value())
	controller.NoContent(c)

}

func (controller *BlogController) DeleteBlogById(c *gin.Context) {
	req := &request.GetBlogByIdRequest{}
	err := c.ShouldBindUri(req)
	if err != nil {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(c)
	controller.blogService.DeleteBlogById(req.BlogId, userId)
	c.Set(constants.OperationKey, constants.DeleteBlog.Value())
	controller.NoContent(c)
}

func (controller *BlogController) UpdateTopStatus(c *gin.Context) {
	req, err := controller.parseUpdateTopParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(c)
	controller.blogService.UpdateBlogTop(req.BlogId, int(userId), req.IsTop)
	c.Set(constants.OperationKey, constants.UpdateBlogTopById.Value())
	controller.NoContent(c)
}

func (controller *BlogController) UpdateRecommendStatus(c *gin.Context) {
	req, err := controller.parseUpdateRecommendParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(c)
	controller.blogService.UpdateBlogRecommend(req.BlogId, int(userId), req.IsRecommend)
	c.Set(constants.OperationKey, constants.UpdateBlogRecommendById.Value())
	controller.NoContent(c)
}

func (controller *BlogController) UpdateVisibilityStatus(c *gin.Context) {
	req, err := controller.parseUpdateVisibilityParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(c)
	controller.blogService.UpdateBlogVisibility(req.BlogId, int(userId), req)
	c.Set(constants.OperationKey, constants.UpdateVisibility.Value())
	controller.NoContent(c)
}

func (controller *BlogController) VisitorReadBlogController(ctx *gin.Context) {
	req := &request.VisitorReadBlogRequest{}
	err := ctx.ShouldBindUri(req)
	if err != nil {
		panic(api_error.ParamError)
	}
	if e := ctx.ShouldBindQuery(req); e != nil {
		panic(api_error.ParamError)
	}
	v, _ := ctx.Get(constants.VisitorBlogOwnerId)
	userId := v.(uint64)
	blog := controller.blogService.GetBlogById(req.BlogId)
	blog.Content = utils.MarkdownToHtml(blog.Content)
	blog.Description = utils.MarkdownToHtml(blog.Description)
	if req.Password != blog.Password {
		panic(api_error.Unauthorized)
	}
	blog.ViewCount = blog.ViewCount + 1
	controller.blogService.IncrViewCount(int(userId), int(req.BlogId), int(blog.ViewCount))
	controller.Success(ctx, blog)
}

func (controller *BlogController) parseGetBlogsParams(c *gin.Context) (*request.GetBlogsRequest, error) {
	req := &request.GetBlogsRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		return nil, err
	}
	if req.Page == 0 {
		req.Page = constants.DefaultPage
	}
	if req.PerPage == 0 {
		req.PerPage = constants.DefaultPerPage
	}
	return req, nil
}

func (controller *BlogController) parseUpdateBlogParams(c *gin.Context) (*request.UpdateBlogRequest, error) {
	req := &request.UpdateBlogRequest{}
	c.ShouldBindUri(req)
	if err := c.ShouldBindJSON(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (controller *BlogController) parseUpdateTopParams(c *gin.Context) (*request.TopRequest, error) {
	req := &request.TopRequest{}
	c.ShouldBindUri(req)
	if err := c.ShouldBindQuery(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (controller *BlogController) parseUpdateRecommendParams(c *gin.Context) (*request.RecommendRequest, error) {
	req := &request.RecommendRequest{}
	c.ShouldBindUri(req)
	if err := c.ShouldBindQuery(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (controller *BlogController) parseUpdateVisibilityParams(c *gin.Context) (*request.VisibilityRequest, error) {
	req := &request.VisibilityRequest{}
	c.ShouldBindUri(req)
	if err := c.ShouldBindJSON(req); err != nil {
		return nil, err
	}
	return req, nil
}
