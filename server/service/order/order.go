package order

import (
	"fmt"
	"github.com/zqhhh/gomall/model"
)

type Service interface {
	Create(ctx *context) error
	List(userID, limit, offset uint) (data []CustomOrder, err error)
	Get(selectedField []string, orderSn int64) (model.Order, error)
}

func NewOrderService() Service {
	return service{}
}

type service struct {
}

func (s service) Create(ctx *context) error {
	h := &next{}
	tx := model.DB.Begin()
	h.setNext(&Address{}).
		setNext(&StockSubtract{}).
		setNext(&Price{}).
		setNext(&DbOrder{tx: tx}).
		setNext(&DbOrderSku{tx: tx})
	if err := h.Run(ctx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
func (s service) List(userID, limit, offset uint) (data []CustomOrder, err error) {
	sql := fmt.Sprintf("SELECT o.id as order_id,o.order_status, o.order_price,o.order_sn,og.goods_id,"+
		"og.goods_name,og.number,og.list_pic_url,og.retail_price as goods_retail_price,og.goods_specification_value FROM order_goods as og "+
		"LEFT JOIN `order` as o on o.id=og.id WHERE o.user_id=%d", userID)
	err = model.DB.Raw(sql).Scan(&data).Error
	return data, err
}

func (s service) Get(selectedField []string, orderSn int64) (model.Order, error) {
	order := model.Order{}
	err := model.DB.Select(selectedField).Where("order_sn=?", orderSn).First(&order).Error
	return order, err
}
