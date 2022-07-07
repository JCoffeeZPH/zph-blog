package service

import (
	"zph/constants"
	"zph/dao"
	"zph/lib/common"
	"zph/models/db"
	"zph/models/request"
	"zph/models/response"
	"zph/utils"
)

type BlogService struct {
	blogDao       dao.BlogDao
	categoryDao   dao.CategoryDao
	tagDao        dao.TagDao
	blogTagDao    dao.BlogTagDao
	blogUserIdDao dao.BlogUserDao
}

func NewBlogService() *BlogService {
	return &BlogService{
		blogDao:       dao.NewBlogDao(),
		categoryDao:   dao.NewCategoryDao(),
		tagDao:        dao.NewTagDao(),
		blogTagDao:    dao.NewBlogTagDao(),
		blogUserIdDao: dao.NewBlogUserDao(),
	}
}

func (service *BlogService) UpdateBlog(req *request.UpdateBlogRequest, userId uint64) {
	categoryId := service.checkCategory(req.Category, userId)
	tagIds := service.checkTags(req.Tags, userId)
	params := service.getParams(req, categoryId)

	service.blogDao.UpdateBlog(req.BlogId, params)
	service.blogTagDao.DeleteBlogIdTags(req.BlogId)
	service.blogTagDao.InsertIntoTagBlog(req.BlogId, tagIds)
}

func (service *BlogService) getParams(req *request.UpdateBlogRequest, categoryId int) map[string]interface{} {
	params := make(map[string]interface{})
	if len(req.Title) > 0 {
		params["title"] = req.Title
	}
	if len(req.FirstPicture) > 0 {
		params["first_picture"] = req.FirstPicture
	}

	if len(req.Content) > 0 {
		params["content"] = req.Content
	}

	if len(req.Description) > 0 {
		params["description"] = req.Description
	}
	params["is_published"] = utils.BoolToInt(req.IsPublished)
	params["is_recommend"] = utils.BoolToInt(req.IsRecommend)
	params["is_appreciation"] = utils.BoolToInt(req.IsAppreciation)
	params["is_comment_enabled"] = utils.BoolToInt(req.IsCommentEnabled)
	params["password"] = req.Password
	params["words"] = len(req.Content)
	params["read_time"] = int(len(req.Content)/constants.DefaultReadSpeedPerMinute) + 1
	params["update_time"] = utils.NowTime()
	params["category_id"] = categoryId

	return params
}

func (service *BlogService) CreateNewBlog(userId uint64, req *request.BlogRequest) {
	blog, tagIds := service.BlogReqToDbModel(userId, req)
	blogId := service.blogDao.CreateBlog(blog)
	service.blogTagDao.InsertIntoTagBlog(blogId, tagIds)
	//service.blogUserIdDao.CreateNewRelation(db.BlogUserId{
	//	UserId: userId,
	//	BlogId: blogId,
	//})
}

func (service *BlogService) BlogReqToDbModel(userId uint64, req *request.BlogRequest) (*db.Blog, []int) {
	nowTime := utils.NowTime()
	readTime := int(len(req.Content)/constants.DefaultReadSpeedPerMinute) + 1
	// 检查tag是否需要新建
	tagIds := service.checkTags(req.Tags, userId)

	blog := &db.Blog{
		Title:            req.Title,
		FirstPicture:     req.FirstPicture,
		Content:          req.Content,
		Description:      req.Description,
		IsPublished:      utils.BoolToInt(req.IsPublished),
		IsRecommend:      utils.BoolToInt(req.IsRecommend),
		IsAppreciation:   utils.BoolToInt(req.IsAppreciation),
		IsCommentEnabled: utils.BoolToInt(req.IsCommentEnabled),
		UpdateTime:       nowTime,
		CreateTime:       nowTime,
		WordCount:        len(req.Content),
		ReadTime:         readTime,
		CategoryId:       service.checkCategory(req.Category, userId),
		IsTop:            utils.BoolToInt(req.IsTop),
		Password:         req.Password,
	}
	if userId != 0 {
		blog.UserId = userId
	} else {
		panic(common.ServiceError{Err: "userId is empty", API: "BlogReqToDbModel"})
	}

	return blog, tagIds
}

func (service *BlogService) checkCategory(category interface{}, userId uint64) int {
	categoryId := 0
	// 检查category是否需要新建
	switch category.(type) {
	case string:
		categoryId = service.categoryDao.CreateCategory(category.(string), userId)
		break
	default:
		categoryId = int(category.(float64))
	}
	return categoryId
}

func (service *BlogService) checkTags(tags []interface{}, userId uint64) []int {
	tagIds := make([]int, 0)
	for _, tag := range tags {
		switch tag.(type) {
		case string:
			req := &request.NewTagRequest{
				TagName: tag.(string),
			}
			tagId := service.tagDao.CreateTag(req, userId)
			tagIds = append(tagIds, tagId)
		default:
			tagIds = append(tagIds, int(tag.(float64)))
		}
	}
	return tagIds
}

func (service *BlogService) GetBlogById(blogId uint64) response.BlogResponse {
	blog := service.blogDao.GetBlogById(blogId)
	resp := service.NewGetBlogResponse(blog)
	return resp
}

func (service *BlogService) GetBlogs(userId, offset, limit, categoryId int, title string) ([]response.BlogResponse, int) {
	blogs := service.blogDao.GetBlogs(userId, offset, limit, categoryId, title)
	total := service.blogDao.GetBlogCount(userId)
	resp := make([]response.BlogResponse, 0)
	for _, blog := range blogs {
		resp = append(resp, service.NewGetBlogResponse(blog))
	}
	return resp, total
}

func (service *BlogService) GetBlogCount(userId uint64) int {
	return service.blogDao.GetBlogCount(int(userId))
}

func (service *BlogService) DeleteBlogById(blogId, userId uint64) {
	service.blogDao.DeleteBlogById(blogId, userId)
	service.blogTagDao.DeleteBlogIdTags(int(blogId))
	//service.blogUserIdDao.DeleteRelationByBlogIdUserId(userId, blogId)
}

func (service *BlogService) NewGetBlogResponse(blog db.Blog) response.BlogResponse {
	blogId := blog.BlogId
	categoryId := blog.CategoryId
	category := service.categoryDao.GetCategoryById(categoryId)
	tagsIds := service.blogTagDao.GetTagIdsByBlogId(int(blogId))
	tags := service.tagDao.GetTagsByIds(tagsIds)
	tagResp := make([]response.TagResponse, 0)
	for _, tag := range tags {
		t := response.TagResponse{
			TagId:   tag.TagId,
			TagName: tag.TagName,
			Color:   tag.Color,
		}
		tagResp = append(tagResp, t)
	}
	return response.BlogResponse{
		BlogId:           blog.BlogId,
		Title:            blog.Title,
		FirstPicture:     blog.FirstPicture,
		Content:          blog.Content,
		Description:      blog.Description,
		IsPublished:      utils.IntToBool(blog.IsPublished),
		IsRecommend:      utils.IntToBool(blog.IsRecommend),
		IsAppreciation:   utils.IntToBool(blog.IsAppreciation),
		IsCommentEnabled: utils.IntToBool(blog.IsCommentEnabled),
		CreateTime:       utils.TimeFormat(blog.CreateTime),
		UpdateTime:       utils.TimeFormat(blog.UpdateTime),
		ViewCount:        blog.ViewCount,
		WordCount:        blog.WordCount,
		ReadTime:         blog.ReadTime,
		IsTop:            utils.IntToBool(blog.IsTop),
		Password:         blog.Password,
		UserId:           blog.UserId,
		Category: response.CategoryResponse{
			CategoryId:   category.CategoryId,
			CategoryName: category.CategoryName,
		},
		Tags: tagResp,
	}
}

func (service *BlogService) UpdateBlogTop(blogId, userId int, isTop bool) {
	service.blogDao.UpdateBlogTopStatus(blogId, userId, isTop)
}

func (service *BlogService) UpdateBlogRecommend(blogId, userId int, isRecommend bool) {
	service.blogDao.UpdateBlogRecommendStatus(blogId, userId, isRecommend)
}

func (service *BlogService) UpdateBlogVisibility(blogId, userId int, visibilityRequest *request.VisibilityRequest) {
	params := map[string]interface{}{
		"is_top":             utils.BoolToInt(visibilityRequest.IsTop),
		"is_appreciation":    utils.BoolToInt(visibilityRequest.IsAppreciation),
		"is_comment_enabled": utils.BoolToInt(visibilityRequest.IsCommentEnabled),
		"is_published":       utils.BoolToInt(visibilityRequest.IsPublished),
		"is_recommend":       utils.BoolToInt(visibilityRequest.IsRecommend),
		"password":           visibilityRequest.Password,
		"update_time":        utils.NowTime(),
	}
	service.blogDao.UpdateBlogVisibilityStatus(blogId, userId, params)
}

func (service *BlogService) GetRandomBlogList(userId uint64) []response.RandomBlog {
	ids := service.blogDao.GetAllBlogIds(userId)
	if len(ids) == 0 {
		return []response.RandomBlog{}
	}
	var randomIds []int64
	if len(ids) <= 5 {
		randomIds = ids
	} else {
		randomIds = utils.GetRandomIds(ids, 5)
	}
	blogs := service.blogDao.GetBlogsByIds(randomIds, userId)
	resp := make([]response.RandomBlog, 0)
	for _, blog := range blogs {
		resp = append(resp, response.RandomBlog{
			Id:           blog.BlogId,
			CreateTime:   utils.TimeFormat(blog.CreateTime),
			FirstPicture: blog.FirstPicture,
			Password:     blog.Password,
			Privacy:      utils.GetBlogPrivacy(blog.IsPublished),
			Title:        blog.Title,
		})
	}
	return resp
}

func (service *BlogService) GetBlogCountByCategoryId(userId, categoryId int) int {
	return service.blogDao.GetBlogCountByUserIdCategoryId(userId, categoryId)
}

func (service *BlogService) GetBlogCountByTagId(tagId int) int {
	return service.blogTagDao.GetBlogCountByTagId(tagId)
}

func (service *BlogService) IncrViewCount(userId, blogId, target int) {
	service.blogDao.IncrBlogViewCount(userId, blogId, target)
}

// todo
func (service *BlogService) BatchGetBlogsByIds(blogIds []uint64) []db.Blog {
	//var res = make(chan db.Blog)
	//defer close(res)
	//maxArrLen := 10
	//for i := 0; i < len(blogIds); i += maxArrLen {
	//	tempNum := i
	//	tempIds := blogIds[i : i+maxArrLen]
	//	f := func() {
	//		fmt.Printf("execute serial num is: %d, tempIds len is: %d\n", tempNum, len(tempIds))
	//		blogs := service.blogDao.BatchGetBlogs(tempIds)
	//		for _, blog := range blogs {
	//			res <- blog
	//		}
	//		fmt.Printf("tempNum: %d ======================================================\n", tempNum)
	//	}
	//	utils.Execute(f)
	//}
	//blogs := make([]db.Blog, 0)
	//for blog := range res {
	//	blogs = append(blogs, blog)
	//}
	//fmt.Println(blogs)
	//return blogs
	return []db.Blog{}
}
