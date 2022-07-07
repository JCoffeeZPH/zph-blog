package db

type Blog struct {
	BlogId int `gorm:"column:id"`
	Title string `gorm:"title"`
	FirstPicture string `gorm:"first_picture"`
	Content string `gorm:"content"`
	Description string `gorm:"description"`
	IsPublished int `gorm:"is_published"`
	IsRecommend int `gorm:"is_recommend"`
	IsAppreciation int `gorm:"is_appreciation"`
	IsCommentEnabled int `gorm:"is_comment_enabled"`
	CreateTime uint32 `gorm:"column:create_time"`
	UpdateTime uint32 `gorm:"column:update_time"`
	ViewCount uint32 `gorm:"column:views"`
	WordCount int `gorm:"column:words"`
	ReadTime int `gorm:"read_time"`
	CategoryId int `gorm:"category_id"`
	IsTop int `gorm:"is_top"`
	Password string `gorm:"password"`
	UserId uint64 `gorm:"user_id"`
}

func (Blog) TableName()string  {
	return "blog_tab"
}