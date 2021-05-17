package address

import (
	"github.com/zqhhh/gomall/model"
	"github.com/zqhhh/gomall/objs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Add(c *gin.Context) {
	user := objs.BuilderUser(c)
	req := ReqAddress{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("error", err)
		return
	}
	userAddr := model.UserAddress{
		UserID:    uint(user.ID),
		Consignee: req.Consignee,
		Country:   req.Country,
		Province:  req.Province,
		City:      req.City,
		District:  req.District,
		Address:   req.Address,
		Mobile:    req.Mobile,
		IsDefault: req.IsDefault,
	}
	_ = userAddr.Insert()
	c.JSON(http.StatusOK, gin.H{})
}
