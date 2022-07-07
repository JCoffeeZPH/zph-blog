package db

type BlogUserId struct {
	Id uint64 `gorm:"id"`
	BlogId int `gorm:"blog_id"`
	UserId uint64 `gorm:"user_id"`
}

func (BlogUserId) TableName() string {
	return "blog_user_id_tab"
}
