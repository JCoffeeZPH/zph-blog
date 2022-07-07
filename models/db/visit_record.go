package db

type VisitRecord struct {
	Id int `gorm:"id"`
	PV int `gorm:"pv"`
	UV int `gorm:"uv"`
	Date string `gorm:"date"`
	UserId int `gorm:"user_id"`
}

func (VisitRecord) TableName() string {
	return "visit_record_tab"
}