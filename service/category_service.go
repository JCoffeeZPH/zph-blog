package service

import (
	"zph/dao"
	"zph/models/db"
)

type CategoryService struct {
	categoryDao dao.CategoryDao
	blogDao dao.BlogDao
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		categoryDao: dao.NewCategoryDao(),
		blogDao: dao.NewBlogDao(),
	}
}

func (service *CategoryService) GetCategories(offset, limit, userId int)([]db.Category, uint64) {
	total := service.categoryDao.GetCategoryCount(userId)
	categories := service.categoryDao.GetCategories(offset, limit, userId)
	return categories, total
}

func (service *CategoryService) UpdateCategory(categoryId uint64, categoryName string)  {
	params := map[string]string{}
	params["category_name"] = categoryName
	service.categoryDao.UpdateCategoryById(categoryId, params)
}

func (service *CategoryService)CreateNewCategory(categoryName string, userId uint64)  {
	service.categoryDao.CreateCategory(categoryName, userId)
}

func (service *CategoryService) DeleteCategoryById(userId, categoryId uint64) {
	service.categoryDao.DeleteCategory(categoryId)
	service.blogDao.DeleteAllBlogsByCategoryId(userId, categoryId)
}

func (service *CategoryService)GetCategoriesCount(userId int)uint64  {
	return service.categoryDao.GetCategoryCount(userId)
}

func (service *CategoryService) GetAllCategories(userId uint64)[]db.Category {
	return service.categoryDao.GetAllCategories(userId)
}