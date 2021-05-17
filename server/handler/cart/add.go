package cart

import (
	"github.com/zqhhh/gomall/model"
	"github.com/zqhhh/gomall/objs"
	cartSrv "github.com/zqhhh/gomall/service/cart"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Add(c *gin.Context) {
	user := objs.BuilderUser(c)
	req := ReqAdd{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("error", err)
		return
	}
	var (
		h   Handler
		ctx Context
	)
	ctx.goodsID = req.GoodsID
	h = &NilHandler{}
	h.setNext(&GoodsHandler{})
	if err := h.Run(&ctx); err != nil {
		c.Set("error", err)
		return
	}
	service := cartSrv.NewCartService()
	cart := model.Cart{}
	cart.Checked = true
	{
		cart.UserID = uint(user.ID)
		cart.GoodsID = req.GoodsID
		cart.Number = req.Number
		cart.GoodsPic = ctx.goods.ListPicUrl
		cart.RetailPrice = ctx.goods.RetailPrice
		cart.GoodsName = ctx.goods.Name
		cart.GoodsSn = req.GoodsSn
		cart.GoodsSpecificationIDs = req.GoodsSpecificationIDs
		cart.GoodsSpecificationValue = req.GoodsSpecificationValue
	}
	if err := service.Add(cart); err != nil {
		c.Set("error", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
