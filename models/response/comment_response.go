package response

type GetCommentResponse struct {
	Id                    int                  `json:"id"`
	Nickname              string               `json:"nickname"`
	Content               string               `json:"content"`
	Avatar                string               `json:"avatar"`
	CreateTime            string               `json:"create_time"`
	Website               string               `json:"website"`
	AdminComment          bool                 `json:"is_admin_comment"`
	ParentCommentId       int                  `json:"parent_comment_id"`
	ParentCommentNickname string               `json:"parent_comment_nickname"`
	ReplyComments         []GetCommentResponse `json:"reply_comments"`
}

type CommentResponse struct {
	AllCommentCount    int                  `json:"all_comment_count"`
	ClosedCommentCount int                  `json:"closed_comment_count"`
	Comments           []GetCommentResponse `json:"comments"`
}

type AdminGetCommentResponse struct {
	Id              int                       `json:"id"`
	Nickname        string                    `json:"nickname"`
	Content         string                    `json:"content"`
	Avatar          string                    `json:"avatar"`
	CreateTime      string                    `json:"create_time"`
	Email           string                    `json:"email"`
	IP              string                    `json:"ip"`
	FromPage        string                    `json:"from_page"`
	IsPublished     bool                      `json:"is_published"`
	QQ              string                    `json:"qq"`
	IsNotice        bool                      `json:"is_notice"`
	Website         string                    `json:"website"`
	AdminComment    bool                      `json:"is_admin_comment"`
	ParentCommentId int                       `json:"parent_comment_id"`
	ReplyComments   []AdminGetCommentResponse `json:"reply_comments"`
	Blog            Blog                      `json:"blog"`
}

type AdminCommentResponse struct {
	Comments []AdminGetCommentResponse `json:"comments"`
}

type Blog struct {
	BlogId    uint64 `json:"blog_id"`
	BlogTitle string `json:"blog_title"`
}
