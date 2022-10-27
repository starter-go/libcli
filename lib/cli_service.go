package lib

import (
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

// CliServiceImpl ...
type CliServiceImpl struct {
	markup.Component `class:"life"`

	ContextHolder CliContextHolder `inject:"#cli.ContextHolder"`
	Client        cli.Client       `inject:"#cli.Client"`
	Server        cli.Server       `inject:"#cli.Server"`

	Handlers []cli.HandlerRegistry `inject:".cli.HandlerRegistry"`
	Filters  []cli.FilterRegistry  `inject:".cli.FilterRegistry"`
}

func (inst *CliServiceImpl) _Impl() application.LifeRegistry {
	return inst
}

// GetLifeRegistration ...
func (inst *CliServiceImpl) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnInit: inst.init,
	}
}

func (inst *CliServiceImpl) init() error {
	ctx := inst.ContextHolder.GetContext()
	inst.initContext(ctx)
	return nil
}

func (inst *CliServiceImpl) initContext(ctx *cli.Context) {

	cb := cli.ContextBuilder{}
	handlers := inst.Handlers
	filters := inst.Filters

	cb.Init(ctx)

	for _, h := range handlers {
		cb.RegisterHandler(h)
	}

	for _, f := range filters {
		cb.RegisterFilter(f)
	}

	cb.Create()
}
