package db

type SiteSetting struct {
	Id int `gorm:"id"`
	NameEn string `gorm:"name_en"`
	NameZh string `gorm:"name_zh"`
	Type int8 `gorm:"type"`
	Value string `json:"value"`
	UserId int `gorm:"user_id"`
}

func (SiteSetting) TableName() string {
	return "site_setting_tab"
}
