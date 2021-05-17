package cart

import (
	"github.com/zqhhh/gomall/objs"
	cartSrv "github.com/zqhhh/gomall/service/cart"
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(c *gin.Context) {
	user := objs.BuilderUser(c)
	service := cartSrv.NewCartService()
	cartList := make([]customCart, 0)
	carts, _ := service.GetMyCart(uint(user.ID))
	for _, c := range carts {
		cartList = append(cartList, customCart{
			ID:                      c.ID,
			GoodsID:                 c.GoodsID,
			GoodsName:               c.GoodsName,
			Image:                   c.GoodsPic,
			Number:                  c.Number,
			Checked:                 c.Checked,
			RetailPrice:             c.RetailPrice,
			GoodsSpecificationValue: c.GoodsSpecificationValue,
		})
	}
	c.JSON(http.StatusOK, cartList)
}
