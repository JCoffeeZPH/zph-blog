package db

type OperationLog struct {
	OperationLogId int `gorm:"column:id"`
	Username string `gorm:"username"`
	Uri string `gorm:"uri"`
	Method string `gorm:"method"`
	Param string `gorm:"param"`
	Description string `gorm:"description"`
	Ip string `gorm:"ip"`
	IpSource string `gorm:"ip_source"`
	OS string `gorm:"os"`
	Browser string `gorm:"browser"`
	Times int64 `gorm:"times"`
	CreateTime uint32 `gorm:"create_time"`
	UserAgent string `gorm:"user_agent"`
	UserId uint64 `gorm:"user_id"`
}

func (OperationLog) TableName() string {
	return "operation_log_tab"
}
