package db

type Tag struct {
	TagId int `gorm:"column:id"`
	TagName string `gorm:"tag_name"`
	Color string `gorm:"color"`
	UserId int `gorm:"user_id"`
}

func (t Tag) TableName() string {
	return "tag_tab"
}