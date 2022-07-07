package request

type BlogRequest struct {
	BlogId int `json:"blog_id" binding:"-"`
	Title string `json:"title" binding:"required"`
	FirstPicture string `json:"first_picture" binding:"required"`
	Content string `json:"content" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsPublished bool `json:"is_published" binding:"-"`
	IsRecommend bool `json:"is_recommend" binding:"-"`
	IsAppreciation bool `json:"is_appreciation" binding:"-"`
	IsCommentEnabled bool `json:"is_comment_enabled" binding:"-"`
	Category interface{} `json:"category" binding:"required"`
	IsTop bool `json:"is_top" binding:"-"`
	Password string `json:"password" binding:"-"`
	UserId uint64 `json:"user_id" binding:"-"`
	Tags []interface{}  `json:"tag_list" binding:"required"`
}

type GetBlogByIdRequest struct {
	BlogId uint64 `uri:"blog_id"`
}

type GetBlogsRequest struct {
	Title string `form:"title" binding:"-"` //搜索时调用
	CategoryId int `form:"category_id" binding:"-"`
	Page int `form:"page" binding:"gt=0"`
	PerPage int `form:"per_page" binding:"gt=0"`
}

type UpdateBlogRequest struct {
	BlogId int `uri:"blog_id" binding:"required"`
	Title string `json:"title" binding:"required"`
	FirstPicture string `json:"first_picture" binding:"required"`
	Content string `json:"content" binding:"required"`
	Description string `json:"description" binding:"-"`
	IsPublished bool `json:"is_published" binding:"-"`
	IsRecommend bool `json:"is_recommend" binding:"-"`
	IsAppreciation bool `json:"is_appreciation" binding:"-"`
	IsCommentEnabled bool `json:"is_comment_enabled" binding:"-"`
	Category interface{} `json:"category" binding:"required"`
	IsTop bool `json:"is_top" binding:"-"`
	Password string `json:"password" binding:"-"`
	Tags []interface{}  `json:"tag_list" binding:"required"`
}

type TopRequest struct {
	BlogId int `uri:"blog_id" binding:"required"`
	IsTop bool `form:"is_top" binding:"-"`
}

type RecommendRequest struct {
	BlogId int `uri:"blog_id" binding:"required"`
	IsRecommend bool `form:"is_recommend" binding:"-"`
}

type VisibilityRequest struct {
	BlogId int `uri:"blog_id" binding:"required"`
	IsTop bool `json:"is_top" binding:"-"`
	IsAppreciation bool `json:"is_appreciation" binding:"-"`
	IsCommentEnabled bool `json:"is_comment_enabled" binding:"-"`
	IsPublished bool `json:"is_published" binding:"-"`
	IsRecommend bool `json:"is_recommend" binding:"-"`
	Password string `json:"password" binding:"-"`
}

type VisitorReadBlogRequest struct {
	//Username string `uri:"username"`
	BlogId uint64 `uri:"blog_id"`
	Password string `json:"password"`
}
