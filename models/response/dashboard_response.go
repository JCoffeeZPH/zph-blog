package response

type DashboardResponse struct {
	BlogCount int `json:"blog_count"`
	Category Category `json:"category"`
	CommentCount int `json:"comment_count"`
	PV int `json:"pv"`
	UV int `json:"uv"`
	Tag Tag `json:"tag"`
	VisitRecord VisitRecord `json:"visit_record"`
	CityVisitor []CityVisitor `json:"city_visitor"`
}

type Category struct {
	Legend []string `json:"legend"`
	Series []DashboardCategory `json:"series"`
}

type DashboardCategory struct {
	CategoryId int `json:"id"`
	CategoryName string `json:"name"`
	Value int `json:"value"` // 分类下的文章数
}

type Tag struct {
	Legend []string `json:"legend"`
	Series []DashboardTag `json:"series"`
}

type DashboardTag struct {
	TagId int `json:"id"`
	TagName string `json:"name"`
	Value int `json:"value"`
}

type CityVisitor struct {
	City string `json:"city"`
	PV int `json:"pv"`
}