package order

import (
	"github.com/zqhhh/gomall/model"
	"github.com/zqhhh/gomall/pkg/snowflake"
	"github.com/jinzhu/gorm"
)

type context struct {
	UserID     uint
	ProductIds []uint
	GoodsList  []Goods
	AddressID  uint
	OrderSn    int64

	orderID    uint
	address    model.UserAddress
	prices     map[uint]float64
	TotalPrice float64
}

func NewContext() context {
	return context{
		UserID:     0,
		ProductIds: []uint{},
		GoodsList:  []Goods{},
		AddressID:  0,
		orderID:    0,
		address:    model.UserAddress{},
		prices:     make(map[uint]float64),
		TotalPrice: 0,
	}
}

type Goods struct {
	GoodsID                 uint
	GoodsName               string
	GoodsSn                 string
	ProductID               uint
	ByNumber                uint
	GoodsSpecificationValue string
	GoodsSpecificationIDs   string
	ListPicUrl              string
}

type IHandler interface {
	setNext(h IHandler) IHandler
	Do(ctx *context) error
	Run(ctx *context) error
}

type next struct {
	nextHandler IHandler
}

func (n *next) setNext(h IHandler) IHandler {
	n.nextHandler = h
	return h
}

func (n *next) Do(_ *context) error {
	return nil
}

func (n *next) Run(ctx *context) error {
	if n.nextHandler != nil {
		if err := n.nextHandler.Do(ctx); err != nil {
			return err
		}
		return n.nextHandler.Run(ctx)
	}
	return nil
}

type Address struct {
	next
}

func (a Address) Do(ctx *context) error {
	address := model.UserAddress{}
	address.ID = ctx.AddressID
	address.UserID = ctx.UserID
	err := address.Get()
	ctx.address = address
	return err
}

type StockSubtract struct {
	next
}

// 减库存
func (s StockSubtract) Do(ctx *context) error {
	return nil
}

// 价格计算
type Price struct {
	next
}

func (p *Price) Do(ctx *context) error {
	for _, g := range ctx.GoodsList {
		product := model.Product{}
		product.ID = g.ProductID
		if err := product.Get(); err != nil {
			return err
		}
		ctx.TotalPrice += product.RetailPrice
		ctx.prices[g.ProductID] = product.RetailPrice
	}
	return nil
}

type DbOrder struct {
	tx *gorm.DB
	next
}

// 写订单表
func (o DbOrder) Do(ctx *context) error {
	address := ctx.address
	orderSn := snowflake.GetSnowflakeID()
	ctx.OrderSn = orderSn
	order := model.Order{
		OrderSn:        orderSn,
		UserID:         ctx.UserID,
		OrderStatus:    1,
		ShippingStatus: 0,
		PayStatus:      0,
		Consignee:      address.Consignee,
		Country:        address.Country,
		Province:       address.Province,
		City:           address.City,
		District:       address.District,
		Address:        address.Address,
		Mobile:         address.Mobile,
		ActualPrice:    0,
		OrderPrice:     ctx.TotalPrice,
	}
	err := o.tx.Create(&order).Error
	ctx.orderID = order.ID
	return err
}

// 写订单商品表
type DbOrderSku struct {
	tx *gorm.DB
	next
}

func (o DbOrderSku) Do(ctx *context) error {
	for _, g := range ctx.GoodsList {
		orderGoods := model.OrderGoods{
			OrderID:                 ctx.orderID,
			GoodsID:                 g.GoodsID,
			GoodsName:               g.GoodsName,
			GoodsSn:                 g.GoodsSn,
			ProductID:               g.ProductID,
			Number:                  g.ByNumber,
			MarketPrice:             0,
			RetailPrice:             ctx.prices[g.ProductID],
			GoodsSpecificationValue: g.GoodsSpecificationValue,
			GoodsSpecificationIDs:   g.GoodsSpecificationIDs,
			ListPicUrl:              g.ListPicUrl,
		}
		if err := o.tx.Create(&orderGoods).Error; err != nil {
			return err
		}
	}
	return nil
}
