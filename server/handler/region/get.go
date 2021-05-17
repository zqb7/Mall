package region

import (
	"github.com/zqhhh/gomall/model"
	regionSrv "github.com/zqhhh/gomall/service/region"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {
	p := params{}
	err := c.ShouldBindQuery(&p)
	if err != nil {
		c.JSON(http.StatusOK, []struct{}{})
		return
	}
	ctx := regionSrv.Context{
		ParentID: p.ParentID,
		Regions:  []model.Region{},
	}
	service := regionSrv.NewRegionService()
	if err := service.Get(&ctx); err != nil {
		c.Set("error", err)
		return
	}
	c.JSON(http.StatusOK, ctx.Regions)
}
