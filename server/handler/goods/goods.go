package goods

import (
	goodsSrv "github.com/zqhhh/gomall/service/goods"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Detail(c *gin.Context) {
	p := params{}
	if err := c.ShouldBindUri(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	var ctx goodsSrv.Context
	ctx.GoodsID = p.ID
	service := goodsSrv.NewGoodsService()
	if err := service.Detail(&ctx); err != nil {

	}
	res := ResGoodsDetail{
		Goods:             customGoods{},
		Gallery:           []customGallery{},
		Issue:             []customIssue{},
		SpecificationList: []specification{},
	}

	res.Goods = customGoods{
		ID:         ctx.GoodsDetail.ID,
		Name:       ctx.GoodsDetail.Name,
		GoodsBrief: ctx.GoodsDetail.GoodsBrief,
		GoodsDesc:  ctx.GoodsDetail.GoodsDesc,
		ListPicUrl: ctx.GoodsDetail.ListPicUrl,
	}
	for _, g := range ctx.GoodsGalleries {
		res.Gallery = append(res.Gallery, customGallery{
			ID:     g.ID,
			ImgUrl: g.ImgUrl,
		})
	}

	for _, i := range ctx.GoodsIssue {
		res.Issue = append(res.Issue, customIssue{
			ID:       i.ID,
			Question: i.Question,
			Answer:   i.Answer,
		})
	}
	res.Attribute = ctx.GoodsAttr

	speMap := make(map[uint]int)
	for _, s := range ctx.GoodsSpec {
		if i, ok := speMap[s.SpecKeyID]; ok {
			res.SpecificationList[i].ValueList = append(res.SpecificationList[i].ValueList, specificationValue{
				ID:                 s.ID,
				GoodsID:            s.GoodsID,
				Value:              s.Value,
				PicUrl:             s.PicUrl,
				Name:               s.Name,
				SpecificationKeyID: s.SpecKeyID,
			})
		} else {
			spec := specification{
				ID:        s.SpecKeyID,
				Name:      s.Name,
				ValueList: []specificationValue{},
			}
			spec.ValueList = append(spec.ValueList, specificationValue{
				ID:                 s.ID,
				GoodsID:            s.GoodsID,
				Value:              s.Value,
				PicUrl:             s.PicUrl,
				Name:               s.Name,
				SpecificationKeyID: s.SpecKeyID,
			})
			res.SpecificationList = append(res.SpecificationList, spec)
			speMap[s.SpecKeyID] = len(res.SpecificationList) - 1
		}
	}

	res.ProductList = ctx.Products
	c.JSON(http.StatusOK, res)
}
