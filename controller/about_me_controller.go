package controller

import (
	"github.com/gin-gonic/gin"
	"zph/constants"
	api_error "zph/error"
	"zph/models/request"
	"zph/service"
	"zph/session"
)

type AboutMeController struct {
	BaseController
	aboutMeService *service.AboutMeService
}

func NewAboutMeController() *AboutMeController {
	return &AboutMeController{
		aboutMeService: service.NewAboutMeService(),
	}
}

func (controller *AboutMeController) GetAboutMeInfo(ctx *gin.Context) {
	userId := session.GetUserId(ctx)
	resp := controller.aboutMeService.GetAboutMe(userId)
	ctx.Set(constants.OperationKey, constants.GetAboutMeInfo.Value())
	controller.Success(ctx, resp)
}

func (controller *AboutMeController) UpdateAboutMe (ctx *gin.Context) {
	req, err := controller.parseUpdateAboutMeParams(ctx)
	if err != nil {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(ctx)
	controller.aboutMeService.UpdateAboutMe(int(userId), req)
	ctx.Set(constants.OperationKey, constants.UpdateAboutMeInfo.Value())
	controller.NoContent(ctx)
}

func (controller *AboutMeController) parseUpdateAboutMeParams(ctx *gin.Context) (*request.UpdateAboutMeRequest, error) {
	req := &request.UpdateAboutMeRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		return nil, err
	}
	return req, nil
}