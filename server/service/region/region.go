package region

import "github.com/zqhhh/gomall/model"

type Context struct {
	ParentID uint
	Regions  []model.Region
}
type Service interface {
	Get(ctx *Context) error
}

func NewRegionService() Service {
	return service{}
}

type service struct {
}

func (s service) Get(ctx *Context) error {
	ctx.Regions = model.GetRegions(ctx.ParentID)
	return nil
}
