package cart

type ReqAdd struct {
	GoodsID                 uint   `json:"goods_id" binding:"required"`
	Number                  uint   `json:"number"`
	GoodsSn                 string `json:"goods_sn" binding:"required"`
	GoodsName               string `json:"goods_name" binding:"required"`
	GoodsPic                string `json:"goods_pic" binding:"required"`
	GoodsSpecificationValue string `json:"goods_specification_value" binding:"required"`
	GoodsSpecificationIDs   string `json:"goods_specification_ids" binding:"required"`
}

type customCart struct {
	ID                      uint    `json:"id"`
	GoodsID                 uint    `json:"goods_id"`
	GoodsName               string  `json:"goods_name"`
	Image                   string  `json:"image"`
	Number                  uint    `json:"number"`
	Checked                 bool    `json:"checked"`
	RetailPrice             float64 `json:"retail_price"`
	GoodsSpecificationValue string  `json:"goods_specification_value"`
}

type Params struct {
	ID uint `uri:"id" binding:"required,min=1"`
}

type ReqChecked struct {
	Checked bool `json:"checked"`
}

type ReqUpdateNumber struct {
	Number uint `json:"number"`
}
