package request

type MomentRequest struct {
	Content string`json:"content" binding:"required"`
	Published bool `json:"is_published" binding:"required"`
}

type GetMomentsRequest struct {
	Page int `form:"page" binding:"gt=0"`
	PerPage int `form:"per_page" binding:"gt=0"`
}

type UpdatePublishedRequest struct {
	MomentId int `uri:"moment_id" binding:"required"`
	IsPublished bool `form:"is_published"`
}

type MomentByIdRequest struct {
	MomentId int `uri:"moment_id" binding:"required"`
}

type UpdateMomentRequest struct {
	MomentId int `uri:"moment_id" binding:"required"`
	Content string`json:"content" binding:"required"`
	IsPublished bool `json:"is_published" binding:"-"`
}