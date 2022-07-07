package dao

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	api_error "zph/error"
	"zph/global"
	"zph/lib/common"
	"zph/models/db"
)

type VisitRecordDao interface {
	CreateNewVisitRecord(visitRecord db.VisitRecord)
	UpdateVisitRecord(userId uint64, params map[string]interface{})
	GetVisitRecord(userId uint64) []db.VisitRecord
}

type VisitRecordDaoImpl struct {
	client *gorm.DB
}

func NewVisitRecordDao() VisitRecordDao {
	return &VisitRecordDaoImpl{
		client: global.MysqlClient,
	}
}

func (dao *VisitRecordDaoImpl) CreateNewVisitRecord(visitRecord db.VisitRecord) {
	err := dao.client.Create(&visitRecord).Error
	if err == nil {
		return
	}
	driverErr, ok := err.(*mysql.MySQLError)
	if ok && driverErr.Number == 1062 {
		log.Errorf("CreateNewVisitRecord duplicate, params is: %+v, err is: %+v", visitRecord, err.Error())
		panic(api_error.AlreadyExists)
	}
	log.Errorf("CreateNewVisitRecord failed, params is: %+v, err is: %+v", visitRecord, err.Error())
	panic(common.DBError{Err: err})
}

// 暂时用不到
func (dao *VisitRecordDaoImpl) UpdateVisitRecord(userId uint64, params map[string]interface{})  {
	err := dao.client.Model(&db.VisitRecord{}).Where("user_id = ?", userId).Update(params).Error
	if err != nil {
		log.Errorf("UpdateVisitRecord failed, userId is: %d, params is: %+v, err is: %+v", userId, params, err)
		panic(common.DBError{Err: err})
	}
}

func (dao *VisitRecordDaoImpl) GetVisitRecord(userId uint64) []db.VisitRecord {
	var records []db.VisitRecord
	if err := dao.client.Where("user_id = ?", userId).Offset(0).Limit(30).Order("id desc").Find(&records).Error; err != nil{
		log.Errorf("UpdateVisitRecord failed, userId is: %d, err is: %+v", userId, err)
		panic(common.DBError{Err: err})
	}
	return records
}