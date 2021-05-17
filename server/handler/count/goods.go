package count

import (
	goodsSrv "github.com/zqhhh/gomall/service/goods"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Goods(c *gin.Context) {
	service := goodsSrv.NewGoodsService()
	c.JSON(http.StatusOK, gin.H{
		"goods_count": service.Count(),
	})
}
