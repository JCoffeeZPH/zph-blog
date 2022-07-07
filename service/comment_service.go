package service

import (
	"net/http"
	"regexp"
	"strings"
	"zph/config"
	"zph/constants"
	"zph/dao"
	api_error "zph/error"
	"zph/lib/cache"
	"zph/lib/client"
	"zph/models/db"
	"zph/models/request"
	"zph/models/response"
	"zph/utils"
)

type CommentService struct {
	commentDao dao.CommentDao
	blogDao    dao.BlogDao
	userDao    dao.UserDao
	ipClient   client.IPClient
	qqClient   *client.QQHttpClient
}

func NewCommentService() *CommentService {
	return &CommentService{
		commentDao: dao.NewCommentDao(),
		blogDao:    dao.NewBlogDao(),
		userDao:    dao.NewUserDao(),
		ipClient:   client.NewIPClient(),
		qqClient:   client.NewQQHttpClient(),
	}
}

const REG string = "^[1-9][0-9]{4,10}$"

func (service *CommentService) CreateComment(req *request.NewCommentRequest) {
	userInfo := service.userDao.GetUserByUsername(req.Username)
	ipDetail, err := service.ipClient.GetIPAndAttribution()
	if err != nil {
		panic(api_error.NewApiStatusError(http.StatusInternalServerError, "get_ip_failed", ""))
	}

	dbComment := &db.Comment{
		BlogId:          req.BlogId,
		Content:         req.Content,
		Email:           req.Email,
		IsNotice:        utils.BoolToInt(req.Notice),
		FromPage:        req.FromPage,
		ParentCommentId: req.ParentCommentId,
		Website:         req.Website,
		UserId:          userInfo.Id,
		Ip:              ipDetail.IP,
		CreateTime:      utils.NowTime(),
	}

	avatar := config.GetDefaultAvatarUrl()
	isQQ, _ := regexp.MatchString(REG, req.Nickname)
	nickname := req.Nickname
	if isQQ {
		qqDetails, e := service.qqClient.GetQQImageCDNUrl(req.Nickname)
		if e != nil {
			log.Warnf("GetQQImageCDNUrl failed, qq is: %s, err is: %+v, qqImageUrl is: %+v", req.Nickname, e, qqDetails)
		} else {
			dbComment.QQ = nickname
			avatar = qqDetails.QQAvatarCDNUrl
			nickname = qqDetails.Nickname
		}
	}
	dbComment.Avatar = avatar
	dbComment.Nickname = nickname

	if strings.EqualFold(req.Email, userInfo.Email) {
		dbComment.IsAdminComment = 1
	}
	service.commentDao.CreateComment(dbComment)
}

func (service *CommentService) GetCommentService(req *request.GetCommentRequest, limit, offset int) response.CommentResponse {
	userInfo := service.userDao.GetUserByUsername(req.Username)
	allCount := service.commentDao.GetCommentCountByPublished(req.BlogId, constants.CommentAllStatus, int(userInfo.Id), req.FromPage)
	closedCount := service.commentDao.GetCommentCountByPublished(req.BlogId, constants.CommentUnPublished, int(userInfo.Id), req.FromPage)
	comments := service.getCommentsByCommentId(req.FromPage, req.BlogId, constants.CommentDefaultParentId, limit, offset, int(userInfo.Id))

	resp := response.CommentResponse{
		AllCommentCount:    allCount,
		ClosedCommentCount: closedCount,
		Comments:           comments,
	}
	return resp
}

func (service *CommentService) getCommentsByCommentId(fromPage, blogId, parentCommentId, limit, offset, userId int) []response.GetCommentResponse {
	parentComments := service.commentDao.GetComments(fromPage, blogId, parentCommentId, limit, offset, userId)
	commentResponse := service.convertCommentResponse(parentComments)
	resp := make([]response.GetCommentResponse, 0)
	for _, comment := range commentResponse {
		replyComments := service.getCommentsByCommentId(fromPage, blogId, comment.Id, 0, 0, userId)
		comment.ReplyComments = replyComments
		resp = append(resp, comment)
	}
	return resp
}

func (service *CommentService) convertCommentResponse(dbComment []db.PageComment) []response.GetCommentResponse {
	resp := make([]response.GetCommentResponse, 0)
	for _, comment := range dbComment {
		resp = append(resp, response.GetCommentResponse{
			Id:                    comment.Id,
			Nickname:              comment.Nickname,
			Content:               comment.Content,
			Avatar:                comment.Avatar,
			CreateTime:            utils.TimeFormat(comment.CreateTime),
			Website:               comment.Website,
			AdminComment:          utils.IntToBool(comment.AdminComment),
			ParentCommentNickname: comment.ParentCommentNickname,
			ParentCommentId:       comment.ParentCommentId,
		})
	}
	return resp
}

func (service *CommentService) AdminGetCommentService(limit, offset int, userId uint64) response.AdminCommentResponse {
	comments := service.getAdminCommentsByCommentId(constants.CommentDefaultParentId, limit, offset, userId)
	resp := response.AdminCommentResponse{
		Comments: comments,
	}
	return resp
}

func (service *CommentService) convertAdminCommentResponse(dbComments []db.Comment) []response.AdminGetCommentResponse {
	resp := make([]response.AdminGetCommentResponse, 0)
	for _, comment := range dbComments {
		blogId := comment.BlogId
		commentBlog := response.Blog{
			BlogId: blogId,
		}
		title, exist := cache.GetTitleByBlogId(blogId)
		if !exist {
			blog := service.blogDao.GetBlogById(blogId)
			commentBlog.BlogTitle = blog.Title
			go cache.SetTitleByBlogId(blogId, blog.Title)
		} else {
			commentBlog.BlogTitle = title
		}
		resp = append(resp, response.AdminGetCommentResponse{
			Id:              comment.CommentId,
			Nickname:        comment.Nickname,
			Content:         comment.Content,
			Avatar:          comment.Avatar,
			CreateTime:      utils.TimeFormat(comment.CreateTime),
			Email:           comment.Email,
			IP:              comment.Ip,
			FromPage:        constants.FromPage(comment.FromPage).ToString(),
			IsPublished:     utils.IntToBool(comment.IsPublished),
			QQ:              comment.QQ,
			IsNotice:        utils.IntToBool(comment.IsNotice),
			Website:         comment.Website,
			AdminComment:    utils.IntToBool(comment.IsAdminComment),
			ParentCommentId: comment.ParentCommentId,
			Blog:            commentBlog,
		})
	}
	return resp
}

func (service *CommentService) getAdminCommentsByCommentId(parentCommentId, limit, offset int, userId uint64) []response.AdminGetCommentResponse {
	parentComments := service.commentDao.AdminGetComments(parentCommentId, limit, offset, userId)
	commentResponse := service.convertAdminCommentResponse(parentComments)
	resp := make([]response.AdminGetCommentResponse, 0)
	for _, comment := range commentResponse {
		replyComments := service.getAdminCommentsByCommentId(comment.Id, 0, 0, userId)
		comment.ReplyComments = replyComments
		resp = append(resp, comment)
	}
	return resp
}

func (service *CommentService) UpdateNoticeCommentById(req *request.NoticeRequest) {
	params := map[string]interface{}{
		"is_notice": utils.BoolToInt(req.IsNotice),
	}
	service.commentDao.UpdateCommentById(req.CommentId, params)
}

func (service *CommentService) UpdatePublishCommentById(req *request.PublishedRequest) {
	params := map[string]interface{}{
		"is_published": utils.BoolToInt(req.IsPublished),
	}
	service.commentDao.UpdateCommentById(req.CommentId, params)
}

func (service *CommentService) UpdateCommentById(req *request.AdminUpdateCommentRequest) {
	params := map[string]interface{}{
		"nickname": req.Nickname,
		"avatar":   req.Avatar,
		"email":    req.Email,
		"website":  req.Website,
		"ip":       req.IP,
		"content":  req.Content,
	}
	service.commentDao.UpdateCommentById(req.CommentId, params)
}

func (service *CommentService) DeleteCommentById(commentId int, userId uint64) {
	toDelCommentIds := service.getCommentIdsForDel(commentId, userId)
	for _, toDelCommentId := range toDelCommentIds {
		service.commentDao.DeleteCommentById(toDelCommentId)
	}
}

func (service *CommentService) getCommentIdsForDel(commentId int, userId uint64) []int {
	comments := service.commentDao.AdminGetComments(commentId, 0, 0, userId)
	commentIds := make([]int, 0)
	for _, comment := range comments {
		commentIds = append(commentIds, comment.CommentId)
		childCommentIds := service.getCommentIdsForDel(comment.CommentId, userId)
		commentIds = append(commentIds, childCommentIds...)
	}

	return commentIds
}
