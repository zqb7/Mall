package topic

type Response struct {
	Count       int     `json:"count"`
	PageSize    int     `json:"page_size"`
	CurrentPage int     `json:"current_page"`
	TopicList   []topic `json:"topic_list"`
	HasNextPage bool    `json:"has_next_page"`
}

type topic struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	PriceInfo   float64 `json:"price_info"`
	ScenePicUrl string  `json:"scene_pic_url"`
	Subtitle    string  `json:"subtitle"`
}

type query struct {
	Page int `form:"page" binding:"number"`
	Size int `form:"size" binding:"number"`
}
