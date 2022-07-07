package db

type FriendChain struct {
	Id int `gorm:"id"`
	Nickname string `gorm:"nickname"`
	Website string `gorm:"website"`
	Description string `gorm:"description"`
	Avatar string `gorm:"avatar"`
	IsPublished int `gorm:"is_published"`
	Views int `gorm:"views"`
	CreateTime uint32 `gorm:"create_time"`
	UserId int `gorm:"user_id"`
}

func (FriendChain) TableName() string {
	return "friend_tab"
}