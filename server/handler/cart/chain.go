package cart

import "github.com/zqhhh/gomall/model"

type Context struct {
	goodsID uint
	goods   model.Goods
}

type Handler interface {
	setNext(handler Handler) Handler
	Do(ctx *Context) error
	Run(ctx *Context) error
}

type Next struct {
	nextHandler Handler
}

func (n *Next) setNext(h Handler) Handler {
	n.nextHandler = h
	return h
}

func (n *Next) Run(ctx *Context) error {
	if n.nextHandler != nil {
		if err := n.nextHandler.Do(ctx); err != nil {
			return err
		}
		return n.nextHandler.Run(ctx)
	}
	return nil
}

type NilHandler struct {
	Next
}

func (h *NilHandler) Do(ctx *Context) error {
	return nil
}

type GoodsHandler struct {
	Next
}

func (h *GoodsHandler) Do(ctx *Context) error {
	goods := model.Goods{}
	goods.ID = ctx.goodsID
	err := goods.Get([]string{"*"})
	ctx.goods = goods
	return err
}
