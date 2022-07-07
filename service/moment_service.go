package service

import (
	"zph/dao"
	api_error "zph/error"
	"zph/models/db"
	"zph/models/request"
	"zph/models/response"
	"zph/utils"
)

type MomentService struct {
	momentDao dao.MomentDao
}

func NewMomentService() *MomentService {
	return &MomentService{
		momentDao: dao.NewMomentDao(),
	}
}

func (service *MomentService) CreateMoment(req *request.MomentRequest, userId int)  {
	moment := &db.Moment{
		Content: req.Content,
		IsPublished: utils.BoolToInt(req.Published),
		CreateTime: utils.NowTime(),
		UpdateTime: utils.NowTime(),
		UserId: userId,
	}
	service.momentDao.CreateNewMoment(moment)
}

func (service *MomentService) GetMoments(offset, limit, userId int) response.GetMomentsResponse {
	total := service.momentDao.GetMomentCount(userId)
	moments := service.momentDao.GetMoments(offset, limit, userId)
	resp := response.GetMomentsResponse{
		Total: total,
	}
	momentResp := make([]response.MomentResponse, 0)
	for _, moment := range moments {
		momentResp = append(momentResp, response.MomentResponse{
			MomentId: moment.MomentId,
			Content: moment.Content,
			CreateTime: utils.TimeFormat(moment.CreateTime),
			UpdateTime: utils.TimeFormat(moment.UpdateTime),
			Likes: moment.Likes,
			IsPublished: utils.IntToBool(moment.IsPublished),
		})
	}
	resp.Moments = momentResp
	return resp
}

func (service *MomentService) UpdatePublishedStatus(userId int, req *request.UpdatePublishedRequest)  {
	params := map[string]interface{}{
		"is_published": req.IsPublished,
		"update_time": utils.NowTime(),
	}
	service.momentDao.UpdateMoment(req.MomentId, userId, params)
}

func (service *MomentService) GetMomentById(userId, momentId int) response.MomentResponse {
	moment := service.momentDao.GetMomentById(momentId, userId)
	if moment == nil {
		panic(api_error.NotFound)
	}
	resp := response.MomentResponse{
		MomentId: moment.MomentId,
		Content: moment.Content,
		CreateTime: utils.TimeFormat(moment.CreateTime),
		UpdateTime: utils.TimeFormat(moment.UpdateTime),
		Likes: moment.Likes,
		IsPublished: utils.IntToBool(moment.IsPublished),
	}
	return resp
}

func (service *MomentService) UpdateMoment(req *request.UpdateMomentRequest, userId int) {
	params := map[string]interface{}{
		"content": req.Content,
		"is_published": utils.BoolToInt(req.IsPublished),
		"update_time": utils.NowTime(),
	}
	service.momentDao.UpdateMoment(req.MomentId, userId, params)
}

func (service *MomentService) DeleteMomentById(momentId, userId int) {
	service.momentDao.DeleteMomentById(momentId, userId)
}