package goods

import (
	"fmt"
	"github.com/zqhhh/gomall/model"
)

type Context struct {
	GoodsID        uint
	GoodsDetail    model.Goods
	GoodsGalleries []model.Gallery
	GoodsIssue     []model.Issue
	GoodsAttr      []Attribute
	GoodsSpec      []Specification
	Products       []Product
}
type IHandler interface {
	setNext(h IHandler) IHandler
	Do(ctx *Context) error
	Run(ctx *Context) error
}

type next struct {
	nextHandler IHandler
}

func (n *next) setNext(h IHandler) IHandler {
	n.nextHandler = h
	return h
}

func (n *next) Do(ctx *Context) error {
	return nil
}

func (n *next) Run(ctx *Context) error {
	if n.nextHandler != nil {
		if err := n.nextHandler.Do(ctx); err != nil {
			return err
		}
		return n.nextHandler.Run(ctx)
	}
	return nil
}

type Detail struct {
	next
}

func (d *Detail) Do(ctx *Context) error {
	goods := model.Goods{}
	goods.ID = ctx.GoodsID
	if err := goods.Get([]string{"id", "name", "goods_brief", "goods_desc", "list_pic_url"}); err != nil {
		return err
	}
	ctx.GoodsDetail = goods
	return nil
}

type Gallery struct {
	next
}

func (g *Gallery) Do(ctx *Context) error {
	galleries, err := model.GetGalleries(ctx.GoodsDetail.ID)
	if err != nil {
		return err
	}
	ctx.GoodsGalleries = galleries
	return nil
}

type Issue struct {
	next
}

func (i *Issue) Do(ctx *Context) error {
	issue, _ := model.GetIssues(ctx.GoodsID)
	ctx.GoodsIssue = issue
	return nil
}

type Attribute struct {
	next
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (a *Attribute) Do(ctx *Context) error {
	sql := fmt.Sprintf("SELECT k.name, v.value FROM attribute_value as v "+
		"LEFT JOIN attribute_key as k ON k.id = v.attribute_id "+
		"WHERE v.goods_id = %d ORDER BY k.sort_order ASC", ctx.GoodsID)
	attributes := make([]Attribute, 0)
	if err := model.DB.Raw(sql).Scan(&attributes).Error; err != nil {
		return err
	}
	ctx.GoodsAttr = attributes
	return nil
}

type Specification struct {
	next
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	PicUrl    string `json:"pic_url"`
	GoodsID   uint   `json:"goods_id"`
	SpecKeyID uint   `json:"spec_key_id"`
}

func (s *Specification) Do(ctx *Context) error {
	sql := fmt.Sprintf("SELECT k.name,k.id as spec_key_id,v.* FROM specification_value as v "+
		"LEFT JOIN specification_key as k ON k.id = v.specification_id "+
		"WHERE v.goods_id = %d ORDER BY k.sort_order ASC", ctx.GoodsID)
	specification := make([]Specification, 0)
	if err := model.DB.Raw(sql).Scan(&specification).Error; err != nil {
		return err
	}
	ctx.GoodsSpec = specification
	return nil
}

type Product struct {
	next
	ID                    uint    `json:"id"`
	GoodsID               uint    `json:"goods_id"`
	GoodsSn               string  `json:"goods_sn"`
	GoodsSpecificationIDs string  `json:"goods_specification_ids"`
	GoodsNumber           uint    `json:"goods_number"`
	RetailPrice           float64 `json:"retail_price"`
}

func (p *Product) Do(ctx *Context) error {
	products, err := model.GetProducts(ctx.GoodsID)
	if err != nil {
		return err
	}
	for _, p := range products {
		ctx.Products = append(ctx.Products, Product{
			ID:                    p.ID,
			GoodsID:               p.GoodsID,
			GoodsSn:               p.GoodsSn,
			GoodsSpecificationIDs: p.GoodsSpecificationIDs,
			GoodsNumber:           p.GoodsNumber,
			RetailPrice:           p.RetailPrice,
		})
	}
	return nil
}
