package handler

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"zph/lib/cache"
	"zph/models/db"
	"zph/service"
	"zph/utils"
)

const (
	// SyncCacheToDB 每天凌晨1点同步昨天PV和UV
	SyncCacheToDB = "00 00 01 * * ?"
)


func SyncBlogDataHandler() {
	c := cron.New(cron.WithSeconds())
	yesterday := utils.GetYesterday()
	c.AddFunc(SyncCacheToDB, func() {
		s := service.NewVisitRecordService()
		pvMap, pvKeys := cache.GetAllPV(yesterday)
		uvMap, uvKeys := cache.GetAllUV(yesterday)
		userIdMap := make(map[int]bool)
		for userId := range pvMap {
			userIdMap[userId] = true
		}
		for userId := range uvMap {
			userIdMap[userId] = true
		}

		if len(userIdMap) == 0 {
			return
		}

		for userId := range userIdMap {
			var record db.VisitRecord
			if pv, ok := pvMap[userId]; ok {
				record.PV = pv
			}

			if uv, ok := uvMap[userId]; ok {
				record.UV = uv
			}

			record.Date = yesterday
			record.UserId = userId

			s.CreateVisitRecord(record)
		}

		cache.Del(pvKeys)
		cache.Del(uvKeys)
		cache.Del([]string{fmt.Sprintf("blog.%s.UV", yesterday)})
	})
	c.Start()

}
