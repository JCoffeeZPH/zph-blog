package response

type GetMomentsResponse struct {
	Total int64 `json:"total"`
	Moments []MomentResponse `json:"moments"`
}

type MomentResponse struct {
	MomentId int `json:"moment_id"`
	Content string `json:"content"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	Likes int `json:"likes"`
	IsPublished bool `json:"is_published"`
}
