package lib

import (
	"context"

	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/starter/markup"
)

// CliClientFacade ...
type CliClientFacade struct {
	markup.Component `id:"cli.Client"`

	ContextHolder CliContextHolder `inject:"#cli.ContextHolder"`
}

func (inst *CliClientFacade) _Impl() cli.Client {
	return inst
}

func (inst *CliClientFacade) getInner() cli.Client {
	return inst.ContextHolder.GetContext().Client
}

// RunCCA ...
func (inst *CliClientFacade) RunCCA(ctx context.Context, cmd string, args []string) error {
	return inst.getInner().RunCCA(ctx, cmd, args)
}

// Run ...
func (inst *CliClientFacade) Run(t *cli.Task) error {
	return inst.getInner().Run(t)
}
