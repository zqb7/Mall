package objs

import "github.com/gin-gonic/gin"

func BuilderUser(c *gin.Context) *User {
	userI, ok := c.Get("user")
	if ok {
		user, _ := userI.(*User)
		return user
	}
	return nil
}
