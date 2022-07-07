package dao

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	api_error "zph/error"
	"zph/global"
	"zph/lib/common"
	"zph/models/db"
	"zph/utils"
)

type FriendChainDao interface {
	CreateNewFriendChain(friend *db.FriendChain)
	GetFriendChainCount(userId int) int64
	GetFriendChains(offset, limit, userId int) []db.FriendChain
	DeleteFriendChain(chainId int)
	UpdateFriendChain(chainId, userId int, params interface{})
	UpdateFriendChainIsPublished(chainId, userId int, isPublished bool)
}

type FriendChainDaoImpl struct {
	client *gorm.DB
}

func NewFriendChainDao() FriendChainDao {
	return &FriendChainDaoImpl{
		client: global.MysqlClient,
	}
}

func (dao *FriendChainDaoImpl) CreateNewFriendChain(friend *db.FriendChain) {
	err := dao.client.Create(friend).Error
	if err == nil {
		return
	}
	driverErr, ok := err.(*mysql.MySQLError)
	if ok && driverErr.Number == 1062 {
		log.Errorf("CreateNewFriendChain duplicate, params is: %+v, err is: %+v", friend, err.Error())
		panic(api_error.AlreadyExists)
	}
	log.Errorf("CreateNewFriendChain failed, params is: %+v, err is: %+v", friend, err.Error())
	panic(common.DBError{Err: err})
}

func (dao *FriendChainDaoImpl) GetFriendChainCount(userId int) int64 {
	var count int64
	if err := dao.client.Model(&db.FriendChain{}).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		log.Errorf("GetFriendChainCount failed, userId is: %d", userId)
		panic(common.DBError{Err: err})
	}
	return count
}

func (dao *FriendChainDaoImpl) GetFriendChains(offset, limit, userId int) []db.FriendChain {
	var friendChains []db.FriendChain
	if err := dao.client.Where("user_id = ?", userId).Offset(offset).Limit(limit).Find(&friendChains).Error; err != nil {
		log.Errorf("GetFriendChains failed, userId is: %d, offset is: %d, limit is: %d, err is: %+v", userId, offset, limit, err)
		panic(common.DBError{Err: err})
	}
	return friendChains
}

func (dao *FriendChainDaoImpl) DeleteFriendChain(chainId int) {
	err := dao.client.Where("id = ?", chainId).Delete(&db.FriendChain{}).Error
	if err != nil {
		log.Errorf("DeleteFriendChain failed, chainId is: %d, err is: %+v", chainId, err)
		panic(common.DBError{Err: err})
	}
}

func (dao *FriendChainDaoImpl) UpdateFriendChain(chainId, userId int, params interface{}) {
	err := dao.client.Model(&db.FriendChain{}).Where("id = ? and user_id = ?", chainId, userId).Update(params).Error
	if err != nil {
		log.Errorf("UpdateFriendChain failed, chainId is: %d, userId is: %d, params is: %+v, err is: %+v", chainId, userId, params, err)
		panic(common.DBError{Err: err})
	}
}

func (dao *FriendChainDaoImpl) UpdateFriendChainIsPublished(chainId, userId int, isPublished bool) {
	if err := dao.client.Model(&db.FriendChain{}).Where("id = ? and user_id = ?", chainId, userId).
		Update(map[string]interface{}{"is_published": utils.BoolToInt(isPublished)}).Error; err != nil {
			log.Errorf("UpdateFriendChainIsPublished failed, chainId is: %d, userId is: %d, isPublished is: %+v, err is: %+v", chainId, userId, isPublished, err)
			panic(common.DBError{Err: err})
	}
}