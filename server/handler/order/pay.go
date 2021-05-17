package order

import (
	orderSrv "github.com/zqhhh/gomall/service/order"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用于生成各个平台支付参数
func Pay(c *gin.Context) {
	params := UriParams{}
	if err := c.ShouldBindUri(&params); err != nil {
		c.Set("error", err)
		return
	}
	service := orderSrv.NewOrderService()
	order, err := service.Get([]string{"order_status", "retail_price"}, params.OrderSn)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"price": order.OrderPrice,
		"pay":   gin.H{}, //预留，返回的是支付宝等app需要的支付参数
	})
}
