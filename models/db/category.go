package db

type Category struct {
	CategoryId int `gorm:"column:id"`
	CategoryName string `gorm:"category_name"`
	UserId uint64 `gorm:"user_id"`
}

func (c Category) TableName()string {
	return "category_tab"
}
