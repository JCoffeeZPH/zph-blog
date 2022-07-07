package bo

type FriendChain struct {
	Id int `json:"id"`
	Nickname string `json:"nickname"`
	Website string `json:"website"`
	Description string `json:"description"`
	Avatar string `json:"avatar"`
	IsPublished bool `json:"is_published"`
	Views int `json:"views"`
	CreateTime string `json:"create_time"`
	UserId int `json:"user_id"`
}
