package order

import (
	"github.com/zqhhh/gomall/objs"
	orderSrv "github.com/zqhhh/gomall/service/order"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(c *gin.Context) {
	req := ReqOrder{GoodsList: []goods{}}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("error", err)
		return
	}
	user := objs.BuilderUser(c)
	ctx := orderSrv.NewContext()
	ctx.UserID = uint(user.ID)
	for _, r := range req.GoodsList {
		ctx.ProductIds = append(ctx.ProductIds, r.ProductID)
		ctx.GoodsList = append(ctx.GoodsList, orderSrv.Goods{
			GoodsID:                 r.GoodsID,
			GoodsName:               r.GoodsName,
			GoodsSn:                 r.GoodsSn,
			ProductID:               r.ProductID,
			ByNumber:                r.ByNumber,
			GoodsSpecificationValue: r.GoodsSpecificationValue,
			GoodsSpecificationIDs:   r.GoodsSpecificationIDs,
			ListPicUrl:              r.ListPicUrl,
		})
	}
	service := orderSrv.NewOrderService()
	if err := service.Create(&ctx); err != nil {
		c.Set("error", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"order_sn": ctx.OrderSn,
		"price":    ctx.TotalPrice,
		"pay":      gin.H{}, //这里存放支付需要的参数
	})
}
