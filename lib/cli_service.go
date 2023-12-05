package lib

import (
	"github.com/starter-go/application"
	"github.com/starter-go/cli"
)

////////////////////////////////////////////////////////////////////////////////

// CliServiceImpl ...
type CliServiceImpl struct {
	// markup.Component `class:"life"`
	//starter:component
	_as func(application.Lifecycle) //starter:as(".")

	ContextHolder CliContextHolder //starter:inject("#")
	Client        cli.Client       //starter:inject("#")
	Server        cli.Server       //starter:inject("#")

	Handlers []cli.HandlerRegistry //starter:inject(".")
	Filters  []cli.FilterRegistry  //starter:inject(".")
}

func (inst *CliServiceImpl) _Impl() application.Lifecycle {
	return inst
}

// Life ...
func (inst *CliServiceImpl) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.init,
	}
}

func (inst *CliServiceImpl) init() error {
	ctx := inst.ContextHolder.GetContext()
	inst.initContext(ctx)
	return nil
}

func (inst *CliServiceImpl) initContext(ctx *cli.Context) {

	// cb := cli.DefaultContextFactory{}
	// handlers := inst.Handlers
	// filters := inst.Filters

	// cb.Init(ctx)

	// for _, h := range handlers {
	// 	cb.RegisterHandler(h)
	// }

	// for _, f := range filters {
	// 	cb.RegisterFilter(f)
	// }

	// // cb.Create()
	// cb.NewContext()

	ctx.Client = inst.Client
	ctx.Server = inst.Server

}
