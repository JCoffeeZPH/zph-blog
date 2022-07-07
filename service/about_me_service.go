package service

import (
	"zph/dao"
	"zph/models/request"
	"zph/models/response"
	"zph/utils"
)

type AboutMeService struct {
	aboutMeDao dao.AboutMeDao
}

func NewAboutMeService() *AboutMeService {
	return &AboutMeService{
		aboutMeDao: dao.NewAboutMeDao(),
	}
}

func (service *AboutMeService) GetAboutMe(userId uint64)response.AboutMeResponse  {
	dbAboutMe := service.aboutMeDao.GetAboutMe(userId)
	resp := response.AboutMeResponse{
		CommentEnabled: utils.IntToBool(dbAboutMe.CommentEnabled),
		Content: dbAboutMe.Content,
		MusicId: dbAboutMe.MusicId,
		Title: dbAboutMe.Title,
	}
	return resp
}

func (service *AboutMeService) UpdateAboutMe(userId int, req *request.UpdateAboutMeRequest) {
	params := map[string]interface{}{
		"title": req.Title,
		"music_id": req.MusicId,
		"content": req.Content,
		"comment_enabled": utils.BoolToInt(req.CommentEnabled),
	}
	service.aboutMeDao.UpdateAboutMe(userId, params)
}