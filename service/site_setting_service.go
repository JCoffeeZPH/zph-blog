package service

import (
	"strconv"
	"zph/constants"
	"zph/dao"
	"zph/lib/common"
	"zph/models/bo"
	"zph/models/db"
	"zph/models/request"
	"zph/models/response"
	"zph/utils"
)

type SiteSettingService struct {
	siteSettingDao dao.SiteSettingDao
}

func NewSiteSettingService() *SiteSettingService {
	return &SiteSettingService{
		siteSettingDao: dao.NewSiteSettingDao(),
	}
}

func (service *SiteSettingService) GetBasicSettings(userId int) []db.SiteSetting {
	settings := service.siteSettingDao.GetSiteSettings(userId, constants.BasicSetting)
	return settings
}

func (service *SiteSettingService) GetDataCardSettings(userId int) []db.SiteSetting {
	settings := service.siteSettingDao.GetSiteSettings(userId, constants.DataCardSetting)
	return settings
}

func (service *SiteSettingService) GetFooterSettings(userId int) []db.SiteSetting {
	settings := service.siteSettingDao.GetSiteSettings(userId, constants.FooterSetting)
	return settings
}

func (service *SiteSettingService) GetFriendChainSettings(userId uint64) response.FriendChainResponse {
	settings := service.siteSettingDao.GetSiteSettings(int(userId), constants.FriendChainSetting)
	resp := response.FriendChainResponse{

	}

	for _, setting := range settings {
		if setting.NameEn == "friendContent" {
			resp.Content = setting.Value
		}else if setting.NameEn == "friendCommentEnabled" {
			status, err := strconv.Atoi(setting.Value)
			if err != nil {
				panic(common.ServiceError{Err: err, API: "GetFriendChainSettings"})
			}
			resp.CommentEnabled = utils.IntToBool(status)
		}

	}
	return resp
}

func (service *SiteSettingService) GetSiteSettings(userId int) response.SiteSettingResponse {
	basicSettings := service.GetBasicSettings(userId)
	footerSettings := service.GetFooterSettings(userId)
	dataCardSettings := service.GetDataCardSettings(userId)

	resp := response.SiteSettingResponse{}

	for _, setting := range basicSettings {
		resp.BasicSettings = append(resp.BasicSettings, service.dbToResponseModel(setting))
	}

	for _, setting := range footerSettings {
		resp.FooterSettings = append(resp.FooterSettings, service.dbToResponseModel(setting))
	}

	for _, setting := range dataCardSettings {
		resp.DataCardSettings = append(resp.DataCardSettings, service.dbToResponseModel(setting))
	}
	return resp
}

func (service *SiteSettingService) UpdateSiteSettings(req *request.SiteSettingRequest, userId uint64)  {
	dbModels := make([]*db.SiteSetting, 0)
	for _, setting := range req.Settings {
		dbModels = append(dbModels, service.requestToDbModel(setting))
	}
	service.siteSettingDao.UpdateSiteSettings(dbModels, int(userId))
	service.siteSettingDao.DeleteFavorite(req.DeleteIds, userId)
}

func (service *SiteSettingService) dbToResponseModel(db db.SiteSetting) bo.SiteSettingModel {
	return bo.SiteSettingModel{
		Id: db.Id,
		NameEn: db.NameEn,
		NameZh: db.NameZh,
		Type: db.Type,
		Value: db.Value,
		UserId: db.UserId,
	}
}

func (service *SiteSettingService) requestToDbModel(request bo.SiteSettingModel) *db.SiteSetting {
	return &db.SiteSetting{
		Id: request.Id,
		NameEn: request.NameEn,
		NameZh: request.NameZh,
		Type: request.Type,
		Value: request.Value,
		UserId: request.UserId,
	}
}

func (service *SiteSettingService) UpdateMyFriendChainInfo(userId int, content string) {
	service.siteSettingDao.UpdateMyFriendChainInfo(userId, content)
}

func (service *SiteSettingService) UpdateMyFriendChainPageCommentSwitch(userId int, status bool) {
	service.siteSettingDao.UpdateMyFriendChainPageCommentSwitch(userId, status)
}