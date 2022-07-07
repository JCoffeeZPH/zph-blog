package db

type AboutMe struct {
	Id int `gorm:"id"`
	Title string `gorm:"title"`
	MusicId int `gorm:"music_id"`
	CommentEnabled int `gorm:"comment_enabled"`
	Content string `gorm:"content"`
	UserId int `gorm:"user_id"`
}

func (AboutMe) TableName() string {
	return "about_me_tab"
}
