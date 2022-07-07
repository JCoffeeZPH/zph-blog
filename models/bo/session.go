package bo

type Session struct {
	SessionId string `json:"session_id"`
	UserId uint64 `json:"user_id"`
	Username string `json:"username"`
	Secret string `json:"secret"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	Email string `json:"email"`
	Role string `json:"role"`
	UpdateTime uint32 `json:"update_time"`
	CreateTime uint32 `json:"create_time"`
}

