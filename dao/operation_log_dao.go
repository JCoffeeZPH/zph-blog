package dao

import (
	"github.com/jinzhu/gorm"
	"zph/global"
	"zph/lib/common"
	"zph/models/db"
	"zph/utils"
)

type OperationLogDao interface {
	CreateOperationLog(operationLog *db.OperationLog)
	GetOperations(startTime, endTime uint32, offset, limit, userId int) []db.OperationLog
	DeleteLogByLogId(logId int)
	GetOperationLogCount(startTime, endTime uint32, userId int)uint64
}

type OperationLogDaoImpl struct {
	client *gorm.DB
}

func NewOperationLogDao() OperationLogDao {
	return &OperationLogDaoImpl{
		client: global.MysqlClient,
	}
}

func (dao *OperationLogDaoImpl) CreateOperationLog(operationLog *db.OperationLog) {
	if err := dao.client.Create(operationLog).Error; err != nil {
		log.Errorf("CreateOperationLog failed, operationLog is: %+v, err is: %+v", operationLog, err.Error())
		panic(common.DBError{Err: err})
	}
}

func (dao *OperationLogDaoImpl) GetOperations(startTime, endTime uint32, offset, limit, userId int) []db.OperationLog {
	var logs []db.OperationLog
	d := dao.client.Offset(offset).Limit(limit).Where("user_id = ?", userId)
	if startTime != 0 && endTime != 0 {
		d = d.Where("create_time > ? and create_time < ?", startTime, endTime)
	}
	err := d.Find(&logs).Error
	if err != nil {
		log.Errorf("GetOperations failed, userId is: %d, startTime is: %s, endTime is: %s, offset is: %d, limit is: %d, err is: %+v", userId, utils.TimeFormat(startTime), utils.TimeFormat(endTime), offset, limit, err)
		panic(common.DBError{Err: err})
	}
	return logs
}

func (dao *OperationLogDaoImpl) DeleteLogByLogId(logId int)  {
	if err := dao.client.Where("id = ?", logId).Delete(&db.OperationLog{}).Error; err != nil{
		log.Errorf("DeleteLogByLogId failed, logId is: %d, err is: %+v", logId, err)
		panic(common.DBError{Err: err})
	}
}

func (dao *OperationLogDaoImpl) GetOperationLogCount(startTime, endTime uint32, userId int)uint64  {
	var count uint64
	d := dao.client.Model(&db.OperationLog{}).Where("user_id = ?", userId)
	if startTime != 0 && endTime != 0 {
		d = d.Where("create_time > ? and create_time < ?", startTime, endTime)
	}
	if err := d.Count(&count).Error; err != nil{
		log.Errorf("GetOperationLogCount failed, startTime is: %s, endTime is: %s, userId is: %d, err is: %+v", utils.TimeFormat(startTime), utils.TimeFormat(endTime), userId, err)
		panic(common.DBError{Err: err})
	}
	return count
}