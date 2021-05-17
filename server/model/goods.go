package model

type Goods struct {
	Model
	CategoryID  uint
	Name        string
	ListPicUrl  string
	RetailPrice float64
	GoodsBrief  string `json:"goods_brief" gorm:"not null;default:''"`
	GoodsDesc   string `gorm:"type:text"` //商品详情页 商品介绍
	SortOrder   uint   `gorm:"type:tinyint(1);default:0;NOT NULL"`
	IsHot       uint   `json:"is_hot" gorm:"type:tinyint(1) unsigned; not null;default:'0'"`
	IsNew       uint   `json:"is_new" gorm:"type:tinyint(1) unsigned; not null;default:'0'"`
}

func (g *Goods) Get(selectFields []string) error {
	err := db.Select(selectFields).First(g).Error
	return err
}

func GetGoods(selectFields []string, where map[string]interface{}, limit, offset int) ([]Goods, error) {
	goods := make([]Goods, 0)
	err := db.Model(Goods{}).Select(selectFields).Where(where).Limit(limit).Offset(offset).Find(&goods).Error
	return goods, err
}

func GetGoodsByCategoryIDs(selectFields []string, ins []uint, limit, offset int) ([]Goods, error) {
	goods := make([]Goods, 0)
	err := db.Model(Goods{}).Select(selectFields).Where("category_id in (?)", ins).
		Limit(limit).Offset(offset).Find(&goods).Error
	return goods, err
}

func (g *Goods) Count() (count int) {
	_ = db.Model(g).Count(&count).Error
	return count
}
