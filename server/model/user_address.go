package model

import "github.com/jinzhu/gorm"

type UserAddress struct {
	Model
	UserID    uint   `json:"user_id"`
	Consignee string `json:"consignee"`
	Country   string `json:"country" gorm:"NOT NULL;DEFAULT:''"`
	Province  string `json:"province" gorm:"NOT NULL;DEFAULT:''"`
	City      string `json:"city" gorm:"NOT NULL;DEFAULT:''"`
	District  string `json:"district" gorm:"NOT NULL;DEFAULT:''"`
	Address   string `json:"address"`
	Mobile    string `json:"mobile" gorm:"NOT NULL;type:varchar(60)"`
	IsDefault bool   `json:"is_default" gorm:"NOT NULL;DEFAULT:0"`
}

func (address *UserAddress) Get() error {
	return db.Where("user_id=?", address.UserID).First(address).Error
}

func (address *UserAddress) Insert() error {
	return db.Create(address).Error
}

func (address *UserAddress) Update(data map[string]interface{}) error {
	return db.Model(UserAddress{}).Where("id=?", address.ID).UpdateColumns(data).Error
}

func (address *UserAddress) GetDefault() error {
	return db.Where("user_id=? and is_default = ?", address.UserID, address.IsDefault).First(address).Error
}

func ListAddress(userID uint) []UserAddress {
	data := make([]UserAddress, 0)
	_ = db.Model(UserAddress{}).Find(&data).Error
	return data
}

func (address *UserAddress) AfterUpdate(tx *gorm.DB) (err error) {
	if address.IsDefault {
		tx.Model(&UserAddress{}).Where("user_id = ?", address.UserID).
			Not("id = ?", address.ID).Update("is_default", false)
	}
	return
}
