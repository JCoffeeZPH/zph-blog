package controller

import (
	"github.com/gin-gonic/gin"
	"zph/constants"
	api_error "zph/error"
	"zph/models/request"
	"zph/models/response"
	"zph/service"
	"zph/session"
)

type CategoryController struct {
	BaseController
	categoryService *service.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		categoryService: service.NewCategoryService(),
	}
}

func (controller *CategoryController)GetCategories(c *gin.Context)  {
	req, err := controller.parseCategoriesParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	offset := (req.Page - 1) * req.PerPage
	limit := req.PerPage

	userId := session.GetUserId(c)
	categories, total := controller.categoryService.GetCategories(offset, limit, int(userId))
	resCategories := make([]response.CategoryResponse, 0)
	for _, category := range categories {
		resCategories = append(resCategories, response.NewGetCategoryResponse(category))
	}
	resp := &response.GetCategoryResponse{
		Total: total,
		Categories: resCategories,
	}
	c.Set(constants.OperationKey, constants.GetCategories.Value())
	controller.Success(c, resp)
}

func (controller *CategoryController) UpdateCategory(c *gin.Context) {
	req, err := controller.parseUpdateCategoryParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	controller.categoryService.UpdateCategory(req.CategoryId, req.CategoryName)
	c.Set(constants.OperationKey, constants.UpdateCategory.Value())
	controller.NoContent(c)
}

func (controller *CategoryController)CreateNewCategory(c *gin.Context) {
	req, err := controller.parseCreateCategoryParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(c)
	controller.categoryService.CreateNewCategory(req.CategoryName, userId)
	c.Set(constants.OperationKey, constants.CreateCategory.Value())
	controller.NoContent(c)
}

func (controller *CategoryController) DeleteCategoryById(c *gin.Context)  {
	req := &request.DeleteCategoryById{}
	if err := c.ShouldBindUri(req); err != nil {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(c)
	controller.categoryService.DeleteCategoryById(userId, req.CategoryId)
	c.Set(constants.OperationKey, constants.DeleteCategory.Value())
	controller.NoContent(c)
}

func (controller *CategoryController)parseCategoriesParams(c *gin.Context)(*request.GetCategoryRequest, error)  {
	req := &request.GetCategoryRequest{}
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

func (controller *CategoryController) parseUpdateCategoryParams(c *gin.Context) (*request.UpdateCategoryRequest, error) {
	req := &request.UpdateCategoryRequest{}
	c.ShouldBindUri(req)
	if err := c.ShouldBindJSON(req); err != nil{
		return nil, err
	}
	return req, nil
}

func (controller *CategoryController) parseCreateCategoryParams(c *gin.Context)(*request.CategoryRequest, error)  {
	req := &request.CategoryRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		return nil, err
	}
	return req, nil
}
