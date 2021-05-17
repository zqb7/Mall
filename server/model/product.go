package model

// 产品表 由Goods表和SpecificationValue组合出来的同一商品的不同产品
type Product struct {
	Model
	GoodsID               uint    `gorm:"NOT NULL;default:0"`
	GoodsSn               string  //商品编号
	GoodsSpecificationIDs string  `gorm:"type:varchar(100)"`
	GoodsNumber           uint    `gorm:"NOT NULL;default:0"` //库存
	RetailPrice           float64 `gorm:"NOT NULL;type:decimal(10,2)"`
}

func (p *Product) Get() error {
	return db.First(p).Error
}

func GetProducts(GoodsID uint) ([]Product, error) {
	products := make([]Product, 0)
	err := db.Where("goods_id=?", GoodsID).Find(&products).Error
	return products, err
}
