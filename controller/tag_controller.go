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

type TagController struct {
	BaseController
	tagService *service.TagService
}

func NewTagController() *TagController {
	return &TagController{
		tagService: service.NewTagService(),
	}
}

// GetTags 获取博客标签列表
func (controller *TagController)GetTags(c *gin.Context)  {
	req, err := controller.parseGetTagsParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	offset := (req.Page - 1) * req.PerPage
	limit := req.PerPage
	userId := session.GetUserId(c)
	resp, total := controller.tagService.GetTags(offset, limit, int(userId))

	respTags := make([]response.TagResponse, 0)
	for _, tag := range resp {
		respTags = append(respTags, response.NewTagResponse(tag))
	}
	res := &response.GetTagResponse{
		Total: total,
		Tags: respTags,
	}
	c.Set(constants.OperationKey, constants.GetTags.Value())
	controller.Success(c, res)
}

func (controller *TagController) UpdateTag(c *gin.Context)  {
	req, err := controller.parseUpdateTagParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	controller.tagService.UpdateTag(req)
	c.Set(constants.OperationKey, constants.UpdateTag.Value())
	controller.NoContent(c)
}

func (controller *TagController) CreateTag(c *gin.Context)  {
	req, err := controller.parseCreateTagParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(c)
	controller.tagService.CreateNewTag(req, userId)
	c.Set(constants.OperationKey, constants.CreateTag.Value())
	controller.NoContent(c)
}

func (controller *TagController)DeleteTagById(c *gin.Context)  {
	req, err := controller.parseDeleteTagParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}

	controller.tagService.DeleteTagById(req.TagId)
	c.Set(constants.OperationKey, constants.DeleteTag.Value())
	controller.NoContent(c)
}

func (controller *TagController)parseGetTagsParams(c *gin.Context) (*request.GetTagsRequest, error) {
	req := &request.GetTagsRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		log.Errorf("parseGetTagsParams failed, err is: %+v", err)
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

func (controller *TagController)parseUpdateTagParams(c *gin.Context)(*request.UpdateTagRequest, error)  {
	req := &request.UpdateTagRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (controller *TagController)parseCreateTagParams(c *gin.Context)(*request.NewTagRequest, error)  {
	req := &request.NewTagRequest{}
	if err := c.ShouldBindJSON(req); err != nil{
		return nil, err
	}
	return req, nil
}

func (controller *TagController)parseDeleteTagParams(c *gin.Context) (*request.DeleteTagRequest, error) {
	req := &request.DeleteTagRequest{}
	if err := c.ShouldBindUri(req); err != nil {
		return nil, err
	}
	return req, nil
}