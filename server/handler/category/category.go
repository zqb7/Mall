package category

import (
	"github.com/zqhhh/gomall/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Detail(c *gin.Context) {
	id := c.Param("id")
	category := model.Category{}
	err := category.Get([]string{"id", "name", "front_name", "wap_banner_url", "parent_id"}, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	res := ResCategory{
		CurrentCategory:     customCategory{},
		ParentCategory:      customCategory{},
		BrotherCategoryList: []customCategory{},
	}
	res.CurrentCategory = customCategory{
		ID:           category.ID,
		Name:         category.Name,
		FrontName:    category.FrontName,
		WapBannerUrl: category.WapBannerUrl,
	}
	parentCategory := model.Category{}
	_ = parentCategory.Get([]string{"id", "name", "front_name", "wap_banner_url", "parent_id"}, map[string]interface{}{
		"id": category.ParentID,
	})
	res.ParentCategory = customCategory{
		ID:           parentCategory.ID,
		Name:         parentCategory.Name,
		FrontName:    parentCategory.FrontName,
		WapBannerUrl: parentCategory.WapBannerUrl,
	}
	brothers, _ := model.GetCategories(
		[]string{"id", "name", "front_name", "wap_banner_url"},
		map[string]interface{}{"parent_id": parentCategory.ID})
	for _, b := range brothers {
		res.BrotherCategoryList = append(res.BrotherCategoryList, customCategory{
			ID:           b.ID,
			Name:         b.Name,
			FrontName:    b.FrontName,
			WapBannerUrl: b.WapBannerUrl,
		})
	}
	c.JSON(http.StatusOK, res)
}

func Goods(c *gin.Context) {
	res := ResGoods{GoodsList: []goods{}}
	goodsList, _ := model.GetGoods([]string{"id", "name", "list_pic_url", "retail_price"}, map[string]interface{}{
		"category_id": c.Param("id"),
	}, 10000, 0)
	for _, g := range goodsList {
		res.GoodsList = append(res.GoodsList, goods{
			ID:          g.ID,
			Name:        g.Name,
			ListPicUrl:  g.ListPicUrl,
			RetailPrice: g.RetailPrice,
		})
	}
	c.JSON(http.StatusOK, res)
}
