package model

type Issue struct {
	Model
	GoodsID  uint
	Question string
	Answer   string
}

func GetIssues(goodsID uint) ([]Issue, error) {
	var issues []Issue
	err := db.Model(Issue{}).Where("goods_id=?", goodsID).Find(&issues).Error
	return issues, err
}
