package response

type LoginUserResponse struct {
	UserId uint64 `json:"user_id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	Email string `json:"email"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	Role string `json:"role"`
	Token string `json:"token"`
}
