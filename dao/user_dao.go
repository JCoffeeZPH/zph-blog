package dao

import (
	"github.com/jinzhu/gorm"
	"zph/global"
	"zph/lib/common"
	"zph/logger"
	"zph/models/db"
)

var (
	log = logger.Log
)

type UserDao interface {
	GetUser(username, password string) *db.User
	GetUserByUsername(username string) db.User
	GetUserByUserId(userId uint64) db.User
}

type UserDaoImpl struct {
	client *gorm.DB
}

func NewUserDao() UserDao {
	return &UserDaoImpl{
		client: global.MysqlClient,
	}
}

// GetUser password是前端已经加密过的
func (dao *UserDaoImpl) GetUser(username, password string) *db.User {
	var user db.User
	if err := dao.client.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		log.Error("GetUserByUsername failed, err is: %+v", err)
		panic(common.DBError{Err: err})
	}
	return &user
}

func (dao *UserDaoImpl) GetUserByUsername(username string) db.User {
	var user db.User
	err := dao.client.Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Errorf("GetUserByUsername failed, username is: %s, err is: %+v", username, err)
		panic(common.DBError{Err: err})
	}
	return user
}

func (dao *UserDaoImpl) GetUserByUserId(userId uint64) db.User {
	var user db.User
	err := dao.client.Where("id = ?", userId).First(&user).Error
	if err != nil {
		log.Errorf("GetUserByUserId failed, userId is: %d, err is: %+v", userId, err)
		panic(common.DBError{Err: err})
	}
	return user
}
