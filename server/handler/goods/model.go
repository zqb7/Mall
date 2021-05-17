package goods

import (
	"github.com/zqhhh/gomall/service/goods"
)

type params struct {
	ID uint `uri:"id"`
}

type customGoods struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	GoodsBrief string `json:"goods_brief"`
	GoodsDesc  string `json:"goods_desc"`
	ListPicUrl string `json:"list_pic_url"`
}

type customGallery struct {
	ID     uint   `json:"id"`
	ImgUrl string `json:"img_url"`
}

type customIssue struct {
	ID       uint   `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type ResGoodsDetail struct {
	Goods             customGoods       `json:"goods"`
	Gallery           []customGallery   `json:"gallery"`
	Issue             []customIssue     `json:"issue"`
	Attribute         []goods.Attribute `json:"attribute"`
	SpecificationList []specification   `json:"specification_list"`
	ProductList       []goods.Product   `json:"product_list"`
}

type specificationValue struct {
	ID                 uint   `json:"id"`
	GoodsID            uint   `json:"goods_id"`
	Value              string `json:"value"`
	PicUrl             string `json:"pic_url"`
	Name               string `json:"name"`
	SpecificationKeyID uint   `json:"specification_key_id"`
}

type specification struct {
	ID        uint                 `json:"id"`
	Name      string               `json:"name"`
	ValueList []specificationValue `json:"value_list"`
}
