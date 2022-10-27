package lib

import (
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/starter/markup"
)

// CliServerFacade ...
type CliServerFacade struct {
	markup.Component `id:"cli.Server"`

	ContextHolder CliContextHolder `inject:"#cli.ContextHolder"`
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
