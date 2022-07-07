package bo

type User struct {
	Id uint64 `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	Email string `json:"email"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	Role string `json:"role"`
}
