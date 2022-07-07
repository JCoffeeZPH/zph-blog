package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
	"zph/constants"
)

func getVisitRecordKey(visitType string) string {
	return fmt.Sprintf("blog.visit.record.%s", visitType)
}

// GetCurrentPVAndUV 使用两个hash结构存储所有的pv和uv， blog.visit.record.PV userId1 10 userId2 200
func GetCurrentPVAndUV(userId uint64) (pv, uv int) {
	pvKey := getVisitRecordKey(constants.VisitTypePV)
	uvKey := getVisitRecordKey(constants.VisitTypeUV)

	userPV, e1 := redisClient.HGet(pvKey, strconv.FormatUint(userId, 10)).Result()
	if e1 == redis.Nil {
		pv = 0
	}else {
		pv, _ = strconv.Atoi(userPV)
	}
	userUV, e2 := redisClient.HGet(uvKey, strconv.FormatUint(userId, 10)).Result()
	if e2 == redis.Nil {
		uv = 0
	}else {
		uv, _ = strconv.Atoi(userUV)
	}
	return
}

func UpdatePVOrUV(userId uint64, visitType string)  {
	key := getVisitRecordKey(visitType)
	count, _ := redisClient.HGet(key, strconv.FormatUint(userId, 10)).Result()
	newCount, _ := strconv.Atoi(count)
	SetPVOrUV(userId, visitType, newCount + 1)
}

func SetPVOrUV(userId uint64, visitType string, newValue int) {
	key := getVisitorCodeKey(visitType)
	redisClient.HSet(key, strconv.FormatUint(userId, 10), newValue)
}

func getCheckIpSetKey() string {
	return fmt.Sprintf("blog.visit.record.ipset")
}

func SetBlogOwnerUserDetail(username string, userId uint64) uint64 {
	key := generatorBlogOwnerUserDetailKey(username)
	_, err := redisClient.Set(key, userId, 12 * time.Hour).Result()
	if err != nil {
		log.Errorf("SetBlogOwnerUserDetail failed, username is: %s, err is: %+v", username, err)
		panic(err)
	}
	return userId
}

func GetBlogOwnerUserDetail(username string) uint64 {
	key := generatorBlogOwnerUserDetailKey(username)
	userId, err := redisClient.Get(key).Result()
	if err == redis.Nil {
		return 0
	}else if err != nil {
		log.Errorf("GetBlogOwnerUserDetail failed, username is: %s, err is: %+v", username, err)
		panic(err)
	}
	id, _ := strconv.Atoi(userId)
	return uint64(id)
}

func generatorBlogOwnerUserDetailKey(username string) string {
	return fmt.Sprintf("blog.owner.detail.%s", username)
}


