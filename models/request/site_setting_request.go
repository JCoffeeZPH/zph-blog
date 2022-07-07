package request

import "zph/models/bo"

type SiteSettingRequest struct {
	Settings []bo.SiteSettingModel `json:"settings"`
	DeleteIds []int `json:"delete_ids"`
}
