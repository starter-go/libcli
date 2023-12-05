package lib

import (
	"github.com/starter-go/application"
	"github.com/starter-go/cli"
)

// TestPoint ...
type TestPoint struct {
	// markup.Component `class:"life"`
	//starter:component
	_as func(application.Lifecycle) //starter:as(".")

	Enabled bool    //starter:inject("${cli.module.test.enabled}")
	CLI     cli.CLI //starter:inject("#")
}

func (inst *TestPoint) _Impl() application.Lifecycle {
	return inst
}

// Life ...
func (inst *TestPoint) Life() *application.Life {
	return &application.Life{
		OnStart: inst.doTest,
	}
}

func (inst *TestPoint) doTest() error {
	if !inst.Enabled {
		return nil
	}
	cmd := "pwd"
	return inst.CLI.GetClient().RunCCA(nil, cmd, nil)
}
