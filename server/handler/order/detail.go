package order

import "github.com/gin-gonic/gin"

func Detail(c *gin.Context) {
	params := UriParams{}
	if err := c.ShouldBindUri(&params); err != nil {
		c.Set("error", err)
		return
	}
}
