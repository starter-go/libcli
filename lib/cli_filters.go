package lib

import (
	"github.com/starter-go/cli"
	"github.com/starter-go/cli/filters"
)

// CliFiltersConfig ...
type CliFiltersConfig struct {
	// markup.Component `class:"cli.FilterRegistry"`
	//starter:component
	_as func(cli.FilterRegistry) //starter:as(".")

	ContextHolder CliContextHolder //starter:inject("#")
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
