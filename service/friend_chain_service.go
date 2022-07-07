package service

import (
	"zph/dao"
	"zph/models/bo"
	"zph/models/db"
	"zph/models/request"
	"zph/models/response"
	"zph/utils"
)

type FriendChainService struct {
	friendChainDao dao.FriendChainDao
}

func NewFriendChainService() *FriendChainService {
	return &FriendChainService{
		friendChainDao: dao.NewFriendChainDao(),
	}
}

func (service *FriendChainService) CreateFriendChain(req *request.CreateNewFriendChainRequest, userId uint64) {
	dbModel := &db.FriendChain{
		Nickname:    req.Nickname,
		Website:     req.Website,
		Description: req.Description,
		Avatar:      req.Avatar,
		IsPublished: utils.BoolToInt(req.IsPublished),
		CreateTime:  utils.NowTime(),
		UserId: int(userId),
	}
	service.friendChainDao.CreateNewFriendChain(dbModel)
}

func (service *FriendChainService) GetFriendChains(offset, limit, userId int) response.GetFriendChainsResponse {
	dbFriendChains := service.friendChainDao.GetFriendChains(offset, limit, userId)
	count := service.friendChainDao.GetFriendChainCount(userId)
	resp := response.GetFriendChainsResponse{}
	for _, friendChain := range dbFriendChains {
		resp.FriendChains = append(resp.FriendChains, bo.FriendChain{
			Id: friendChain.Id,
			Nickname: friendChain.Nickname,
			Website: friendChain.Website,
			Description: friendChain.Description,
			Avatar: friendChain.Avatar,
			IsPublished: utils.IntToBool(friendChain.IsPublished),
			Views: friendChain.Views,
			CreateTime: utils.TimeFormat(friendChain.CreateTime),
			UserId: friendChain.UserId,
		})
	}
	resp.Total = count
	return resp
}

func (service *FriendChainService) DeleteFriendChainById(chainId int)  {
	service.friendChainDao.DeleteFriendChain(chainId)
}

func (service *FriendChainService) UpdateFriendChain(req *request.UpdateFriendChainRequest, userId int) {
	params := map[string]interface{}{
		"nickname":     req.Nickname,
		"description":  req.Description,
		"website":      req.Website,
		"avatar":       req.Avatar,
		"is_published": utils.BoolToInt(req.IsPublished),
	}
	service.friendChainDao.UpdateFriendChain(req.FriendChainId, userId, params)
}

func (service *FriendChainService) UpdateFriendChainPublishStatus(req *request.UpdateFriendChainIsPublished, userId uint64) {
	service.friendChainDao.UpdateFriendChainIsPublished(req.FriendChainId, int(userId), req.IsPublished)
}


