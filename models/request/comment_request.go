package request

type NewCommentRequest struct {
	BlogId          uint64 `json:"blog_id" binding:"required"`
	Content         string `json:"content" binding:"required,not_empty_max=500"`
	Email           string `json:"email" binding:"required,empty_or_max=60"`
	Nickname        string `json:"nickname" binding:"required,empty_or_max=200"`
	Notice          bool   `json:"is_notice"`
	FromPage        int    `json:"from_page" binding:"required, in=1 2 3"` // 0普通文章, 1关于我, 2友链
	ParentCommentId int    `json:"parent_comment_id"`
	Website         string `json:"website" binding:"empty_or_max=100"`
	Username        string `json:"username" binding:"required"`
}

type GetCommentRequest struct {
	Page     int    `form:"page"`
	PerPage  int    `form:"per_page"`
	FromPage int    `form:"from_page" binding:"required"`
	BlogId   int    `form:"blog_id"`
	Username string `form:"username" binding:"required"`
}

type AdminGetCommentsRequest struct {
	Page    int `form:"page"`
	PerPage int `form:"per_page"`
}

type NoticeRequest struct {
	CommentId uint64 `uri:"comment_id"`
	IsNotice  bool   `json:"is_notice"`
}

type PublishedRequest struct {
	CommentId   uint64 `uri:"comment_id"`
	IsPublished bool   `json:"is_published"`
}

type AdminUpdateCommentRequest struct {
	CommentId uint64 `uri:"comment_id"`
	Nickname  string `json:"nickname" binding:"required,not_empty_max=200"`
	Avatar    string `json:"avatar" binding:"required,not_empty_max=100"`
	Email     string `json:"email" binding:"required,not_empty_max=60"`
	Website   string `json:"website"`
	IP        string `json:"ip" binding:"required,not_empty_max=16"`
	Content   string `json:"content" binding:"required,not_empty_max=500"`
}

type DeleteCommentRequest struct {
	CommentId int `uri:"comment_id"`
}
