package cart

import (
	cartSrv "github.com/zqhhh/gomall/service/cart"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Clear(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetHeader("user_id"))
	service := cartSrv.NewCartService()
	_ = service.Clear(uint(userID))
	c.JSON(http.StatusOK, gin.H{})
}
