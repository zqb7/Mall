package catalog

import (
	"github.com/zqhhh/gomall/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(c *gin.Context) {
	res := ResIndex{}
	parentCategories, _ := model.GetCategories([]string{"id", "name", "wap_banner_url"}, map[string]interface{}{"parent_id": 0})
	res.CategoryList = make([]customCategory, 0)
	for _, p := range parentCategories {
		sub := make([]subCategory, 0)
		subC, _ := model.GetCategories([]string{"id", "name", "front_name", "wap_banner_url"}, map[string]interface{}{"parent_id": p.ID})
		for _, s := range subC {
			sub = append(sub, subCategory{
				ID:           s.ID,
				Name:         s.Name,
				WapBannerUrl: s.WapBannerUrl,
			})
		}
		res.CategoryList = append(res.CategoryList, customCategory{
			ID:              p.ID,
			Name:            p.Name,
			WapBannerUrl:    p.WapBannerUrl,
			SubCategoryList: sub,
		})
	}
	if len(res.CategoryList) > 0 {
		res.CurrentCategory = res.CategoryList[0]
	}
	c.JSON(http.StatusOK, res)
}

func Current(c *gin.Context) {
	id := c.Param("id")
	current := customCategory{
		ID:              0,
		Name:            "",
		WapBannerUrl:    "",
		SubCategoryList: make([]subCategory, 0),
	}
	category := model.Category{}
	err := category.Get([]string{"id", "name", "front_name", "wap_banner_url"}, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		c.JSON(http.StatusOK, struct{}{})
		return
	}
	current.ID = category.ID
	current.Name = category.Name
	current.FrontName = category.FrontName
	current.WapBannerUrl = category.WapBannerUrl
	sub := make([]subCategory, 0)
	subC, _ := model.GetCategories([]string{"id", "name", "front_name", "wap_banner_url"}, map[string]interface{}{"parent_id": category.ID})
	for _, s := range subC {
		sub = append(sub, subCategory{
			ID:           s.ID,
			Name:         s.Name,
			FrontName:    s.FrontName,
			WapBannerUrl: s.WapBannerUrl,
		})
	}
	current.SubCategoryList = sub
	c.JSON(http.StatusOK, gin.H{
		"current_category": current,
	})
}
