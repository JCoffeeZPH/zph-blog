package request

type GetTagsRequest struct {
	Page int `form:"page" binding:"-"`
	PerPage int `form:"per_page"binding:"-"`
}

type UpdateTagRequest struct {
	TagId uint64 `json:"id" binding:"required"`
	TagName string `json:"tag_name" binding:"-"`
	Color 	string `json:"color" binding:"-"`
}

type NewTagRequest struct {
	TagName string `json:"tag_name" binding:"required"`
	Color 	string `json:"color" binding:"required"`
}

type DeleteTagRequest struct {
	TagId uint64 `uri:"tag_id" binding:"required"`
}