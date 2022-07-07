package response

import "zph/models/bo"

type FriendChainResponse struct {
	CommentEnabled bool `json:"comment_enabled"`
	Content string `json:"content"`
}

type GetFriendChainsResponse struct {
	Total int64 `json:"total"`
	FriendChains []bo.FriendChain `json:"friend_chains"`
}
