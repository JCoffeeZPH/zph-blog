package request

type GetUserByUsernameRequest struct {
	Username string `uri:"username" binding:"required"`
}
