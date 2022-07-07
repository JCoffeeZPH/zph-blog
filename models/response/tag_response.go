package response

import "zph/models/db"

type GetTagResponse struct {
	Total uint64 `json:"total"`
	Tags []TagResponse `json:"tags"`
}

type TagResponse struct {
	TagId int `json:"id"`
	TagName string `json:"tag_name"`
	Color 	string `json:"color"`
}

func NewTagResponse(tag db.Tag) TagResponse {
	return TagResponse{
		TagId: tag.TagId,
		TagName: tag.TagName,
		Color: tag.Color,
	}
}