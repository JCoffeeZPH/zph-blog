package bo


type SiteSettingModel struct {
	Id int `json:"id"`
	NameEn string `json:"name_en"`
	NameZh string `json:"name_zh"`
	Type int8 `json:"type"`
	Value string `json:"value"`
	UserId int `json:"user_id"`
}
