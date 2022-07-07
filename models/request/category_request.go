package request

type CategoryRequest struct {
	CategoryName string `json:"category_name" binding:"required"`
}

type GetCategoryRequest struct {
	Page int `form:"page" binding:"-"`
	PerPage int `form:"per_page" binding:"-"`
}

type UpdateCategoryRequest struct {
	CategoryId uint64 `uri:"category_id"`
	CategoryName string `json:"category_name" binding:"required"`
}

type DeleteCategoryById struct {
	CategoryId uint64 `uri:"category_id"`
}