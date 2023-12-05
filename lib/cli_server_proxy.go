package lib

import (
	"github.com/starter-go/cli"
)

// CliServerFacade ...
type CliServerFacade struct {
	// markup.Component `id:"cli.Server"`
	//starter:component
	_as func(cli.Server) //starter:as("#")

	ContextHolder CliContextHolder //starter:inject("#")
}

func (inst *CliServerFacade) _Impl() cli.Server {
	return inst
}

func (inst *CliServerFacade) getInner() cli.Server {
	return inst.ContextHolder.GetContext().Server
}

// FindHandler ...
func (inst *CliServerFacade) FindHandler(name string) (*cli.HandlerRegistration, error) {
	return inst.getInner().FindHandler(name)
}

// ListNames ...
func (inst *CliServerFacade) ListNames() []string {
	return inst.getInner().ListNames()
}

// RegisterHandler 。。。
func (inst *CliServerFacade) RegisterHandler(hr *cli.HandlerRegistration) error {
	return inst.getInner().RegisterHandler(hr)
}

// RegisterHandlers 。。。
func (inst *CliServerFacade) RegisterHandlers(hr cli.HandlerRegistry) error {
	return inst.getInner().RegisterHandlers(hr)
}
