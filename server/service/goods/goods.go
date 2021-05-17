package goods

import "github.com/zqhhh/gomall/model"

type service struct {
}

type Service interface {
	Count() int
	Detail(ctx *Context) error
}

func NewGoodsService() Service {
	return service{}
}

func (s service) Count() int {
	goods := model.Goods{}
	return goods.Count()
}
func (s service) Detail(ctx *Context) error {
	h := &next{}
	h.setNext(&Detail{}).
		setNext(&Gallery{}).
		setNext(&Issue{}).
		setNext(&Attribute{}).
		setNext(&Specification{}).
		setNext(&Product{})
	if err := h.Run(ctx); err != nil {
		return err
	}
	return nil
}
