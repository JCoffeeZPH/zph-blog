package dao

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"zph/constants"
	api_error "zph/error"
	"zph/global"
	"zph/lib/common"
	"zph/models/db"
)

type CommentDao interface {
	CreateComment(comment *db.Comment)
	GetCommentsByFromPage(fromPage, blogId, limit, offset int) []db.Comment
	GetComments(fromPage, blogId, parentCommentId, limit, offset, userId int) []db.PageComment
	GetCommentCountByPublished(blogId, commentPublishStatus, userId, fromPage int) int
	AdminGetComments(parentCommentId, limit, offset int, userId uint64) []db.Comment
	UpdateCommentById(commentId uint64, params map[string]interface{})
	DeleteCommentById(commentId int)
}

type CommentDaoImpl struct {
	client *gorm.DB
}

func NewCommentDao() CommentDao {
	return &CommentDaoImpl{
		client: global.MysqlClient,
	}
}

func (dao *CommentDaoImpl) CreateComment(comment *db.Comment) {
	err := dao.client.Create(comment).Error
	if err == nil {
		return
	}
	driverErr, ok := err.(*mysql.MySQLError)
	if ok && driverErr.Number == 1062 {
		log.Errorf("CreateComment duplicate, comment is: %+v, err is: %+v", comment, err.Error())
		panic(api_error.AlreadyExists)
	}
	log.Errorf("CreateComment failed, comment is: %+v, err is: %+v", comment, err.Error())
	panic(common.DBError{Err: err})
}

func (dao *CommentDaoImpl) GetCommentsByFromPage(fromPage, blogId, limit, offset int) []db.Comment {
	comments := make([]db.Comment, 0)
	err := dao.client.Where("from_page = ? and blog_id = ?", fromPage, blogId).Limit(limit).Offset(offset).Find(&comments).Error
	if err != nil {
		log.Errorf("GetCommentsByFromPage failed, err is: %+v, fromPage is: %d, blogId is: %d, limit is: %d, offset is: %d", err, fromPage, blogId, limit, offset)
		panic(common.DBError{Err: err})
	}
	return comments
}

func (dao *CommentDaoImpl) GetComments(fromPage, blogId, parentCommentId, limit, offset, userId int) []db.PageComment {
	res := make([]db.PageComment, 0)

	sql := "select c1.id, c1.nickname, c1.content, c1.avatar, c1.create_time, c1.is_admin_comment, c1.website," +
		"c1.parent_comment_id as parent_comment_id, c2.nickname as parent_comment_nickname " +
		"from comment_tab as c1 left join comment_tab as c2 on c1.parent_comment_id = c2.id " +
		"where c1.from_page = ? " +
		"and c1.parent_comment_id = ? " +
		"and c1.is_published = 1 " +
		"and c1.user_id = ?"
	var d *gorm.DB
	if fromPage == 0 && blogId > 0 {
		sql += " and c1.blog_id = ?"
		d = dao.client.Raw(sql, fromPage, parentCommentId, blogId, userId)
	} else {
		d = dao.client.Raw(sql, fromPage, parentCommentId, userId)
	}
	if limit > 0 {
		d = d.Limit(limit).Offset(offset)
	}
	d = d.Order("c1.create_time desc").Scan(&res)
	err := d.Error
	if err != nil {
		log.Errorf("GetComments failed, err is: %+v, fromPage is: %d, blogId is: %d, parentCommentId is: %d, userId is: %d", err, fromPage, blogId, parentCommentId, userId)
		panic(common.DBError{Err: err})
	}
	return res
}

func (dao *CommentDaoImpl) GetCommentCountByPublished(blogId, commentPublishStatus, userId, fromPage int) int {
	var count int
	d := dao.client.Model(&db.Comment{}).Where("user_id = ? and from_page = ?", userId, fromPage)
	if blogId > 0 {
		d = d.Where("blog_id = ?", blogId)
	}
	if commentPublishStatus == constants.CommentUnPublished {
		d = d.Where("is_published = 0")
	} else if commentPublishStatus == constants.CommentPublished {
		d = d.Where("is_published = 1")
	}
	err := d.Count(&count).Error
	if err != nil {
		log.Errorf("GetCommentCountByPublished failed, err is: %v, blogId is: %d, commentPublishStatus is: %d", err, blogId, commentPublishStatus)
		panic(common.DBError{Err: err})
	}
	return count
}

func (dao *CommentDaoImpl) AdminGetComments(parentCommentId, limit, offset int, userId uint64) []db.Comment {
	res := make([]db.Comment, 0)
	sql := "select * from comment_tab as c1 left join comment_tab as c2 on c1.parent_comment_id = c2.id where c1.parent_comment_id = ? and c1.user_id = ?"
	d := dao.client.Raw(sql, parentCommentId, userId)
	if limit > 0 {
		d = d.Limit(limit).Offset(offset)
	}
	err := d.Order("c1.create_time desc").Scan(&res).Error
	if err != nil {
		log.Errorf("AdminGetComments failed, err is: %+v, parentCommentId is: %d, limit is: %d, offset is: %d", parentCommentId, limit, offset)
		panic(common.DBError{Err: err})
	}
	return res
}

func (dao *CommentDaoImpl) UpdateCommentById(commentId uint64, params map[string]interface{}) {
	err := dao.client.Model(&db.Comment{}).Where("id = ?", commentId).Update(params).Error
	if err != nil {
		log.Errorf("UpdateCommentNoticePublishedById failed, commentId is: %d, params is: %v, err is: %+v", commentId, params, err)
		panic(common.DBError{Err: err})
	}
}

func (dao *CommentDaoImpl) DeleteCommentById(commentId int) {
	err := dao.client.Model(&db.Comment{}).Where("id = ?", commentId)
	if err != nil {
		log.Errorf("DeleteCommentById failed, err is: %+v, commentId is: %d", err, commentId)
		panic(common.DBError{Err: err})
	}
}
