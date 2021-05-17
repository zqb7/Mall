package middleware

import (
	"github.com/zqhhh/gomall/objs"
	"github.com/zqhhh/gomall/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token不能为空",
			})
		} else {
			claim, err := util.ParseToken(token)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
			} else {
				c.Set("user_id", claim.UserID)
				c.Request.Header.Set("user_id", strconv.Itoa(int(claim.UserID)))
				c.Set("user", &objs.User{ID: int(claim.UserID)})
				c.Next()
			}
		}

	}
}
