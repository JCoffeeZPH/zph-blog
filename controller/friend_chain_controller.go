package controller

import (
	"github.com/gin-gonic/gin"
	"zph/constants"
	api_error "zph/error"
	"zph/models/request"
	"zph/service"
	"zph/session"
)

type FriendChainController struct {
	BaseController
	siteSettingService *service.SiteSettingService
	friendChainService *service.FriendChainService
}

func NewFriendChainController() *FriendChainController  {
	return &FriendChainController{
		siteSettingService: service.NewSiteSettingService(),
		friendChainService: service.NewFriendChainService(),
	}
}

func (controller *FriendChainController) GetFriendChain(ctx *gin.Context) {
	userId := session.GetUserId(ctx)
	resp := controller.siteSettingService.GetFriendChainSettings(userId)
	ctx.Set(constants.OperationKey, constants.GetFriendChainPageInfo.Value())
	controller.Success(ctx, resp)
}

func (controller *FriendChainController) CreateNewFriendChain(ctx *gin.Context) {
	req, err := controller.parseCreateNewFriendChainParams(ctx)
	if err != nil {
		panic(api_error.ParamError)
	}

	userId := session.GetUserId(ctx)
	controller.friendChainService.CreateFriendChain(req, userId)
	ctx.Set(constants.OperationKey, constants.CreateNewFriendChain.Value())
	controller.NoContent(ctx)
}

func (controller *FriendChainController) UpdateFriendChain(ctx *gin.Context) {
	req, err := controller.parseUpdateFriendChainParams(ctx)
	if err != nil {
		panic(api_error.ParamError)
	}

	userId := session.GetUserId(ctx)
	controller.friendChainService.UpdateFriendChain(req, int(userId))
	ctx.Set(constants.OperationKey, constants.UpdateFriendChain.Value())
	controller.NoContent(ctx)
}

func (controller *FriendChainController) GetFriendChains(ctx *gin.Context) {
	req, err := controller.parseGetFriendChainsParams(ctx)
	if err != nil {
		panic(api_error.ParamError)
	}

	userId := session.GetUserId(ctx)
	offset := (req.Page - 1) * req.PerPage
	limit := req.PerPage
	resp := controller.friendChainService.GetFriendChains(offset, limit, int(userId))
	ctx.Set(constants.OperationKey, constants.GetFriendChains.Value())
	controller.Success(ctx, resp)
}

func (controller *FriendChainController) DeleteFriendChainById(ctx *gin.Context) {
	req := &request.DeleteFriendChainRequest{}
	err := ctx.ShouldBindUri(req)
	if err != nil{
		panic(api_error.ParamError)
	}
	controller.friendChainService.DeleteFriendChainById(req.FriendChainId)
	ctx.Set(constants.OperationKey, constants.DeleteFriendChain.Value())
	controller.NoContent(ctx)
}

func (controller *FriendChainController) UpdateFriendChainCommentStatus(ctx *gin.Context) {
	req := &request.UpdateCommentStatusRequest{}
	err := ctx.ShouldBindQuery(req)
	if err != nil {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(ctx)
	controller.siteSettingService.UpdateMyFriendChainPageCommentSwitch(int(userId), req.CommentEnabled)
	ctx.Set(constants.OperationKey, constants.UpdateCommentStatus.Value())
	controller.NoContent(ctx)
}

func (controller *FriendChainController) UpdateFriendChainPageInfo(ctx *gin.Context) {
	req := &request.UpdatePageInfoRequest{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(ctx)
	controller.siteSettingService.UpdateMyFriendChainInfo(int(userId), req.Content)
	ctx.Set(constants.OperationKey, constants.UpdatePageInfo.Value())
	controller.NoContent(ctx)
}

func (controller *FriendChainController) UpdateFriendChainIsPublished(ctx *gin.Context) {
	req := &request.UpdateFriendChainIsPublished{}
	ctx.ShouldBindUri(req)
	if err := ctx.ShouldBindQuery(req); err != nil{
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(ctx)
	controller.friendChainService.UpdateFriendChainPublishStatus(req, userId)
	ctx.Set(constants.OperationKey, constants.UpdateIsPublished.Value())
	controller.NoContent(ctx)
}

func (controller *FriendChainController) parseCreateNewFriendChainParams(ctx *gin.Context)(*request.CreateNewFriendChainRequest, error)  {
	req := &request.CreateNewFriendChainRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (controller *FriendChainController) parseUpdateFriendChainParams(ctx *gin.Context) (*request.UpdateFriendChainRequest, error) {
	req := &request.UpdateFriendChainRequest{}
	ctx.ShouldBindUri(req)
	if err := ctx.ShouldBindJSON(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (controller *FriendChainController) parseGetFriendChainsParams(ctx *gin.Context)(*request.GetFriendChainsRequest, error)  {
	req := &request.GetFriendChainsRequest{}
	if err := ctx.ShouldBindQuery(req); err != nil {
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
