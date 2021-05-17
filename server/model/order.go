package model

// 订单表
type Order struct {
	Model
	OrderSn        int64  `gorm:"UNIQUE;NOT NULL"`
	UserID         uint   `gorm:"NOT NULL"`
	OrderStatus    uint   `gorm:"NOT NULL;DEFAULT:0"`
	ShippingStatus uint   `gorm:"NOT NULL;DEFAULT:0"` //发货状态
	PayStatus      uint   `gorm:"NOT NULL;DEFAULT:0"`
	Consignee      string // 收货人
	Country        string `gorm:"NOT NULL;DEFAULT:''"`
	Province       string `gorm:"NOT NULL;DEFAULT:''"`
	City           string `gorm:"NOT NULL;DEFAULT:''"`
	District       string `gorm:"NOT NULL;DEFAULT:''"`
	Address        string
	Mobile         string `gorm:"NOT NULL;type:varchar(60)"`
	ActualPrice    float64
	OrderPrice     float64
}

// 订单商品表
type OrderGoods struct {
	Model
	OrderID                 uint
	GoodsID                 uint
	GoodsName               string
	GoodsSn                 string
	ProductID               uint
	Number                  uint    `gorm:"NOT NULL; default:1"`
	MarketPrice             float64 `gorm:"type:decimal(10,2)"`
	RetailPrice             float64 `gorm:"type:decimal(10,2)"`
	GoodsSpecificationValue string  `gorm:"type:text"`
	GoodsSpecificationIDs   string  `gorm:"type:varchar(100)"`
	ListPicUrl              string
}

func (o *Order) Insert() error {
	return db.Create(o).Error
}
