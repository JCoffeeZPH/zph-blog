package db

type Comment struct {
	CommentId       int    `gorm:"column:id"`
	Nickname        string `gorm:"nickname"`
	Email           string `gorm:"email"`
	Content         string `gorm:"content"`
	Avatar          string `gorm:"avatar"`
	CreateTime      uint32 `gorm:"create_time"`
	Ip              string `gorm:"ip"`
	IsPublished     int    `gorm:"is_published"`
	IsAdminComment  int    `gorm:"is_admin_comment"`
	FromPage        int    `gorm:"from_page"`
	IsNotice        int    `gorm:"is_notice"`
	BlogId          uint64 `gorm:"blog_id"`
	ParentCommentId int    `gorm:"parent_comment_id"`
	Website         string `gorm:"website"`
	QQ              string `gorm:"qq"`
	UserId          uint64 `gorm:"user_id"`
}

func (Comment) TableName() string {
	return "comment_tab"
}

type PageComment struct {
	Id                    int    `gorm:"id"`
	Nickname              string `gorm:"nickname"`
	Content               string `gorm:"content"`
	Avatar                string `gorm:"avatar"`
	CreateTime            uint32 `gorm:"create_time"`
	Website               string `gorm:"website"`
	AdminComment          int    `gorm:"is_admin_comment"`
	ParentCommentId       int    `gorm:"parent_comment_id"`
	ParentCommentNickname string `gorm:"parent_comment_nickname"`
}
