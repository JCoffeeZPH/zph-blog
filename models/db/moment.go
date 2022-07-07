package db

type Moment struct {
	MomentId int `gorm:"column:id"`
	Content string `gorm:"content"`
	Likes int `gorm:"likes"`
	IsPublished int `gorm:"is_published"`
	CreateTime uint32 `gorm:"create_time"`
	UpdateTime uint32 `gorm:"update_time"`
	UserId int `gorm:"user_id"`
}

func (Moment) TableName() string {
	return "moment_tab"
}