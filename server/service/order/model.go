package order

type CustomOrder struct {
	OrderID                 uint    `json:"order_id"`
	OrderSn                 int64   `json:"order_sn"`
	OrderStatus             uint    `json:"order_status"`
	GoodsID                 uint    `json:"goods_id"`
	GoodsName               string  `json:"goods_name"`
	GoodsRetailPrice        float64 `json:"goods_retail_price"`
	Number                  uint    `json:"number"`
	RetailPrice             float64 `json:"retail_price"`
	GoodsSpecificationValue string  `json:"goods_specification_value"`
	ListPicUrl              string  `json:"list_pic_url"`
}
