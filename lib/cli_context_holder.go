package lib

import (
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/starter/markup"
)

// CliContextHolder ...
type CliContextHolder interface {
	GetContext() *cli.Context
}

// CliContextHolderImpl ...
type CliContextHolderImpl struct {
	markup.Component `id:"cli.ContextHolder"`

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
