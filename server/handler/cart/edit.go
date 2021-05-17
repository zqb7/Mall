package cart

import (
	"github.com/zqhhh/gomall/objs"
	cartSrv "github.com/zqhhh/gomall/service/cart"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Update(c *gin.Context) {
	user := objs.BuilderUser(c)
	params := Params{}
	if err := c.ShouldBindUri(&params); err != nil {
		c.Set("error", err)
		return
	}
	req := ReqUpdateNumber{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("error", err)
		return
	}
	service := cartSrv.NewCartService()
	if err := service.UpdateNumber(params.ID, uint(user.ID), req.Number); err != nil {
		c.Set("error", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
