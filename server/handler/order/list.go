package order

import (
	"github.com/zqhhh/gomall/objs"
	orderSrv "github.com/zqhhh/gomall/service/order"
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(c *gin.Context) {
	user := objs.BuilderUser(c)
	service := orderSrv.NewOrderService()
	data, err := service.List(uint(user.ID), 0, 0)
	if err != nil {
		c.Set("error", err)
		return
	}
	res := make([]CustomOrder, 0)
	m := make(map[int64]int, 0)
	for _, d := range data {
		if v, ok := m[d.OrderSn]; ok {
			res[v].GoodsList = append(res[v].GoodsList, CustomOrderGoods{
				GoodsID:                 d.GoodsID,
				GoodsName:               d.GoodsName,
				Number:                  d.Number,
				OrderSn:                 d.OrderSn,
				RetailPrice:             d.GoodsRetailPrice,
				GoodsSpecificationValue: d.GoodsSpecificationValue,
				ListPicUrl:              d.ListPicUrl,
			})
		} else {
			res = append(res, CustomOrder{
				OrderID:     d.OrderID,
				OrderSn:     d.OrderSn,
				OrderStatus: d.OrderStatus,
				RetailPrice: d.RetailPrice,
				GoodsList: []CustomOrderGoods{{
					GoodsID:                 d.GoodsID,
					GoodsName:               d.GoodsName,
					Number:                  d.Number,
					OrderSn:                 d.OrderSn,
					RetailPrice:             d.GoodsRetailPrice,
					GoodsSpecificationValue: d.GoodsSpecificationValue,
					ListPicUrl:              d.ListPicUrl,
				}},
			})
			m[d.OrderSn] = len(m)
		}
	}
	c.JSON(http.StatusOK, res)
}
