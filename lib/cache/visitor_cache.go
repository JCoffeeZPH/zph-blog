package cache

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
	"time"
	"zph/constants"
	api_error "zph/error"
	"zph/models/bo"
)

func getVisitorCodeKey(visitorCode string) string {
	return fmt.Sprintf("blog.visitor.code.%s", visitorCode)
}

func SetVisitorCode(visitorCode string, userId uint64) {
	visitorCacheModel := &bo.VisitorCacheModel{
		BlogOwnerId: userId,
	}
	visitorCacheByte, _ := json.Marshal(visitorCacheModel)
	key := getVisitorCodeKey(visitorCode)
	redisClient.Set(key, visitorCacheByte, constants.VisitorDefaultTimeOut)
}

func GetVisitorCode(visitorCode string) uint64 {
	key := getVisitorCodeKey(visitorCode)
	var visitorModel bo.VisitorCacheModel
	v, err := redisClient.Get(key).Result()
	if err == redis.Nil {
		log.Debugf("GetVisitorCode not found, key is %s", key)
		return 0
	}
	e := json.Unmarshal([]byte(v), &visitorModel)
	if e != nil {
		log.Errorf("GetVisitorCode failed, err is %+v", err)
		return 0
	}
	return visitorModel.BlogOwnerId
}

func generateKey(userId uint64, visitType, date string) string {
	return fmt.Sprintf("blog.%s.%d.%s", date, userId, visitType)
}

func generateUVIpKey(visitType string) string {
	date := time.Now().Format("20060102")
	return fmt.Sprintf("blog.%s.%s", date, visitType)
}

func UpdatePV(userId uint64) {
	date := time.Now().Format("20060102")
	key := generateKey(userId, constants.VisitTypePV, date)
	_, err := redisClient.Get(key).Result()
	if err == redis.Nil {
		// 第二天凌晨两点超时并在此时间段完成同步
		redisClient.Set(key, 1, time.Duration(GetTodayRemainSecond() + 60 * 60 * 2) * time.Second)
	} else if err != nil {
		log.Errorf("UpdatePV err, userId is: %d, err is: %+v", userId, err)
		panic(api_error.CacheError)
	}else {
		redisClient.Incr(key)
	}

}

func GetTodayPV(userId uint64, date string) int {
	key := generateKey(userId, constants.VisitTypePV, date)
	count, err := redisClient.Get(key).Result()
	if err == redis.Nil {
		redisClient.Set(key, 0, time.Duration(GetTodayRemainSecond() + 60 * 60 * 2) * time.Second)
		return 0
	}else if err != nil {
		log.Errorf("GetTodayPV failed, userId is: %d, err is: %+v", userId, err)
		panic(api_error.CacheError)
	}else {
		total, _ := strconv.Atoi(count)
		return total
	}
}

func UpdateUV(userId uint64, ip string) {
	// 查set是否存在此ip今日已经访问过此用户
	key := generateUVIpKey(constants.VisitTypeUV)
	value := strconv.Itoa(int(userId)) + "-" + ip
	exist, err := redisClient.SIsMember(key, value).Result()
	if err != nil {
		log.Errorf("UpdateUV failed, userId is: %d, ip is: %s, err is: %+v", userId, ip, err)
		panic(api_error.CacheError)
	}
	if exist {
		return
	}
	// 若没有则将ip加入set，并更新uv
	redisClient.SAdd(key, value)
	date := time.Now().Format("20060102")
	uvKey := generateKey(userId, constants.VisitTypeUV, date)
	_, e := redisClient.Get(uvKey).Result()
	if e == redis.Nil {
		// 第二天凌晨两点超时并在此时间段完成同步
		redisClient.Set(uvKey, 1, time.Duration(GetTodayRemainSecond() + 60 * 60 * 2) * time.Second)
	} else if e != nil {
		log.Errorf("UpdateUV failed, userId is: %d, ip is: %s, err is: %+v", userId, ip, err)
		panic(api_error.CacheError)
	}else {
		redisClient.Incr(uvKey)
	}
}

func GetTodayUV(userId uint64, date string) int {
	key := generateKey(userId, constants.VisitTypeUV, date)
	count, err := redisClient.Get(key).Result()
	if err == redis.Nil {
		redisClient.Set(key, 0, time.Duration(GetTodayRemainSecond() + 60 * 60 * 2) * time.Second)
		return 0
	}else if err != nil {
		log.Errorf("GetTodayUV failed, userId is: %d, err is: %+v", userId, err)
		panic(api_error.CacheError)
	}else {
		total, _ := strconv.Atoi(count)
		return total
	}
}


func GetTodayRemainSecond() int64 {
	todayLast := time.Now().Format("2006-01-02") + " 23:59:59"
	todayLastTime, _ := time.ParseInLocation("2006-01-02 15:04:05", todayLast, time.Local)
	remainSecond := todayLastTime.Unix() - time.Now().Local().Unix()
	return remainSecond
}

func generatorAllPVKey(date string) string {
	return fmt.Sprintf("blog.%s.*.PV", date)
}

func generatorAllUVKey(date string) string {
	return fmt.Sprintf("blog.%s.*.UV", date)
}

func GetAllPV(date string) (map[int]int, []string){
	pattern := generatorAllPVKey(date)
	// blog.20211227.2.PV
	keys, err := redisClient.Keys(pattern).Result()
	if err != nil {
		log.Errorf("GetAllPV failed, date is: %s, err is: %+v", date, err)
		panic(api_error.CacheError)
	}
	pvMap := make(map[int]int)
	for _, key := range keys {
		key = key[:strings.LastIndex(key, ".")]
		userId := key[strings.LastIndex(key, ".") + 1:]
		id, _ := strconv.Atoi(userId)
		pv := GetTodayPV(uint64(id), date)
		pvMap[id] = pv
	}
	return pvMap, keys
}

func GetAllUV(date string) (map[int]int, []string) {
	pattern := generatorAllUVKey(date)
	// blog.20211227.3.UV
	keys, err := redisClient.Keys(pattern).Result()
	if err != nil {
		log.Errorf("GetAllPV failed, date is: %s, err is: %+v", date, err)
		panic(api_error.CacheError)
	}
	uvMap := make(map[int]int)
	for _, key := range keys {
		key = key[:strings.LastIndex(key, ".")]
		userId := key[strings.LastIndex(key, ".") + 1:]
		id, _ := strconv.Atoi(userId)
		pv := GetTodayUV(uint64(id), date)
		uvMap[id] = pv
	}
	return uvMap, keys
}

func Del(keys []string) {
	redisClient.Del(keys ...)
}