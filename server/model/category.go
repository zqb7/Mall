package model

type Category struct {
	Model
	Name         string `gorm:"type:varchar(255)"`
	ParentID     uint
	FrontName    string
	FrontDesc    string
	SortOrder    uint `gorm:"type:tinyint(1);default:0;NOT NULL"`
	WapBannerUrl string
}

func GetCategories(selectFields []string, where map[string]interface{}) (categories []Category, err error) {
	categories = make([]Category, 0)
	err = db.Model(Category{}).Select(selectFields).Where(where).Find(&categories).Error
	return categories, err
}

func (c *Category) Get(selectFields []string, where map[string]interface{}) error {
	err := db.Select(selectFields).Where(where).First(c).Error
	return err
}
