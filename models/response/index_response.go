package response

type IndexResponse struct {
	Badges []BadgeSetting `json:"badges"`
	CategoryList []CategoryResponse `json:"category_list"`
	DataCardSetting DataCardSetting `json:"data_card_setting"`
	BasicSetting BasicSetting `json:"basic_setting"`
	TagList []TagResponse `json:"tag_list"`
	LatestBlogList []LatestBlog `json:"latest_blog"`
	RandomBlogList []RandomBlog `json:"random_blog_list"`
}

type BadgeSetting struct {
	Title string `json:"title"`
	Url string `json:"url"`
	Subject string `json:"subject"`
	Value string `json:"value"`
	Color string `json:"color"`
}

type DataCardSetting struct {
	Avatar string `json:"avatar"`
	BiliBili string `json:"bilibili"`
	Email string `json:"email"`
	Favorites []Favorite `json:"favorites"`
	Github string `json:"github"`
	Name string `json:"name"`
	Netease string `json:"netease"`
	QQ string `json:"qq"`
	RollText []string `json:"rollText"`
}

type Favorite struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

type BasicSetting struct {
	BeiAn string `json:"beian"`
	BlogName string `json:"blog_name"`
	CommentAdminFlag string `json:"comment_admin_flag"`
	Copyright string `json:"copyright"`
	FooterImgTitle string `json:"footer_img_title"`
	FooterImgUrl string `json:"footer_img_url"`
	Reward string `json:"reward"`
	WebTitleSuffix string `json:"web_title_suffix"`
}

type LatestBlog struct {
	Id int `json:"id"`
	Password string `json:"password"`
	Privacy bool `json:"privacy"`
	Title string `json:"title"`
}

type RandomBlog struct {
	Id int `json:"id"`
	CreateTime string `json:"create_time"`
	FirstPicture string `json:"first_picture"`
	Password string `json:"password"`
	Privacy bool `json:"privacy"`
	Title string `json:"title"`
}