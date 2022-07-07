package response

type GetTagsAndCategoriesResponse struct {
	Tags []TagResponse `json:"tags"`
	Categories []CategoryResponse `json:"categories"`
}

type BlogResponse struct {
	BlogId int `json:"id"`
	Title string `json:"title"`
	FirstPicture string `json:"first_picture"`
	Content string `json:"content"`
	Description string `json:"description"`
	IsPublished bool `json:"is_published"`
	IsRecommend bool `json:"is_recommend"`
	IsAppreciation bool `json:"is_appreciation"`
	IsCommentEnabled bool `json:"is_comment_enabled"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	ViewCount uint32 `json:"views"`
	WordCount int `json:"words"`
	ReadTime int `json:"read_time"`
	Category CategoryResponse `json:"category"`
	IsTop bool `json:"is_top"`
	Password string `json:"password"`
	UserId uint64 `json:"user_id"`
	Tags []TagResponse `json:"tags"`
}

type BlogsResponse struct {
	Total int `json:"total"`
	Blogs []BlogResponse `json:"blogs"`
	Categories []CategoryResponse `json:"categories"`
}