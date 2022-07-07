package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"zph/constants"
)

func generateBlogKey(blogId uint64) string {
	return fmt.Sprintf("blog.%d.title", blogId)
}

func GetTitleByBlogId(blogId uint64) (string, bool) {
	key := generateBlogKey(blogId)
	title, err := redisClient.Get(key).Result()
	if err == redis.Nil {
		return "", false
	}
	return title, true
}

func SetTitleByBlogId(blogId uint64, title string) {
	key := generateBlogKey(blogId)
	_, err := redisClient.Set(key, title, constants.BLogTitleDefaultTimeOut).Result()
	if err != nil {
		log.Errorf("setTitleByBlogId failed, err is: %+v, blogId is: %d, title is: %s", err, blogId, title)
		panic(err)
	}
}
