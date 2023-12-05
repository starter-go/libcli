package unit

import (
	"context"

	"github.com/starter-go/application"
	"github.com/starter-go/cli"
)

// Test1 ...
type Test1 struct {
	//starter:component
	//// _as func(application.Lifecycle) //starter:as(".")

	CLI cli.CLI //starter:inject("#")
}

func (inst *Test1) _impl() application.Lifecycle {
	return inst
}

// Life ...
func (inst *Test1) Life() *application.Life {
	return &application.Life{
		OnStart: inst.start,
	}
}

func (inst *Test1) start() error {
	ctx := context.Background()
	return inst.CLI.GetClient().RunCCA(ctx, "help", nil)
}
