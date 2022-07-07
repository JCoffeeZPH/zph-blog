package session

import (
	"context"
	"github.com/gin-gonic/gin"
)

func SetUserId(ctx *gin.Context, userId uint64) {
	ctx.Set(string(KeyUserId), userId)
}

func SetUsername(ctx *gin.Context, username string) {
	ctx.Set(string(KeyUsername), username)
}

func SetNickname(ctx *gin.Context, nickname string) {
	ctx.Set(string(KeyNickname), nickname)
}

func SetAvatar(ctx *gin.Context, avatar string) {
	ctx.Set(string(KeyAvatar), avatar)
}

func SetEmail(ctx *gin.Context, email string) {
	ctx.Set(string(KeyEmail), email)
}

func SetRole(ctx *gin.Context, role string) {
	ctx.Set(string(KeyRole), role)
}

func GetUserId(ctx context.Context) (userId uint64) {
	if c, ok := ctx.(*gin.Context); ok{
		userId = c.GetUint64(string(KeyUserId))
		return userId
	}else {
		panic("Unsupported context")
	}
	return
}

func GetUsername(ctx context.Context) (username string) {
	if c, ok := ctx.(*gin.Context); ok{
		username = c.GetString(string(KeyUsername))
		return username
	}else {
		panic("Unsupported context")
	}
	return
}