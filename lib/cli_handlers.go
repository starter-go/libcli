package lib

import (
	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/cli/handlers"

	"github.com/bitwormhole/starter/markup"
)

// CliHandlersConfig ...
type CliHandlersConfig struct {
	markup.Component `class:"cli.HandlerRegistry"`
}

func (inst *CliHandlersConfig) _Impl() cli.HandlerRegistry {
	return inst
}

// GetHandlers ...
func (inst *CliHandlersConfig) GetHandlers() []*cli.HandlerRegistration {

	list1 := make([]cli.HandlerRegistry, 0)

	list1 = append(list1, &handlers.ChdirHandler{})
	list1 = append(list1, &handlers.HelpHandler{})
	list1 = append(list1, &handlers.LsHandler{})
	list1 = append(list1, &handlers.MkdirHandler{})
	list1 = append(list1, &handlers.PwdHandler{})
	list1 = append(list1, &handlers.SleepHandler{})

	return inst.makeRegistrationList(list1)
}

func (inst *CliHandlersConfig) makeRegistrationList(src []cli.HandlerRegistry) []*cli.HandlerRegistration {
	dst := make([]*cli.HandlerRegistration, 0)
	for _, hr1 := range src {
		group := hr1.GetHandlers()
		dst = append(dst, group...)
	}
	return dst
}
