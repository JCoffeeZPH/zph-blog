package utils

import (
	"math/rand"
	"strings"
	"time"
	"zph/lib/client"
)

func NowTime() uint32 {
	return uint32(time.Now().Unix())
}

func TimeFormat(timestamp uint32) string {
	timezone := "Asia/Shanghai"
	location, _ := time.LoadLocation(timezone)
	time := time.Unix(int64(timestamp), 0)
	return time.In(location).Format("2006-01-02T15:04:05-07:00")
}

func BoolToInt(flag bool) int {
	if flag {
		return 1
	}
	return 0
}

func IntToBool(i int) bool {
	if i == 0 {
		return false
	}
	return true
}

func GetBlogPrivacy(privacy int) bool {
	if privacy == 1 {
		return false
	}
	return true
}

func GetPrivacy(privacy bool) bool {
	if privacy {
		return false
	}
	return true
}

//时间转时间戳
func StrToUnix(timeStr string) uint32 {
	if len(timeStr) == 0 {
		return 0
	}
	timeStr = strings.ReplaceAll(timeStr, " ", "T")
	timeStr += "+08:00"
	local, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	tt, _ := time.ParseInLocation("2006-01-02T15:04:05-07:00", timeStr, local)
	return uint32(tt.Unix())
}

func GetRandomIds(origin []int64, count int) []int64 {
	tmpOrigin := make([]int64, len(origin))
	copy(tmpOrigin, origin)
	//一定要seed
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(tmpOrigin), func(i int, j int) {
		tmpOrigin[i], tmpOrigin[j] = tmpOrigin[j], tmpOrigin[i]
	})

	result := make([]int64, 0, count)
	for index, value := range tmpOrigin {
		if index == count {
			break
		}
		result = append(result, value)
	}
	return result
}

func GetYesterday() string {
	yesterday := time.Now().AddDate(0, 0, -1)
	return yesterday.Format("20060102")
}

func MarkdownToHtml(markdown string) string {
	c := client.NewMarkdownClient()
	html, err := c.GetMarkdownHtml(markdown)
	if err != nil {
		panic(err)
	}
	return html.Data["html"].(string)
}
