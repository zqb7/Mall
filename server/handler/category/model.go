package category

type customCategory struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	FrontName    string `json:"front_name"`
	WapBannerUrl string `json:"wap_banner_url"`
}

type ResCategory struct {
	CurrentCategory     customCategory   `json:"current_category"`
	ParentCategory      customCategory   `json:"parent_category"`
	BrotherCategoryList []customCategory `json:"brother_category_list"`
}

type goods struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	ListPicUrl  string  `json:"list_pic_url"`
	RetailPrice float64 `json:"retail_price"`
}

type ResGoods struct {
	GoodsList []goods `json:"goods_list"`
}
