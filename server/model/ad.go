package model

type AD struct {
	Model
	Name     string `json:"name"`
	Link     string `json:"link"`
	ImageUrl string `json:"image_url"`
	Content  string `json:"content"`
	Enabled  bool   `json:"enabled" gorm:"defalut:0;NOT NULL"`
}

func GetADs() (ads []AD, err error) {
	ads = make([]AD, 0)
	err = db.Model(AD{}).Not("enabled", 0).Find(&ads).Error
	return
}
