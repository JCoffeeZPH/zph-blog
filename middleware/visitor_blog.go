package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"zph/constants"
	api_error "zph/error"
	"zph/lib/cache"
	"zph/lib/client"
	"zph/service"
)

func VisitorVisitBlogPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		userId := cache.GetBlogOwnerUserDetail(username)
		if userId == 0 {
			userId = getUseridByName(username)
			cache.SetBlogOwnerUserDetail(username, userId)
		}
		//userId := constants.BlogOwnerId
		ctx.Set(constants.VisitorBlogOwnerId, userId)

		go cache.UpdatePV(userId)
		if ip := getIpDetails(); len(ip) > 0 {
			go cache.UpdateUV(userId, ip)
		}

		// 新加city pv
		go updateCityPV(userId)

		ctx.Next()
	}
}

func getIpDetails() string {
	ipClient := client.NewIPClient()
	ipDetails, err := ipClient.GetIPAndAttribution()
	if err != nil {
		panic(err)
	}
	if len(ipDetails.IP) > 0 {
		return strings.ReplaceAll(ipDetails.IP, ".", "-")
	}
	return ""
}

func getUseridByName(username string) uint64 {
	s := service.NewUserService()
	user := s.GetUserByUsername(username)
	if user == nil {
		panic(api_error.NotFound)
	}
	return user.Id
}

func updateCityPV(userId uint64) {
	ipClient := client.NewIPClient()
	ipDetails, err := ipClient.GetIPAndAttribution()
	if err != nil {
		panic(err)
	}

	cityVisitorService := service.NewCityVisitorService()

	// ipSource 中国|广东省|深圳市|联通 / 广东省|深圳市
	city := ipDetails.IpSource[strings.LastIndex(ipDetails.IpSource, "|")+1:]
	pv := cityVisitorService.GetCityPVByCity(userId, city)
	if pv < 0 {
		cityVisitorService.CreateNewCity(city, userId)
	} else {
		cityVisitorService.UpdateCityVisitor(int(userId), pv+1, city)
	}
}

func ReadBlogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")
		if len(userId) == 0 {
			panic(api_error.ParamError)
		}
		id, _ := strconv.Atoi(userId)
		go cache.UpdatePV(uint64(id))
		if ip := getIpDetails(); len(userId) > 0 {
			go cache.UpdateUV(uint64(id), ip)
		}

		// 新加city pv
		go updateCityPV(uint64(id))

		ctx.Next()
	}
}
