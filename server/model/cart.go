package model

type Cart struct {
	Model
	UserID                  uint
	GoodsID                 uint
	GoodsSn                 string
	GoodsName               string
	GoodsPic                string  `gorm:"NOT NULL"`
	Checked                 bool    `gorm:"NOT NULL;default:0"`
	Number                  uint    `gorm:"NOT NULL; default:1"`
	MarketPrice             float64 `gorm:"type:decimal(10,2)"`
	RetailPrice             float64 `gorm:"type:decimal(10,2)"`
	GoodsSpecificationValue string  `gorm:"type:text"`
	GoodsSpecificationIDs   string  `gorm:"type:varchar(100)"`
}

func (cart *Cart) Insert() error {
	return db.Create(cart).Error
}

func (cart *Cart) UpdateFields(fields map[string]interface{}) error {
	return db.Model(cart).Where("user_id=?", cart.UserID).UpdateColumns(fields).Error
}

func GetUserCart(selectFields []string, userID uint) (data []Cart, err error) {
	err = db.Select(selectFields).Where("user_id=?", userID).Find(&data).Error
	return data, err
}

func (cart *Cart) Delete() error {
	err := db.Where(map[string]interface{}{
		"id":      cart.ID,
		"user_id": cart.UserID,
	}).Delete(Cart{}).Error
	return err
}

func (cart *Cart) Clear() error {
	err := db.Where(map[string]interface{}{
		"user_id": cart.UserID,
	}).Delete(Cart{}).Error
	return err
}
