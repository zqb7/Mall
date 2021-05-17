package address

import (
	"github.com/zqhhh/gomall/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

func Update(c *gin.Context) {
	req := ReqUpdate{}
	uriParams := UriParams{}
	if err := c.ShouldBindUri(&uriParams); err != nil {
		c.Set("error", err)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Set("error", err)
		return
	}
	data := structToMap(req)
	address := model.UserAddress{}
	address.ID = uriParams.AddressID
	err := address.Update(data)
	if err != nil {

	}
	c.JSON(http.StatusOK, gin.H{})
}

func structToMap(obj interface{}) (data map[string]interface{}) {
	data = make(map[string]interface{}, 0)
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("json")
		if tag == "" {
			data[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
		} else if strings.Contains(tag, "omitempty") && v.Field(i).Interface() == nil {
			continue
		} else {
			data[strings.Split(tag, ",")[0]] = v.Field(i).Interface()
		}

	}
	return data
}
