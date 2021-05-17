package model

/*
商品SPU 关键属性, 不会影响库存和价格的属性
*/

type AttributeKey struct {
	Model
	Name      string `gorm:"NOT NULL"`
	SortOrder uint   `gorm:"type:tinyint(1);default:0;NOT NULL"`
}

type AttributeValue struct {
	Model
	Value       string `gorm:"NOT NULL;type:text"`
	AttributeID uint   `gorm:"NOT NULL"`
	GoodsID     uint   `gorm:"NOT NULL"`
}
