package request

type UserLoginRequest struct {
	Username string `json:"username" binding:"-"`
	Password string `json:"password" binding:"-"`
}
