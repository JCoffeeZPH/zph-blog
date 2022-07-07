package response

import "zph/models/db"

type GetCategoryResponse struct {
	Total uint64 `json:"total"`
	Categories []CategoryResponse `json:"categories"`
}

type CategoryResponse struct {
	CategoryId int `json:"id"`
	CategoryName string `json:"category_name"`
}

func NewGetCategoryResponse(category db.Category) CategoryResponse {
	return CategoryResponse{
		CategoryId: category.CategoryId,
		CategoryName: category.CategoryName,
	}
}