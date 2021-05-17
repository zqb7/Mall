package model

/*
SKU 库存量单位， 销售属性
*/

type SpecificationKey struct {
	Model
	Name      string `gorm:"NOT NULL"`
	SortOrder uint   `gorm:"type:tinyint(1);default:0;NOT NULL"`
}

type SpecificationValue struct {
	Model
	Value           string `gorm:"NOT NULL"`
	SpecificationID uint   `gorm:"NOT NULL"`
	GoodsID         uint   `gorm:"NOT NULL"`
	PicUrl          string `gorm:"NOT NULL;type:text"`
}
