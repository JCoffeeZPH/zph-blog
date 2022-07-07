package controller

import (
	"github.com/gin-gonic/gin"
	"zph/constants"
	api_error "zph/error"
	"zph/models/request"
	"zph/service"
	"zph/session"
)

type SiteSettingController struct {
	BaseController
	siteSettingService *service.SiteSettingService
}

func NewSiteSettingController() *SiteSettingController {
	return &SiteSettingController{
		siteSettingService: service.NewSiteSettingService(),
	}
}

func (controller *SiteSettingController) GetSettings(ctx *gin.Context) {
	userId := session.GetUserId(ctx)
	resp := controller.siteSettingService.GetSiteSettings(int(userId))
	ctx.Set(constants.OperationKey, constants.GetSiteSettings.Value())
	controller.Success(ctx, resp)
}

func (controller *SiteSettingController) UpdateSiteSettings(ctx *gin.Context) {
	req, err := controller.parseUpdateSiteSettings(ctx)
	if err != nil {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(ctx)
	controller.siteSettingService.UpdateSiteSettings(req, userId)
	ctx.Set(constants.OperationKey, constants.UpdateSiteSettings.Value())
	controller.NoContent(ctx)
}

func (controller *SiteSettingController) parseUpdateSiteSettings(c *gin.Context) (*request.SiteSettingRequest, error) {
	req := &request.SiteSettingRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		return nil, err
	}
	return req, nil
}