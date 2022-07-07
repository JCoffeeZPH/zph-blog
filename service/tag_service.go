package service

import (
	"zph/dao"
	"zph/models/db"
	"zph/models/request"
)

type TagService struct {
	tagDao dao.TagDao
}

func NewTagService() *TagService {
	return &TagService{
		tagDao: dao.NewTagDao(),
	}
}

func (service *TagService) GetTags(offset, limit, userId int)([]db.Tag, uint64) {
	total := service.tagDao.GetTagCount(userId)
	tags := service.tagDao.GetTags(offset, limit, userId)
	return tags, total
}

func (service *TagService) UpdateTag(tag *request.UpdateTagRequest)  {
	tagId := tag.TagId
	params := map[string]string{}
	if tag.TagName != "" {
		params["tag_name"] = tag.TagName
	}
	if tag.Color != "" {
		params["color"] = tag.Color
	}
	service.tagDao.UpdateTagById(tagId, params)
}

func (service *TagService)CreateNewTag(tag *request.NewTagRequest, userId uint64)  {
	service.tagDao.CreateTag(tag, userId)
}

func (service *TagService) DeleteTagById(tagId uint64) {
	service.tagDao.DeleteTag(tagId)
}

func (service *TagService)GetTagsCount(userId int)uint64  {
	return service.tagDao.GetTagCount(userId)
}

func (service *TagService) GetAllTags(userId uint64)[]db.Tag {
	return service.tagDao.GetAllTags(userId)
}

