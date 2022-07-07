package db

type BlogTag struct {
	BlogId int `gorm:"blog_id"`
	TagId int `gorm:"tag_id"`
}

func (BlogTag) TableName() string {
	return "blog_tag"
}
