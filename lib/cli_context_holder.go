package lib

import (
	"github.com/starter-go/cli"
)

// CliContextHolder ...
type CliContextHolder interface {
	GetContext() *cli.Context
}

// CliContextHolderImpl ...
type CliContextHolderImpl struct {
	// markup.Component `id:"cli.ContextHolder"`
	//starter:component
	_as func(CliContextHolder) //starter:as("#")

	context *cli.Context
}

func (inst *CliContextHolderImpl) _Impl() CliContextHolder {
	return inst
}

// GetContext ...
func (inst *CliContextHolderImpl) GetContext() *cli.Context {
	ctx := inst.context
	if ctx == nil {
		ctx = &cli.Context{}
		inst.context = ctx
	}
	return ctx
}
