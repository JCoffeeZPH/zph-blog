package middleware

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"zph/config"
	api_error "zph/error"
	"zph/lib/cache"
	"zph/models/bo"
	"zph/models/util"
	"zph/session"
)

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := getJwtToken(c)
		if len(jwtToken) == 0 {
			panic(api_error.Unauthorized)
		}
		s := util.ParseSession(jwtToken)
		if s == nil {
			panic(api_error.Unauthorized)
		}
		auth(c, jwtToken, s)
		c.Next()
	}
}


func getJwtToken(c *gin.Context) string {
	header := c.GetHeader("Authorization")
	splits := strings.Split(header, " ")
	if len(splits) == 2 && strings.ToLower(splits[0]) == "bearer" {
		return splits[1]
	} else {
		return ""
	}
}

func auth(ctx *gin.Context, jwtToken string, s *bo.Session) {
	if !verifyJwtToken(jwtToken, s.Secret) {
		panic(api_error.Unauthorized)
	}
	go cache.SaveSessionIntoRedis(s)
	setContext(ctx, s)
}

func verifyJwtToken(jwtToken, secret string) bool {
	md5Value := md5.Sum([]byte(secret))
	md5 := fmt.Sprintf("%x", md5Value)

	signingKey := md5 + config.GetJwtSecretSand()

	_, e := jwt.Parse(jwtToken, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(signingKey), nil
	})

	if e != nil {
		return false
	} else {
		return true
	}
}

func setContext(ctx *gin.Context, s *bo.Session)  {
	session.SetUserId(ctx, s.UserId)
	session.SetUsername(ctx, s.Username)
	session.SetNickname(ctx, s.Nickname)
	session.SetEmail(ctx, s.Email)
	session.SetAvatar(ctx, s.Avatar)
	session.SetRole(ctx, s.Role)
}
