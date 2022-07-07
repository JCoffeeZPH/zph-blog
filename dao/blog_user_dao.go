package dao

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	api_error "zph/error"
	"zph/global"
	"zph/lib/common"
	"zph/models/db"
)

type BlogUserDao interface {
	CreateNewRelation(m db.BlogUserId)
	DeleteRelationByBlogIdUserId(userId, blogId uint64)
}

type BlogUserDaoImpl struct {
	client *gorm.DB
}

func NewBlogUserDao() BlogUserDao {
	return &BlogUserDaoImpl{
		client: global.MysqlClient,
	}
}

func (dao *BlogUserDaoImpl) CreateNewRelation(m db.BlogUserId) {
	err := dao.client.Create(m).Error
	if err == nil {
		return
	}
	driverErr, ok := err.(*mysql.MySQLError)
	if ok && driverErr.Number == 1062 {
		log.Errorf("CreateNewRelation duplicate, BlogUserId is: %+v, err is: %+v", m, err.Error())
		panic(api_error.AlreadyExists)
	}
	log.Errorf("CreateNewRelation failed, BlogUserId is: %+v, err is: %+v", m, err.Error())
	panic(common.DBError{Err: err})
}

func (dao *BlogUserDaoImpl) DeleteRelationByBlogIdUserId(userId, blogId uint64)  {
	err := dao.client.Where("blog_id = ? and user_id = ?", blogId, userId).Delete(&db.BlogUserId{}).Error
	if err != nil {
		log.Errorf("DeleteRelationByBlogIdUserId failed, err is:%+v, userId is: %d, blogId is: %d", err, userId, blogId)
		panic(common.DBError{Err: err})
	}
}
