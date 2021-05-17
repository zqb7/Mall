package model

//专题
type Topic struct {
	Model
	Title           string  `json:"title" gorm:"not null; default:''"`
	Content         string  `json:"content" gorm:"type:text"`
	Avatar          string  `json:"avatar" gorm:"not null; default:''"`
	ItemPicUrl      string  `json:"item_pic_url" gorm:"not null; default:''"`
	Subtitle        string  `json:"subtitle" gorm:"not null; default:''"`
	TopicCategoryId uint    `json:"topic_category_id" gorm:"int(11) unsigned; not null; default: '0'"`
	PriceInfo       float64 `json:"price_info" gorm:"type:decimal(10,2); not null;default:'0.00'"`
	ReadCount       string  `json:"read_count" gorm:"not null; default:'0'"`
	ScenePicUrl     string  `json:"scene_pic_url" gorm:"not null; default: '0'"`
	TopicTemplateId uint    `json:"topic_template_id" gorm:"int(11) unsigned; not null; default: '0'"`
	TopicTagId      int     `json:"topic_tag_id" gorm:"int(11) unsigned; not null; default: '0'"`
	SortOrder       uint    `json:"sort_order" gorm:"int(11) unsigned; not null; default: '100'"`
	IsShow          uint    `json:"is_show" gorm:"not null; default:'1'"`
}

func GetTopic(selectFields []string, where map[string]interface{}, limit, offset int) ([]Topic, error) {
	topic := make([]Topic, 0)
	err := db.Model(Topic{}).Select(selectFields).Where(where).
		Limit(limit).Offset(offset).Find(&topic).Error
	return topic, err
}

func CountTopic() (count int) {
	_ = db.Model(Topic{}).Select("*").Count(&count).Error
	return count
}
