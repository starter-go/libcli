package lib

import (
	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/cli/filters"
	"github.com/bitwormhole/starter/markup"
)

// CliFiltersConfig ...
type CliFiltersConfig struct {
	markup.Component `class:"cli.FilterRegistry"`

	ContextHolder CliContextHolder `inject:"#cli.ContextHolder"`
}

func (inst *CliFiltersConfig) _Impl() cli.FilterRegistry {
	return inst
}

// GetFilters ...
func (inst *CliFiltersConfig) GetFilters() []*cli.FilterRegistration {

	ctx := inst.ContextHolder.GetContext()

	frList1 := make([]cli.FilterRegistry, 0)
	frList1 = append(frList1, &filters.BindingFilter{})
	frList1 = append(frList1, &filters.ClientServerInjectingFilter{Context: ctx})
	frList1 = append(frList1, &filters.CommandPrepareFilter{})
	frList1 = append(frList1, &filters.ErrorFilter{})
	frList1 = append(frList1, &filters.HandlingFilter{})
	frList1 = append(frList1, &filters.MultilineCommandFilter{})

	return inst.makeRegistrations(frList1)
}

func (inst *CliFiltersConfig) makeRegistrations(src []cli.FilterRegistry) []*cli.FilterRegistration {
	dst := make([]*cli.FilterRegistration, 0)
	for _, fr1 := range src {
		group := fr1.GetFilters()
		dst = append(dst, group...)
	}
	return dst
}
