package model

//品牌商
type Brand struct {
	Model
	Name       string  `json:"name"`
	PicUrl     string  `json:"pic_url"`
	SimpleDesc string  `json:"simple_desc"`
	Price      float64 `json:"price" gorm:"type:decimal(10,2);NOT NULL"`
	SortOrder  uint    `json:"sort_order" gorm:"type:tinyint(1);default:0;NOT NULL"`
}

func GetLastBrand() ([]Brand, error) {
	brands := make([]Brand, 0)
	err := db.Model(Brand{}).Limit(4).Find(&brands).Error
	return brands, err
}
