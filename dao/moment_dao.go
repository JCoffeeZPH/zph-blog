package dao

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	api_error "zph/error"
	"zph/global"
	"zph/lib/common"
	"zph/models/db"
)

type MomentDao interface {
	CreateNewMoment (moment *db.Moment) int
	GetMoments(offset, limit, userId int)[]db.Moment
	GetMomentCount(userId int) int64
	UpdateMoment(momentId, userId int, params map[string]interface{})
	GetMomentById(momentId, userId int) *db.Moment
	DeleteMomentById(momentId, userId int)
}

type MomentDaoImpl struct {
	client *gorm.DB
}

func NewMomentDao() MomentDao {
	return &MomentDaoImpl{
		client: global.MysqlClient,
	}
}

func (dao *MomentDaoImpl) CreateNewMoment (moment *db.Moment) int {
	err := dao.client.Create(moment).Error
	if err == nil {
		return moment.MomentId
	}
	driverErr, ok := err.(*mysql.MySQLError)
	if ok && driverErr.Number == 1062 {
		log.Errorf("CreateNewMoment duplicate, params is: %+v, err is: %+v", moment, err)
		panic(api_error.AlreadyExists)
	}
	log.Errorf("CreateNewMoment failed, params is: %+v, err is: %+v", moment, err)
	panic(common.DBError{Err: err})
}

func (dao *MomentDaoImpl) GetMoments(offset, limit, userId int)[]db.Moment  {
	var moments []db.Moment
	if err := dao.client.Where("user_id = ?", userId).Offset(offset).Limit(limit).Order("id asc").Find(&moments).Error; err != nil {
		log.Errorf("GetMoments failed, offset is: %d, limit is: %d, err is: %+v", offset, limit, err)
		panic(common.DBError{Err: err})
	}
	return moments
}

func (dao *MomentDaoImpl) GetMomentCount(userId int) int64{
	var count int64
	if err := dao.client.Model(db.Moment{}).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		log.Errorf("GetMomentCount failed, err is: %+v", err)
		panic(common.DBError{Err: err})
	}
	return count
}

func (dao *MomentDaoImpl) UpdateMoment(momentId, userId int, params map[string]interface{})  {
	if err := dao.client.Model(&db.Moment{}).Where("id = ? and user_id = ?", momentId, userId).Update(params).Error; err != nil {
		log.Errorf("UpdatePublished failed, momentId is: %d, params: %+v, err is: %+v", err.Error)
		panic(common.DBError{Err: err})
	}
}

func (dao *MomentDaoImpl) GetMomentById(momentId, userId int) *db.Moment {
	var moment db.Moment
	d := dao.client.Where("id = ? and user_id = ?", momentId, userId).First(&moment)
	if d.RecordNotFound() {
		log.Debugf("GetMomentById not found")
		return nil
	}
	if d.Error != nil {
		log.Errorf("GetMomentById failed, momentId is: %d, userId is: %d, err is: %+v", momentId, userId, d.Error.Error())
		return nil
	}
	return &moment

}

func (dao *MomentDaoImpl) DeleteMomentById(momentId, userId int)  {
	if err := dao.client.Where("id = ? and user_id = ?", momentId, userId).Delete(&db.Moment{}).Error; err != nil {
		log.Errorf("DeleteMomentById failed, momentId is: %d, userId is: %d, err is: %+v", momentId, userId, err.Error())
		panic(common.DBError{Err: err})
	}
}