package response

type AboutMeResponse struct {
	CommentEnabled bool `json:"comment_enabled"`
	Content string `json:"content"`
	MusicId int `json:"music_id"`
	Title string `json:"title"`
}
