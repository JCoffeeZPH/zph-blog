package dao

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	api_error "zph/error"
	"zph/global"
	"zph/lib/common"
	"zph/models/db"
)

type CategoryDao interface {
	GetCategories(offset, limit, userId int)[]db.Category
	GetCategoryCount(userId int) uint64
	UpdateCategoryById(categoryId uint64, params map[string]string)
	CreateCategory(categoryName string, userId uint64) int
	DeleteCategory(categoryId uint64)
	GetAllCategories(userId uint64)[]db.Category
	GetCategoryById(categoryId int)db.Category
}

type CategoryDaoImpl struct {
	client *gorm.DB
}

func NewCategoryDao() CategoryDao {
	return &CategoryDaoImpl{
		client: global.MysqlClient,
	}
}

func (dao *CategoryDaoImpl)GetCategories(offset, limit, userId int)[]db.Category  {
	var categories []db.Category
	if err := dao.client.Where("user_id = ?", userId).Offset(offset).Limit(limit).Order("id asc").Find(&categories).Error; err != nil{
		log.Errorf("GetCategories failed, offset is: %d, limit is: %d, userId is: %d, err is: %+v", offset, limit, userId, err.Error())
		panic(common.DBError{Err: err})
	}
	return categories
}

func (dao *CategoryDaoImpl) GetCategoryCount(userId int) uint64 {
	var count uint64
	if err := dao.client.Model(&db.Category{}).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		log.Errorf("GetCategoryCount failed, userId is: %d err is: %+v", userId, err.Error())
		panic(common.DBError{Err: err})
	}
	return count
}

func (dao *CategoryDaoImpl) UpdateCategoryById(categoryId uint64, params map[string]string)  {
	if err := dao.client.Model(&db.Category{}).Where("id = ?", categoryId).Update(params).Error; err != nil {
		log.Errorf("UpdateCategoryById failed, categoryId = %d, params is: %+v, err is: %+v",categoryId, params, err.Error())
		panic(common.DBError{Err: err})
	}
}

func (dao *CategoryDaoImpl) CreateCategory(categoryName string, userId uint64) int{
	category := &db.Category{
		CategoryName: categoryName,
		UserId: userId,
	}
	err := dao.client.Create(category).Error
	if err == nil {
		return category.CategoryId
	}
	driverErr, ok := err.(*mysql.MySQLError)
	if ok && driverErr.Number == 1062 {
		log.Errorf("CreateCategory duplicate, categoryName is: %s, userId is: %d, err is: %+v", categoryName, userId, err.Error())
		panic(api_error.AlreadyExists)
	}
	log.Errorf("CreateCategory failed, categoryName is: %s, userId is: %d, err is: %+v", categoryName, userId, err.Error())
	panic(common.DBError{Err: err})
}

func (dao *CategoryDaoImpl) DeleteCategory(categoryId uint64) {
	sql := "delete from category_tab where id = ?"
	if err := dao.client.Exec(sql, []uint64{categoryId}).Error; err != nil {
		log.Errorf("DeleteCategory failed, categoryId is: %d, err is: %+v", categoryId, err.Error())
		panic(common.DBError{Err: err})
	}
}

func (dao *CategoryDaoImpl)GetAllCategories(userId uint64)[]db.Category  {
	var categories []db.Category
	if err := dao.client.Where("user_id = ?", userId).Order("id asc").Find(&categories).Error; err != nil {
		log.Errorf("GetAllCategories failed, userId is: %d, err is: %+v", userId, err.Error())
		panic(common.DBError{Err: err})
	}
	return categories
}

func (dao *CategoryDaoImpl) GetCategoryById(categoryId int)db.Category{
	var category db.Category
	if err := dao.client.Where("id = ?", categoryId).First(&category).Error; err != nil {
		log.Errorf("GetCategoryById failed, categoryId: %d, err is: %+v", categoryId, err.Error())
		panic(common.DBError{Err: err})
	}
	return category
}

