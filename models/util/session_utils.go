package util

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"zph/config"
	"zph/lib/cache"
	"zph/logger"
	"zph/models/bo"
)
var log = logger.Log

func ParseJwtToken(token string) string  {
	t, err := jwt.Parse(token, nil)
	if t == nil {
		log.Errorf("parseJwtToken, invalid token: %s, err: %+v", token, err)
		return ""
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		return claims["id"].(string)
	} else {
		return ""
	}
}

func GenerateJwtToken(sessionId, secret string) string {
	now := uint32(time.Now().Unix())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          sessionId,
		"create_time": now,
	})
	key := GenerateJwtTokenKey(secret)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		log.Errorf("GenerateJwtToken failed, err is: %+v", err)
	}
	return tokenStr
}

func GenerateJwtTokenKey(secret string) []byte {
	md5Value := md5.Sum([]byte(secret))
	signingKey := fmt.Sprintf("%x", md5Value) + config.GetJwtSecretSand()

	return []byte(signingKey)
}

func ParseSession(jwtToken string) *bo.Session {
	sessionId := ParseJwtToken(jwtToken)
	if len(sessionId) == 0 {
		return nil
	}
	session := cache.SessionBySessionIdFromRedis(sessionId)
	return session
}

