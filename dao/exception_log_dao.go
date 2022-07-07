package dao

import (
	"github.com/jinzhu/gorm"
	"zph/global"
	"zph/models/db"
)

type ExceptionLogDao interface {

}

type ExceptionDaoImpl struct {
	client *gorm.DB
}

func NewExceptionDao() ExceptionLogDao {
	return &ExceptionDaoImpl{
		client: global.MysqlClient,
	}
}

func CreateLog(log db.ExceptionLog) {

}