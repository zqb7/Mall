package index

import (
	"github.com/zqhhh/gomall/model"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	res := Response{
		Ads:     make([]model.AD, 0),
		Channel: make([]model.Channel, 0),
	}
	floorList := make([]floor, 0)
	res.Ads, _ = model.GetADs()
	res.Channel, _ = model.GetChannels()
	res.Brands, _ = model.GetLastBrand()
	parentCategories, _ := model.GetCategories([]string{"id"}, map[string]interface{}{"parent_id": 0})
	for _, p := range parentCategories {
		resCate := floor{
			ID:        p.ID,
			Name:      p.Name,
			GoodsList: make([]goods, 0),
		}
		childCategories, _ := model.GetCategories([]string{"id"}, map[string]interface{}{"parent_id": p.ID})
		var childCID []uint
		for _, c := range childCategories {
			childCID = append(childCID, c.ID)
		}
		goodsList, _ := model.GetGoodsByCategoryIDs([]string{"id", "name", "list_pic_url", "retail_price"}, childCID, 7, 0)
		for _, g := range goodsList {
			resCate.GoodsList = append(resCate.GoodsList, goods{
				ID:          g.ID,
				Name:        g.Name,
				ListPicUrl:  g.ListPicUrl,
				RetailPrice: g.RetailPrice,
			})
		}
		floorList = append(floorList, resCate)
	}
	hotGoodsList, _ := model.GetGoods([]string{"id", "name", "list_pic_url", "retail_price", "goods_brief"}, map[string]interface{}{
		"is_hot": 1,
	}, 3, 0)
	newGoodsList, _ := model.GetGoods([]string{"id", "name", "list_pic_url", "retail_price"}, map[string]interface{}{
		"is_new": 1,
	}, 4, 0)
	res.HotGoodsList = make([]goods, 0)
	res.NewGoodsList = make([]goods, 0)
	for _, h := range hotGoodsList {
		res.HotGoodsList = append(res.HotGoodsList, goods{
			ID:          h.ID,
			Name:        h.Name,
			ListPicUrl:  h.ListPicUrl,
			RetailPrice: h.RetailPrice,
			GoodsBrief:  h.GoodsBrief,
		})
	}
	for _, n := range newGoodsList {
		res.NewGoodsList = append(res.NewGoodsList, goods{
			ID:          n.ID,
			Name:        n.Name,
			ListPicUrl:  n.ListPicUrl,
			RetailPrice: n.RetailPrice,
		})
	}
	res.FloorList = floorList
	res.TopicList, _ = model.GetTopic([]string{"*"}, map[string]interface{}{}, 3, 0)
	c.JSON(200, res)
}
