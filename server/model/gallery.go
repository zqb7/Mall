package model

//画廊， 商品详情顶部的轮播图片
type Gallery struct {
	Model
	GoodsID   uint `gorm:"type:int(11) unsigned;default:0;NOT NULL"`
	ImgUrl    string
	SortOrder uint `gorm:"type:tinyint(1);default:0;NOT NULL"`
}

func GetGalleries(goodsID uint) ([]Gallery, error) {
	var goods []Gallery
	err := db.Model(Gallery{}).Select([]string{"id", "goods_id", "img_url"}).
		Where("goods_id=?", goodsID).Find(&goods).Error
	return goods, err
}
