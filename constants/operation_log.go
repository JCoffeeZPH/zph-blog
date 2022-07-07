package constants

type OperationType string

const (
	OperationKey string = "OPERATION_NAME"

	GetTagsAndCategories OperationType = "获取全部标签和分类"
	CreateNewBlog OperationType = "创建blog"
	GetBlogById OperationType = "根据Id获取blog"
	GetBlogs OperationType = "按分页获取blogs"
	UpdateBlog OperationType = "更新blog"
	DeleteBlog OperationType = "更新blog"
	UpdateBlogTopById OperationType = "更新blog置顶状态"
	UpdateBlogRecommendById OperationType = "更新blog可评论状态"
	UpdateVisibility OperationType = "更新blog状态"

	GetCategories OperationType = "分页获取分类"
	UpdateCategory OperationType = "更新分类信息"
	CreateCategory OperationType = "创建分类"
	DeleteCategory OperationType = "删除分类"

	LoginOperation OperationType = "登录"

	CreateMoment OperationType = "新建动态"
	GetMoments OperationType = "分页获取动态"
	GetMomentById OperationType = "根据Id获取动态"
	UpdateMomentPublished OperationType = "更新动态发布状态"
	UpdateMoment OperationType = "更新动态信息"
	DeleteMoment OperationType = "删除动态"

	GetTags OperationType = "分页获取标签"
	CreateTag OperationType = "新建标签"
	UpdateTag OperationType = "更新标签信息"
	DeleteTag OperationType = "删除标签"

	GetSiteSettings OperationType = "获取站点配置"
	UpdateSiteSettings OperationType = "更新站点配置"

	GetFriendChainPageInfo OperationType = "获取友链页面信息"
	CreateNewFriendChain OperationType = "新加友链"
	UpdateFriendChain OperationType = "更新友链信息"
	DeleteFriendChain OperationType = "删除友链"
	GetFriendChains OperationType = "获取友链"
	UpdateCommentStatus OperationType = "修改友链页面评论开关"
	UpdateIsPublished OperationType = "修改友链是否公开开关"
	UpdatePageInfo OperationType = "修改友链页面信息"

	GetAboutMeInfo OperationType = "获取AboutMe信息"
	UpdateAboutMeInfo OperationType = "修改AboutMe信息"

	VisitDashboardPage OperationType = "查看首页仪表盘"
)

func (o OperationType) Value() string {
	return string(o)
}