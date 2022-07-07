package db

type LoginLog struct {
	LoginLogId int `gorm:"column:id"`
	Username string `gorm:"username"`
	IP string `gorm:"ip"`
	IPSource string `gorm:"ip_source"`
	OS string `gorm:"os"`
	Browser string `gorm:"browser"`
	Status int `gorm:"status"`
	Description string `gorm:"description"`
	CreateTime uint32 `gorm:"create_time"`
	UserAgent string `gorm:"user_agent"`
	UserId int `gorm:"user_id"`
}

func (LoginLog) TableName() string {
	return "login_log_tab"
}