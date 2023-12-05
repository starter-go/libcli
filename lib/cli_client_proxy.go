package lib

import (
	"context"

	"github.com/starter-go/cli"
)

// CliClientFacade ...
type CliClientFacade struct {
	// markup.Component `id:"cli.Client"`
	//starter:component
	_as func(cli.Client) //starter:as("#")

	ContextHolder CliContextHolder //starter:inject("#")
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
