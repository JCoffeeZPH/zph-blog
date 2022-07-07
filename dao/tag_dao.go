package dao

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	api_error "zph/error"
	"zph/global"
	"zph/lib/common"
	"zph/models/db"
	"zph/models/request"
)

type TagDao interface {
	GetTags(offset, limit, userId int)[]db.Tag
	GetTagCount(userId int) uint64
	UpdateTagById(tagId uint64, params map[string]string)
	CreateTag(tag *request.NewTagRequest, userId uint64) int
	DeleteTag(tagId uint64)
	GetAllTags(userId uint64)[]db.Tag
	GetTagsByIds(tagIds []int)[]db.Tag
}

type TagDaoImpl struct {
	client *gorm.DB
}

func NewTagDao() TagDao {
	return &TagDaoImpl{
		client: global.MysqlClient,
	}
}

func (dao *TagDaoImpl) GetTags(offset, limit, userId int)[]db.Tag  {
	var tags []db.Tag
	if err := dao.client.Where("user_id = ?", userId).Offset(offset).Limit(limit).Order("id asc").Find(&tags).Error; err != nil{
		log.Errorf("GetTags failed, offset is: %d, limit is: %d, userId is: %d err is: %+v",offset, limit, userId, err.Error())
		panic(common.DBError{Err: err})
	}
	return tags
}

func (dao *TagDaoImpl) GetTagCount(userId int) uint64 {
	var count uint64
	if err := dao.client.Model(&db.Tag{}).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		log.Errorf("GetTagCount failed, userId is: %d, err is: %+v", userId, err.Error())
		panic(common.DBError{Err: err})
	}
	return count
}

func (dao *TagDaoImpl) UpdateTagById(tagId uint64, params map[string]string)  {
	if err := dao.client.Model(&db.Tag{}).Where("id = ?", tagId).Update(params).Error; err != nil {
		log.Errorf("UpdateTagById failed, tagId = %d, params is: %+v, err is: %+v",tagId, params, err.Error())
		panic(common.DBError{Err: err})
	}
}

func (dao *TagDaoImpl) CreateTag(req *request.NewTagRequest, userId uint64) int{
	tag := &db.Tag{
		TagName: req.TagName,
		Color: req.Color,
		UserId: int(userId),
	}
	err := dao.client.Create(tag).Error
	if err == nil {
		return tag.TagId
	}
	driverErr, ok := err.(*mysql.MySQLError)
	if ok && driverErr.Number == 1062 {
		log.Errorf("CreateTag duplicate, params is: %+v, err is: %+v", req, err.Error())
		panic(api_error.AlreadyExists)
	}
	log.Errorf("CreateTag failed, params is: %+v, err is: %+v", req, err.Error())
	panic(common.DBError{Err: err})
}

func (dao *TagDaoImpl) DeleteTag(tagId uint64) {
	sql := "delete from tag_tab where id = ?"
	if err := dao.client.Exec(sql, []uint64{tagId}).Error; err != nil {
		log.Errorf("DeleteTag failed, tagId is: %d, err is: %+v", tagId, err.Error())
		panic(common.DBError{Err: err})
	}
}

func (dao *TagDaoImpl) GetAllTags(userId uint64)[]db.Tag  {
	var tags []db.Tag
	if err := dao.client.Where("user_id = ?", userId).Order("id asc").Find(&tags).Error; err != nil {
		log.Errorf("GetAllTags failed, userId is: %d err is: %+v", userId, err.Error())
		panic(common.DBError{Err: err})
	}
	return tags
}

func (dao *TagDaoImpl) GetTagsByIds(tagIds []int) []db.Tag  {
	var tags []db.Tag
	if err := dao.client.Where("id in (?)", tagIds).Find(&tags).Error; err != nil{
		log.Errorf("GetTagsByIds failed, tagIds: %+v, err is: %+v", tagIds, err.Error())
		panic(common.DBError{Err: err})
	}
	return tags
}
