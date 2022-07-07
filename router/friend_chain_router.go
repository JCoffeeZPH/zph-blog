package router

import (
	"github.com/gin-gonic/gin"
	"zph/controller"
)

func initFriendChainRouter(group *gin.RouterGroup) {
	c := controller.NewFriendChainController()

	group.GET("/friend_info", c.GetFriendChain)
	group.GET("/friend_chains", c.GetFriendChains)
	group.POST("/friend_chain", c.CreateNewFriendChain)
	group.PUT("/friend_chain/:friend_chain_id", c.UpdateFriendChain)
	group.DELETE("/friend_chain/:friend_chain_id", c.DeleteFriendChainById)
	group.PUT("/friend_chain/comment_enabled", c.UpdateFriendChainCommentStatus)
	group.PUT("/friend_chain/content", c.UpdateFriendChainPageInfo)
	group.PUT("/friend/is_published/:friend_chain_id", c.UpdateFriendChainIsPublished)

}
