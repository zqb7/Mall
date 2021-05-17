package cart

import (
	"github.com/zqhhh/gomall/objs"
	cartSrv "github.com/zqhhh/gomall/service/cart"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Del(c *gin.Context) {
	user := objs.BuilderUser(c)
	params := Params{}
	if err := c.ShouldBindUri(&params); err != nil {
		c.Set("error", err)
		return
	}
	service := cartSrv.NewCartService()
	_ = service.Delete(params.ID, uint(user.ID))
	c.JSON(http.StatusOK, gin.H{})
}
