package response

import "zph/models/bo"

type SiteSettingResponse struct {
	BasicSettings []bo.SiteSettingModel `json:"basic_settings"`
	FooterSettings []bo.SiteSettingModel `json:"footer_settings"`
	DataCardSettings []bo.SiteSettingModel `json:"data_card_settings"`
}

