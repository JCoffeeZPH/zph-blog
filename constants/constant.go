package constants

import (
	"time"
)

type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December

	// BlogOwnerId 由于前端原因暂时写死
	BlogOwnerId uint64 = 2

	SessionDefaultTimeOut   = 7 * 24 * time.Hour
	VisitorDefaultTimeOut   = 24 * time.Hour
	BLogTitleDefaultTimeOut = 7 * 24 * time.Hour

	DefaultPage    = 1
	DefaultPerPage = 10

	DefaultReadSpeedPerMinute = 300

	OS           string = "操作系统"
	Browser      string = "浏览器"
	LoginSuccess        = "登录成功"
	LoginDefeat         = "登录失败"

	BasicSetting       int = 1
	FooterSetting      int = 2
	DataCardSetting    int = 3
	FriendChainSetting int = 4

	VisitPerPage       = 5
	VisitBlogOwnerId   = "BLOG_OWNER_ID"
	VisitorBlogOwnerId = "user_id"

	VisitTypePV string = "PV"
	VisitTypeUV string = "UV"

	CommentPublished       = 1
	CommentUnPublished     = 0
	CommentAllStatus       = 2
	CommentDefaultParentId = -1
	InternalNetIp          = "内网"
)

type FromPage int8

const (
	Blog        FromPage = 1
	AboutMe     FromPage = 2
	FriendChain FromPage = 3

	BlogPage        string = "Article"
	AboutMePage     string = "About me"
	FriendChainPage string = "Friend Chain"
	UnknownPage     string = "unknown"
)

func (f FromPage) ToString() string {
	if f == Blog {
		return BlogPage
	} else if f == AboutMe {
		return AboutMePage
	} else if f == FriendChain {
		return FriendChainPage
	} else {
		return UnknownPage
	}
}

var FromPageMap = map[int8]string{
	1: BlogPage,
	2: AboutMePage,
	3: FriendChainPage,
}
