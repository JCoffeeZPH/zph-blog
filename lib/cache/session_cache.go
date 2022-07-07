package cache

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"zph/constants"
	"zph/global"
	"zph/logger"
	"zph/models/bo"
)

var (
	redisClient = global.RedisClient
	log = logger.Log
)

func sessionRedisKey(sessionId string) string {
	return fmt.Sprintf("blob.user_session.%s", sessionId)
}

func SessionBySessionIdFromRedis(sessionId string) *bo.Session {
	session, err := redisClient.Get(sessionRedisKey(sessionId)).Result()
	if err == redis.Nil {
		log.Debugf("SessionBySessionIdFromRedis not found session, sessionId is %s", sessionId)
		return nil
	}
	s := &bo.Session{}
	err = json.Unmarshal([]byte(session), s)
	if err != nil {
		log.Errorf("SessionBySessionIdFromRedis failed, err is %+v", err)
		return nil
	}
	return s
}

func SaveSessionIntoRedis(session *bo.Session) {
	key := sessionRedisKey(session.SessionId)
	sessionByte, _ := json.Marshal(session)
	redisClient.Set(key, sessionByte, constants.SessionDefaultTimeOut)
}

