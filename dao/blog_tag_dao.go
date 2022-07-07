package dao

import (
	"github.com/jinzhu/gorm"
	"zph/global"
	"zph/lib/common"
	"zph/models/db"
)

type BlogTagDao interface {
	GetTagIdsByBlogId(blogId int) []int
	InsertIntoTagBlog(blogId int, tagIds []int)
	DeleteBlogIdTags(blogId int)
	GetBlogCountByTagId(tagId int) int
	DeleteBlogIdsByTagId(tagId uint64)
}

type BlogTagImpl struct {
	client *gorm.DB
}

func NewBlogTagDao() BlogTagDao {
	return &BlogTagImpl{
		client: global.MysqlClient,
	}
}

func (dao *BlogTagImpl) InsertIntoTagBlog(blogId int, tagIds []int) {
	sql := "insert into blog_tag (`blog_id`, `tag_id`) values (?, ?)"
	for _, tagId := range tagIds {
		params := []interface{}{blogId, tagId}
		err := dao.client.Debug().Exec(sql, params...).Error
		if err != nil {
			log.Errorf("InsertIntoTagBlog params is: %+v, err is: %+v", params, err)
			panic(common.DBError{Err: err})
		}
	}
}

func (dao *BlogTagImpl) GetTagIdsByBlogId(blogId int) []int {
	var tagIds []TagId
	if err := dao.client.Table("blog_tag").Select("tag_id").Where("blog_id = ?", blogId).Find(&tagIds).Error; err != nil {
		log.Errorf("GetTagIdsByBlogId failed, blogId is: %d, err is: %+v", blogId, err)
		panic(common.DBError{Err: err})
	}
	ids := make([]int, 0)
	for _, id := range tagIds {
		ids = append(ids, id.TagId)
	}
	return ids
}

func (dao *BlogTagImpl) DeleteBlogIdTags(blogId int) {
	var blogTags db.BlogTag
	if err := dao.client.Table("blog_tag").Where("blog_id = ?", blogId).Delete(blogTags).Error; err != nil {
		log.Errorf("DeleteBlogIdTags failed, blogId is: %d, err is: %+v", blogId, err)
		panic(common.DBError{Err: err})
	}
}

type TagId struct {
	TagId int `gorm:"column:tag_id"`
}

func (dao *BlogTagImpl) GetBlogCountByTagId(tagId int) int  {
	var count int
	if err := dao.client.Model(&db.BlogTag{}).Where("tag_id = ?", tagId).Count(&count).Error; err != nil {
		log.Errorf("GetBlogCountByTagId failed, tagId is: %d, err is: %+v", tagId, err)
		panic(common.DBError{Err: err})
	}
	return count
}

func (dao *BlogTagImpl) DeleteBlogIdsByTagId(tagId uint64){
	var blogTags db.BlogTag
	if err := dao.client.Table("blog_tag").Where("tag_id = ?", tagId).Delete(blogTags).Error; err != nil {
		log.Errorf("DeleteBlogIdsByTagId failed, tagId is: %d, err is: %+v", tagId, err)
		panic(common.DBError{Err: err})
	}
}
