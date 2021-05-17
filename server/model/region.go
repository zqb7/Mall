package model

type Region struct {
	Model
	ParentID uint   `gorm:"NOT NULL;default:0" json:"parent_id"`
	Name     string `gorm:"NOT NULL" json:"name"`
	Type     uint   `gorm:"NOT NULL;default:0" json:"type"`
}

func GetRegions(parentID uint) []Region {
	data := make([]Region, 0)
	_ = db.Model(Region{}).Where("parent_id=?", parentID).Find(&data).Error
	return data
}

func GetAllRegions() []Region {
	data := make([]Region, 0)
	_ = db.Model(Region{}).Find(&data).Error
	return data
}
