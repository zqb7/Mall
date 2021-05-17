package index

import "github.com/zqhhh/gomall/model"

type Response struct {
	Ads          []model.AD      `json:"ads"`
	Channel      []model.Channel `json:"channel"`
	Brands       []model.Brand   `json:"brands"`
	NewGoodsList []goods         `json:"new_goods_list"`
	HotGoodsList []goods         `json:"hot_goods_list"`
	TopicList    []model.Topic   `json:"topic_list"`
	FloorList    []floor         `json:"floor_list"`
}

type floor struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	GoodsList []goods `json:"goods_list"`
}

type goods struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	ListPicUrl  string  `json:"list_pic_url"`
	RetailPrice float64 `json:"retail_price"`
	GoodsBrief  string  `json:"goods_brief,omitempty"`
}
