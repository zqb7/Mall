package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

func HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err, ok := c.Get("error")
		if !ok {
			return
		}
		errList := make(map[string]string, 0)
		switch errs := err.(type) {
		case validator.ValidationErrors:
			for _, e := range errs {
				errList[strings.ToLower(e.Field())] = validationErrorToText(e)
			}
		default:
			errList["body"] = "非法的请求体"
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": errList,
		})
	}
}
func validationErrorToText(e validator.FieldError) string {
	fieldName := strings.ToLower(e.Field())
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s 是必选参数", fieldName)
	case "max":
		return fmt.Sprintf("%s 长度必须小于或等于 %s", fieldName, e.Param())
	case "min":
		return fmt.Sprintf("%s 长度必须大于或等于 %s", fieldName, e.Param())
	case "email":
		return fmt.Sprintf("邮箱格式错误")
	case "len":
		return fmt.Sprintf("%s 长度必须为 %s", fieldName, e.Param())
	}
	return fmt.Sprintf("无效的参数 %s", fieldName)
}
