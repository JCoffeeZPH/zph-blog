package request

type CreateNewFriendChainRequest struct {
	Nickname string `json:"nickname" binding:"required"`
	Description string `json:"description" binding:"required"`
	Website string `json:"website" binding:"required"`
	Avatar string `json:"avatar" binding:"required"`
	IsPublished bool `json:"is_published" binding:"-"`
}

type UpdateFriendChainRequest struct {
	FriendChainId int `uri:"friend_chain_id"`
	Nickname string `json:"nickname" binding:"required"`
	Description string `json:"description" binding:"required"`
	Website string `json:"website" binding:"required"`
	Avatar string `json:"avatar" binding:"required"`
	IsPublished bool `json:"is_published" binding:"-"`
}

type GetFriendChainsRequest struct {
	Page int `form:"page" binding:"gt=0"`
	PerPage int `form:"per_page" binding:"gt=0"`
}

type DeleteFriendChainRequest struct {
	FriendChainId int `uri:"friend_chain_id" binding:"required"`
}

type UpdateCommentStatusRequest struct {
	CommentEnabled bool `form:"comment_enabled"`
}

type UpdatePageInfoRequest struct {
	Content string `json:"content" binding:"required"`
}

type UpdateFriendChainIsPublished struct {
	FriendChainId int `uri:"friend_chain_id" binding:"required"`
	IsPublished bool `form:"is_published" binding:"-"`
}