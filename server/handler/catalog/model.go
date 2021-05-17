package catalog

type subCategory struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	FrontName    string `json:"front_name"`
	WapBannerUrl string `json:"wap_banner_url"`
}
type customCategory struct {
	ID              uint          `json:"id"`
	Name            string        `json:"name"`
	FrontName       string        `json:"front_name"`
	WapBannerUrl    string        `json:"wap_banner_url"`
	SubCategoryList []subCategory `json:"sub_category_list"`
}
type ResIndex struct {
	CategoryList    []customCategory `json:"category_list"`
	CurrentCategory customCategory   `json:"current_category"`
}
