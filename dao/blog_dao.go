package dao

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	api_error "zph/error"
	"zph/global"
	"zph/lib/common"
	"zph/models/db"
	"zph/utils"
)

type BlogDao interface {
	CreateBlog(blog *db.Blog) int
	GetBlogById(blogId uint64) db.Blog
	GetBlogs(userId, offset, limit, categoryId int, title string) []db.Blog
	DeleteBlogById(blogId, userId uint64)
	GetBlogCount(userId int) int
	UpdateBlog(blogId int, params map[string]interface{})
	UpdateBlogTopStatus(blogId, userId int, isTop bool)
	UpdateBlogRecommendStatus(blogId, userId int, isRecommend bool)
	UpdateBlogVisibilityStatus(blogId, userId int, params map[string]interface{})
	GetAllBlogIds(userId uint64) []int64
	GetBlogsByIds(ids []int64, userId uint64) []db.Blog
	GetBlogCountByUserIdCategoryId(userId, categoryId int) int
	IncrBlogViewCount(userId, blogId, target int)
	DeleteAllBlogsByCategoryId(userId, categoryId uint64)
	BatchGetBlogs(blogIds []uint64) []db.Blog
}

type BlogDaoImpl struct {
	client *gorm.DB
}

func NewBlogDao() BlogDao {
	return &BlogDaoImpl{
		client: global.MysqlClient,
	}
}

func (dao *BlogDaoImpl) CreateBlog(blog *db.Blog) int {
	err := dao.client.Create(blog).Error
	if err == nil {
		return blog.BlogId
	}
	driverErr, ok := err.(*mysql.MySQLError)
	if ok && driverErr.Number == 1062 {
		log.Errorf("CreateBlog duplicate, params is: %+v, err is: %+v", blog, err.Error())
		panic(api_error.AlreadyExists)
	}
	log.Errorf("CreateBlog failed, params is: %+v, err is: %+v", blog, err.Error())
	panic(common.DBError{Err: err})
}

func (dao *BlogDaoImpl) GetBlogById(blogId uint64) db.Blog {
	var blog db.Blog
	err := dao.client.Where("id = ?", blogId).First(&blog).Error
	if err != nil {
		log.Errorf("GetBlogById failed, blogId: %d, err is: %+v", blogId, err.Error())
		panic(common.DBError{Err: err})
	}
	return blog
}

func (dao *BlogDaoImpl) GetBlogs(userId, offset, limit, categoryId int, title string) []db.Blog {
	var blogs []db.Blog
	querySQL := dao.client.Offset(offset).Limit(limit).Where("user_id = ?", userId)
	if categoryId > 0 {
		querySQL = querySQL.Where("category_id = ?", categoryId)
	}
	if title != "" && len(title) > 0 {
		querySQL = querySQL.Where("title like ?", "%"+title+"%")
	}
	if err := querySQL.Order("create_time desc").Find(&blogs).Error; err != nil {
		log.Errorf("GetBlogs failed, userId is: %d, offset: %d, limit: %d, err is: %+v", userId, offset, limit, err.Error())
		panic(common.DBError{Err: err})
	}
	return blogs
}

func (dao *BlogDaoImpl) DeleteBlogById(blogId, userId uint64) {
	var blog db.Blog
	if err := dao.client.Where("id = ? and user_id = ?", blogId, userId).Delete(&blog).Error; err != nil {
		log.Errorf("DeleteBlogById failed, blogId is: %d, userId: %d, err is: %+v", blogId, userId, err.Error())
		panic(common.DBError{Err: err})
	}
}

func (dao *BlogDaoImpl) GetBlogCount(userId int) int {
	var count int
	if err := dao.client.Model(&db.Blog{}).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		log.Errorf("GetBlogCount failed, userId is: %d, err is: %+v", userId, err.Error())
		panic(common.DBError{Err: err})
	}
	return count
}

func (dao *BlogDaoImpl) UpdateBlog(blogId int, params map[string]interface{}) {
	if err := dao.client.Model(&db.Blog{}).Where("id = ?", blogId).Update(params).Error; err != nil {
		log.Errorf("UpdateBlog failed, blogId: %d, params: %+v, err is: %+v", blogId, params, err.Error())
		panic(common.DBError{Err: err})
	}
}

func (dao *BlogDaoImpl) UpdateBlogTopStatus(blogId, userId int, isTop bool) {
	if err := dao.client.Model(&db.Blog{}).Where("id = ? and user_id = ?", blogId, userId).Update(map[string]interface{}{"is_top": utils.BoolToInt(isTop), "update_time": utils.NowTime()}).Error; err != nil {
		log.Errorf("UpdateBlogTopStatus failed, blogId: %d, userId: %d, isTop: %+v, err is: %+v", blogId, userId, isTop, err.Error())
		panic(common.DBError{Err: err})
	}
}

func (dao *BlogDaoImpl) UpdateBlogRecommendStatus(blogId, userId int, isRecommend bool) {
	if err := dao.client.Model(&db.Blog{}).Where("id = ? and user_id = ?", blogId, userId).Update(map[string]interface{}{"is_recommend": utils.BoolToInt(isRecommend), "update_time": utils.NowTime()}).Error; err != nil {
		log.Errorf("UpdateBlogRecommendStatus failed, blogId: %d, userId: %d, is_recommend: %+v, err is: %+v", blogId, userId, isRecommend, err.Error())
		panic(common.DBError{Err: err})
	}
}

func (dao *BlogDaoImpl) UpdateBlogVisibilityStatus(blogId, userId int, params map[string]interface{}) {
	if err := dao.client.Model(&db.Blog{}).Where("id = ? and user_id = ?", blogId, userId).Update(params).Error; err != nil {
		log.Errorf("UpdateBlogVisibilityStatus failed, blogId: %d, userId: %d, params: %+v, err is: %+v", blogId, userId, params, err.Error())
		panic(common.DBError{Err: err})
	}
}

func (dao *BlogDaoImpl) GetAllBlogIds(userId uint64) []int64 {
	var idModels []BlogIdModel
	d := dao.client.Table("blog_tab").Select("id").Where("user_id = ? and is_published = ?", userId, 1).Find(&idModels)
	if err := d.Error; err != nil {
		log.Errorf("GetAllBlogIds failed, userId is: %d, err is: %+v", userId, err)
		panic(common.DBError{Err: err})
	}
	ids := make([]int64, 0)
	for _, model := range idModels {
		ids = append(ids, model.Id)
	}
	return ids
}

func (dao *BlogDaoImpl) GetBlogsByIds(ids []int64, userId uint64) []db.Blog {
	var blogs []db.Blog
	err := dao.client.Where("user_id = ? and id in (?)", userId, ids).Order("create_time desc").Find(&blogs).Error
	if err != nil {
		log.Errorf("GetBlogsByIds failed, userId is: %d, ids is: %+v, err is: %+v", userId, ids, err)
		panic(common.DBError{Err: err})
	}
	return blogs
}

type BlogIdModel struct {
	Id int64 `json:"id"`
}

func (dao *BlogDaoImpl) GetBlogCountByUserIdCategoryId(userId, categoryId int) int {
	var count int
	if err := dao.client.Model(&db.Blog{}).Where("user_id  = ? and category_id = ?", userId, categoryId).Count(&count).Error; err != nil {
		log.Errorf("GetBlogCountByUserIdCategoryId failed, userId is: %d, categoryId is: %d, err is: %+v", userId, categoryId, err)
		panic(common.DBError{Err: err})
	}
	return count
}

func (dao *BlogDaoImpl) IncrBlogViewCount(userId, blogId, target int) {
	if err := dao.client.Model(&db.Blog{}).Where("user_id = ? and id = ?", userId, blogId).Update(map[string]interface{}{"views": target}).Error; err != nil {
		log.Errorf("IncrBlogViewCount failed, userId is: %d, blogId is: %d, target is: %d, err is: %+v", userId, blogId, target, err)
		panic(common.DBError{Err: err})
	}
}

func (dao *BlogDaoImpl) DeleteAllBlogsByCategoryId(userId, categoryId uint64) {
	var blog db.Blog
	if err := dao.client.Where("category_id = ? and user_id = ?", categoryId, userId).Delete(&blog).Error; err != nil {
		log.Errorf("DeleteAllBlogsByCategoryId failed, categoryId is: %d, userId: %d, err is: %+v", categoryId, userId, err.Error())
		panic(common.DBError{Err: err})
	}
}

func (dao *BlogDaoImpl) BatchGetBlogs(blogIds []uint64) []db.Blog {
	var blogs []db.Blog
	if err := dao.client.Where("id in (?)", blogIds).Find(&blogs).Error; err != nil {
		panic(common.DBError{Err: err})
	}
	return blogs
}
