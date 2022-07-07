package db

type CityVisitor struct {
	City string `gorm:"city"`
	PV int `gorm:"pv"`
	UserId int `gorm:"user_id"`
}

func (CityVisitor) TableName() string {
	return "city_visitor_tab"
}