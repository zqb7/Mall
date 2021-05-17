package auth

import (
	"github.com/zqhhh/gomall/model"
	userSrv "github.com/zqhhh/gomall/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var req ReqRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("error", err)
		return
	}
	user := model.User{
		Username:    req.Username,
		Password:    req.Password,
		Birthday:    req.Birthday,
		LastLoginIP: c.ClientIP(),
		Nickname:    "",
		Mobile:      "",
		RegisterIP:  c.ClientIP(),
		Avatar:      "",
	}
	service := userSrv.NewUserService()
	if err := service.Register(user); err != nil {
		c.Set("error", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
