package dao

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	api_error "zph/error"
	"zph/global"
	"zph/lib/common"
	"zph/models/db"
)

type AboutMeDao interface {
	CreateAboutMe(dbAbout db.AboutMe)
	GetAboutMe(userId uint64) db.AboutMe
	UpdateAboutMe(userId int, params map[string]interface{})
}

type AboutMeDaoImpl struct {
	client *gorm.DB
}

func NewAboutMeDao() AboutMeDao {
	return &AboutMeDaoImpl{
		client: global.MysqlClient,
	}
}

func (dao *AboutMeDaoImpl) CreateAboutMe(dbAbout db.AboutMe) {
	err := dao.client.Create(dbAbout).Error
	if err == nil {
		return
	}
	driverErr, ok := err.(*mysql.MySQLError)
	if ok && driverErr.Number == 1062 {
		log.Errorf("CreateAboutMe duplicate, params is: %+v, err is: %+v", dbAbout, err.Error())
		panic(api_error.AlreadyExists)
	}
	log.Errorf("CreateAboutMe failed, params is: %+v, err is: %+v", dbAbout, err.Error())
	panic(common.DBError{Err: err})
}

func (dao *AboutMeDaoImpl) GetAboutMe(userId uint64) db.AboutMe {
	var aboutMe db.AboutMe
	if err := dao.client.Where("user_id = ?", userId).First(&aboutMe).Error; err != nil {
		log.Errorf("GetAboutMe failed, userId is: %d, err is: %+v", userId, err)
		panic(common.DBError{Err: err})
	}
	return aboutMe
}

func (dao *AboutMeDaoImpl) UpdateAboutMe(userId int, params map[string]interface{}) {
	if err := dao.client.Model(&db.AboutMe{}).Where("user_id = ?", userId).Update(params).Error; err != nil {
		log.Errorf("UpdateSiteSettings failed, err is: %+v", err)
		panic(common.DBError{Err: err})
	}
}
