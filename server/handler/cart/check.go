package cart

import (
	"github.com/zqhhh/gomall/objs"
	cartSrv "github.com/zqhhh/gomall/service/cart"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Checked(c *gin.Context) {
	user := objs.BuilderUser(c)
	params := Params{}
	if err := c.ShouldBindUri(&params); err != nil {
		c.Set("error", err)
		return
	}

	req := ReqChecked{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("error", err)
		return
	}
	service := cartSrv.NewCartService()
	_ = service.UpdateChecked(params.ID, uint(user.ID), req.Checked)
	c.JSON(http.StatusOK, gin.H{})
}
