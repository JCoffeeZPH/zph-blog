package response

type VisitBlogResponse struct {
	Blogs []GetVisitBlogResponse `json:"blogs"`
	TotalPage int `json:"total_page"`
}

type GetVisitBlogResponse struct {
	Id int `json:"id"`
	Category CategoryResponse `json:"category"`
	CreateTime string `json:"create_time"`
	Description string `json:"description"`
	Password string `json:"password"`
	Privacy bool `json:"privacy"`
	ReadTime int `json:"read_time"`
	Tags []TagResponse `json:"tags"`
	Title string `json:"title"`
	Top bool `json:"top"`
	Views uint32 `json:"views"`
	Words int `json:"words"`
	UserId uint64 `json:"user_id"`
}