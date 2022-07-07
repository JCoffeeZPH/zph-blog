package request

type UpdateAboutMeRequest struct {
	Title string `json:"title" binding:"required"`
	MusicId int `json:"music_id" binding:"required"`
	Content string `json:"content" binding:"required"`
	CommentEnabled bool `json:"comment_enabled" binding:"-"`
}
