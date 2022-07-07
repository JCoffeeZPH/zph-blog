package controller

import (
	"github.com/gin-gonic/gin"
	"zph/constants"
	api_error "zph/error"
	"zph/models/request"
	"zph/service"
	"zph/session"
)

type MomentController struct {
	BaseController
	momentService *service.MomentService
}

func NewMomentController() *MomentController {
	return &MomentController{
		momentService: service.NewMomentService(),
	}
}

func (controller *MomentController) CreateMoment(c *gin.Context)  {
	req, err := controller.parseCreateMomentParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(c)
	controller.momentService.CreateMoment(req, int(userId))
	c.Set(constants.OperationKey, constants.CreateMoment.Value())
	controller.NoContent(c)
}

func (controller *MomentController) GetMoments(c *gin.Context) {
	req, err := controller.parseGetMomentsParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	offset := (req.Page - 1) * req.PerPage
	limit := req.PerPage

	userId := session.GetUserId(c)
	resp := controller.momentService.GetMoments(offset, limit, int(userId))
	c.Set(constants.OperationKey, constants.GetMoments.Value())
	controller.Success(c, resp)
}

func (controller *MomentController) UpdatePublished(c *gin.Context) {
	req, err := controller.parseUpdateMomentParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(c)
	controller.momentService.UpdatePublishedStatus(int(userId), req)
	c.Set(constants.OperationKey, constants.UpdateMomentPublished.Value())
	controller.NoContent(c)
}

func (controller *MomentController) GetMomentById(c *gin.Context)  {
	req := &request.MomentByIdRequest{}
	err := c.ShouldBindUri(req)
	if err != nil {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(c)
	resp := controller.momentService.GetMomentById(int(userId), req.MomentId)
	c.Set(constants.OperationKey, constants.GetMomentById.Value())
	controller.Success(c, resp)
}

func (controller *MomentController) UpdateMoment(c *gin.Context)  {
	req, err := controller.parseUpdateMomentByIdParams(c)
	if err != nil {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(c)
	controller.momentService.UpdateMoment(req, int(userId))
	c.Set(constants.OperationKey, constants.UpdateMoment.Value())
	controller.NoContent(c)
}

func (controller *MomentController) DeleteMomentById(c *gin.Context) {
	req := &request.MomentByIdRequest{}
	err := c.ShouldBindUri(req)
	if err != nil {
		panic(api_error.ParamError)
	}

	userId := session.GetUserId(c)
	controller.momentService.DeleteMomentById(req.MomentId, int(userId))
	c.Set(constants.OperationKey, constants.DeleteMoment.Value())
	controller.NoContent(c)
}

func (controller *MomentController) parseCreateMomentParams(c *gin.Context)(*request.MomentRequest, error)  {
	req := &request.MomentRequest{}
	if err := c.ShouldBindJSON(req); err != nil{
		return nil, err
	}

	return req, nil
}

func (controller *MomentController) parseGetMomentsParams(c *gin.Context)(*request.GetMomentsRequest, error)  {
	req := &request.GetMomentsRequest{}
	if err := c.ShouldBindQuery(req); err != nil{
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

func (controller *MomentController) parseUpdateMomentParams(c *gin.Context) (*request.UpdatePublishedRequest, error) {
	req := &request.UpdatePublishedRequest{}
	c.ShouldBindUri(req)
	if err := c.ShouldBindQuery(req); err != nil{
		return nil, err
	}

	return req, nil
}

func (controller *MomentController) parseUpdateMomentByIdParams(c *gin.Context)(*request.UpdateMomentRequest, error)  {
	req := &request.UpdateMomentRequest{}
	c.ShouldBindUri(req)
	if err := c.ShouldBindJSON(req); err != nil {
		return nil, err
	}
	return req, nil
}