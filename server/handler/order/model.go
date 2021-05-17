package order

type ReqOrder struct {
	GoodsList []goods `json:"goods_list"`
	AddressID uint    `json:"address_id"`
}

type goods struct {
	GoodsID                 uint   `json:"goods_id" binding:"required"`
	GoodsName               string `json:"goods_name" binding:"required"`
	GoodsSn                 string `json:"goods_sn" binding:"required"`
	ProductID               uint   `json:"product_id" binding:"required"`
	ByNumber                uint   `json:"by_number" binding:"required"`
	GoodsSpecificationValue string `json:"goods_specification_value" binding:"required"`
	GoodsSpecificationIDs   string `json:"goods_specification_ids" binding:"required"`
	ListPicUrl              string `json:"list_pic_url" binding:"required"`
}

type UriParams struct {
	OrderSn int64 `uri:"orderSn"`
}

type CustomOrder struct {
	OrderID     uint               `json:"order_id"`
	OrderSn     int64              `json:"order_sn"`
	OrderStatus uint               `json:"order_status"`
	RetailPrice float64            `json:"retail_price"`
	GoodsList   []CustomOrderGoods `json:"goods_list"`
}

type CustomOrderGoods struct {
	GoodsID                 uint    `json:"goods_id"`
	GoodsName               string  `json:"goods_name"`
	Number                  uint    `json:"number"`
	OrderSn                 int64   `json:"order_sn"`
	RetailPrice             float64 `json:"retail_price"`
	GoodsSpecificationValue string  `json:"goods_specification_value"`
	ListPicUrl              string  `json:"list_pic_url"`
}
