package db

type ExceptionLog struct {
	Id uint64 `gorm:"id"`
	Uri string `gorm:"uri"`
	Method string `gorm:"method"`
	Errors string `gorm:"errors"`
	Ip string `gorm:"ip"`
	IpSource string `gorm:"ip_source"`
	OS string `gorm:"os"`
	Browser string `gorm:"browser"`
	CreateTime uint32 `gorm:"create_time"`
	UserAgent string `gorm:"user_agent"`
	UserId uint64 `gorm:"user_id"`
}
