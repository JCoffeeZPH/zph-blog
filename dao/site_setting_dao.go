package dao

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"zph/constants"
	"zph/global"
	"zph/lib/common"
	"zph/models/db"
	"zph/utils"
)

type SiteSettingDao interface {
	GetSiteSettings(userId, settingType int) []db.SiteSetting
	UpdateSiteSettings(settings []*db.SiteSetting, userId int)
	DeleteFavorite(ids []int, userId uint64)
	UpdateMyFriendChainInfo(userId int, content string)
	UpdateMyFriendChainPageCommentSwitch(userId int, status bool)
}

type SiteSettingDapImpl struct {
	client *gorm.DB
}

func NewSiteSettingDao() SiteSettingDao {
	return &SiteSettingDapImpl{
		client: global.MysqlClient,
	}
}

func (dao *SiteSettingDapImpl) GetSiteSettings(userId, settingType int) []db.SiteSetting {
	var settings []db.SiteSetting
	d := dao.client.Where("user_id = ? and type = ?", userId, settingType).Find(&settings)
	if err := d.Error; err != nil {
		log.Errorf("GetSiteSettings failed, userId is: %d, settingType is: %d, err is: %+v", userId, settingType, err)
		panic(common.DBError{Err: err})
	}
	return settings
}

func (dao *SiteSettingDapImpl) UpdateSiteSettings(settings []*db.SiteSetting, userId int) {
	for _, setting := range settings {
		if setting.UserId == 0 {
			setting.UserId = userId
		}
		if err := dao.client.Model(&db.SiteSetting{}).Save(setting).Error; err != nil {
			log.Errorf("UpdateSiteSettings failed, err is: %+v", err)
			panic(common.DBError{Err: err})
		}
	}
}

func (dao *SiteSettingDapImpl) DeleteFavorite(ids []int, userId uint64) {
	for _, id := range ids {
		if err := dao.client.Where("user_id = ? and id = ?", userId, id).Delete(&db.SiteSetting{}).Error; err != nil {
			log.Errorf("DeleteFavorite failed, id is: %d, err is: %+v", id, err)
			panic(common.DBError{Err: err})
		}
	}
}

func (dao *SiteSettingDapImpl) UpdateMyFriendChainInfo(userId int, content string) {
	if err := dao.client.Model(&db.SiteSetting{}).Where("user_id = ? and type = ? and name_en = ?", userId, constants.FriendChainSetting, "friendContent").
		Update(map[string]interface{}{"value": content}).Error; err != nil {
			log.Errorf("UpdateMyFriendChainInfo failed, userId is: %d, content is: %s, err is: %+v", userId, content, err)
			panic(common.DBError{Err: err})
	}
}

func (dao *SiteSettingDapImpl) UpdateMyFriendChainPageCommentSwitch(userId int, status bool) {
	if err := dao.client.Model(&db.SiteSetting{}).Where("user_id = ? and type = ? and name_en = ?", userId, constants.FriendChainSetting, "friendCommentEnabled").
		Update(map[string]interface{}{"value": strconv.Itoa(utils.BoolToInt(status))}).Error; err != nil {
		log.Errorf("UpdateMyFriendChainInfo failed, userId is: %d, status is: %+v, err is: %+v", userId, status, err)
		panic(common.DBError{Err: err})
	}
}