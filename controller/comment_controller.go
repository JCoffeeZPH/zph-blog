package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zph/constants"
	api_error "zph/error"
	"zph/models/request"
	"zph/service"
	"zph/session"
)

type CommentController struct {
	BaseController
	commentService *service.CommentService
}

func NewCommentController() *CommentController {
	return &CommentController{
		commentService: service.NewCommentService(),
	}
}

func (controller *CommentController) CreateNewComment(ctx *gin.Context) {
	req, err := controller.parseCreateCommentParams(ctx)
	if err != nil {
		panic(api_error.ParamError)
	}
	controller.commentService.CreateComment(req)
	controller.NoContent(ctx)
}

func (controller *CommentController) GetComments(ctx *gin.Context) {
	req, err := controller.parseGetCommentsParams(ctx)
	if err != nil {
		panic(api_error.ParamError)
	}
	limit := req.PerPage
	offset := (req.Page - 1) * req.PerPage
	resp := controller.commentService.GetCommentService(req, limit, offset)
	controller.Success(ctx, resp)
}

func (controller *CommentController) AdminGetComments(ctx *gin.Context) {
	req, err := controller.parseAdminGetCommentsParams(ctx)
	if err != nil {
		panic(api_error.ParamError)
	}

	limit := req.PerPage
	offset := (req.Page - 1) * req.PerPage
	userId := session.GetUserId(ctx)
	resp := controller.commentService.AdminGetCommentService(limit, offset, userId)
	controller.Success(ctx, resp)
}

func (controller *CommentController) AdminUpdateNotice(ctx *gin.Context) {
	req, err := controller.parseAdminUpdateNoticeParams(ctx)
	if err != nil {
		panic(api_error.ParamError)
	}
	controller.commentService.UpdateNoticeCommentById(req)
	controller.NoContent(ctx)
}

func (controller *CommentController) AdminUpdatePublished(ctx *gin.Context) {
	req, err := controller.parseAdminUpdatePublishedParams(ctx)
	if err != nil {
		panic(api_error.ParamError)
	}
	controller.commentService.UpdatePublishCommentById(req)
	controller.NoContent(ctx)
}

func (controller *CommentController) AdminUpdateComment(ctx *gin.Context) {
	req, err := controller.parseAdminUpdateCommentParams(ctx)
	if err != nil {
		panic(api_error.ParamError)
	}
	controller.commentService.UpdateCommentById(req)
	controller.NoContent(ctx)
}

func (controller *CommentController) DeleteCommentById(ctx *gin.Context) {
	req := &request.DeleteCommentRequest{}
	if err := ctx.ShouldBindUri(req); err != nil {
		log.Errorf("parseDeleteCommentByIdParams failed, err is: %+v", err)
		panic(api_error.ParamError)
	}
	userId := session.GetUserId(ctx)
	controller.commentService.DeleteCommentById(req.CommentId, userId)
	controller.NoContent(ctx)
}

func (controller *CommentController) parseAdminUpdateCommentParams(ctx *gin.Context) (*request.AdminUpdateCommentRequest, error) {
	req := &request.AdminUpdateCommentRequest{}
	ctx.ShouldBindUri(req)
	if err := ctx.ShouldBindJSON(req); err != nil {
		log.Errorf("parseAdminUpdateCommentParams failed, err is: %+v", err)
		return nil, err
	}
	return req, nil
}

func (controller *CommentController) parseAdminUpdateNoticeParams(ctx *gin.Context) (*request.NoticeRequest, error) {
	req := &request.NoticeRequest{}
	ctx.ShouldBindUri(req)
	if err := ctx.ShouldBindJSON(req); err != nil {
		log.Errorf("parseAdminUpdateNoticeParams failed, err is: %+v", err)
		return nil, err
	}

	return req, nil
}

func (controller *CommentController) parseAdminUpdatePublishedParams(ctx *gin.Context) (*request.PublishedRequest, error) {
	req := &request.PublishedRequest{}
	ctx.ShouldBindUri(req)
	if err := ctx.ShouldBindJSON(req); err != nil {
		log.Errorf("parseAdminUpdatePublishedParams failed, err is: %+v", err)
		return nil, err
	}

	return req, nil
}

func (controller *CommentController) parseAdminGetCommentsParams(ctx *gin.Context) (*request.AdminGetCommentsRequest, error) {
	req := &request.AdminGetCommentsRequest{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		log.Errorf("parseAdminGetCommentsParams failed, err is: %+v", err)
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

func (controller *CommentController) parseCreateCommentParams(ctx *gin.Context) (*request.NewCommentRequest, error) {
	req := &request.NewCommentRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		log.Errorf("parseCreateCommentParams failed, err is: %+v")
		return nil, err
	}
	return req, nil
}

func (controller *CommentController) parseGetCommentsParams(ctx *gin.Context) (*request.GetCommentRequest, error) {
	req := &request.GetCommentRequest{}
	if err := ctx.ShouldBindQuery(req); err != nil {
		log.Errorf("parseGetCommentsParams failed, err is: %+v", err)
		return nil, err
	}
	if _, exist := constants.FromPageMap[int8(req.FromPage)]; !exist {
		log.Error("parseGetCommentsParams failed, err is: from page is valid")
		return nil, fmt.Errorf("from page is valid, param is: %+v", req)
	}
	if constants.FromPage(req.FromPage) == constants.Blog && req.BlogId <= 0 {
		log.Error("parseGetCommentsParams failed, err is: from page is blog, but no blogId")
		return nil, fmt.Errorf("from page is blog, but no blogId, param is: %+v", req)
	}
	if req.Page == 0 {
		req.Page = constants.DefaultPage
	}
	if req.PerPage == 0 {
		req.PerPage = constants.DefaultPerPage
	}
	return req, nil

}
