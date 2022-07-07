package db

type User struct {
	Id uint64 `gorm:"id"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Nickname string `gorm:"nickname"`
	Avatar string `gorm:"avatar"`
	Email string `gorm:"email"`
	CreateTime uint32 `gorm:"column:create_time"`
	UpdateTime uint32 `gorm:"column:update_time"`
	Role string `gorm:"role"`
}

func (user User) TableName()string {
	return "user_tab"
}
