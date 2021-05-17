package auth

import (
	"github.com/zqhhh/gomall/pkg/util"
	userSrv "github.com/zqhhh/gomall/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	req := ReqLogin{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("error", err)
		return
	}
	service := userSrv.NewUserService()
	err := service.Login(req.User, req.Password)
	if err != nil {
		c.Set("error", err)
		return
	}
	token, err := util.GenerateToken(util.Claim{
		UserID:   service.GetUserID(),
		UserName: req.User,
		Password: req.Password,
	})
	if err != nil {
		c.Set("error", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
