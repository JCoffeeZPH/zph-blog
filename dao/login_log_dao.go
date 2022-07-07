package dao

import (
	"github.com/jinzhu/gorm"
	"zph/global"
	"zph/lib/common"
	"zph/models/db"
	"zph/utils"
)

type LoginLogDao interface {
	CreateLoginLog(operationLog *db.LoginLog)
	GetLoginLogs(startTime, endTime uint32, offset, limit, userId int) []db.LoginLog
	DeleteLoginLogByLogId(logId int)
	GetLoginLogCount(startTime, endTime uint32, userId int)uint64
}

type LoginLogDaoImpl struct {
	client *gorm.DB
}

func NewLoginLogDao() LoginLogDao {
	return &LoginLogDaoImpl{
		client: global.MysqlClient,
	}
}

func (dao *LoginLogDaoImpl) CreateLoginLog(loginLog *db.LoginLog) {
	if err := dao.client.Create(loginLog).Error; err != nil {
		log.Errorf("CreateLoginLog failed, loginLog is: %+v, err is: %+v", loginLog, err.Error())
		panic(common.DBError{Err: err})
	}
}

func (dao *LoginLogDaoImpl) GetLoginLogs(startTime, endTime uint32, offset, limit, userId int) []db.LoginLog {
	var logs []db.LoginLog
	d := dao.client.Offset(offset).Limit(limit).Where("user_id = ?", userId)
	if startTime != 0 && endTime != 0 {
		d = d.Where("create_time > ? and create_time < ?", startTime, endTime)
	}
	err := d.Find(&logs).Error
	if err != nil {
		log.Errorf("GetLoginLogs failed, userId is: %d, startTime is: %s, endTime is: %s, offset is: %d, limit is: %d, err is: %+v", userId, utils.TimeFormat(startTime), utils.TimeFormat(endTime), offset, limit, err)
		panic(common.DBError{Err: err})
	}
	return logs
}

func (dao *LoginLogDaoImpl) DeleteLoginLogByLogId(logId int)  {
	if err := dao.client.Where("id = ?", logId).Delete(&db.LoginLog{}).Error; err != nil{
		log.Errorf("DeleteLoginLogByLogId failed, logId is: %d, err is: %+v", logId, err)
		panic(common.DBError{Err: err})
	}
}

func (dao *LoginLogDaoImpl) GetLoginLogCount(startTime, endTime uint32, userId int)uint64  {
	var count uint64
	d := dao.client.Model(&db.LoginLog{}).Where("user_id = ?", userId)
	if startTime != 0 && endTime != 0 {
		d = d.Where("create_time > ? and create_time < ?", startTime, endTime)
	}
	if err := d.Count(&count).Error; err != nil{
		log.Errorf("GetLoginLogCount failed, startTime is: %s, endTime is: %s, userId is: %d, err is: %+v", utils.TimeFormat(startTime), utils.TimeFormat(endTime), userId, err)
		panic(common.DBError{Err: err})
	}
	return count
}
