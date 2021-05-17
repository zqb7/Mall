package address

import (
	"github.com/zqhhh/gomall/model"
	"github.com/zqhhh/gomall/objs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(c *gin.Context) {
	user := objs.BuilderUser(c)
	isDefault := c.DefaultQuery("is_default", "")
	if isDefault == "1" {
		address := model.UserAddress{}
		address.UserID = uint(user.ID)
		address.IsDefault = true
		_ = address.GetDefault()
		c.JSON(http.StatusOK, []interface{}{address})
		return
	}
	c.JSON(http.StatusOK, model.ListAddress(uint(user.ID)))
}
