package lib

import (
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// TestPoint ...
type TestPoint struct {
	markup.Component `class:"life"`

	Enabled bool       `inject:"${cli.module.test.enabled}"`
	Client  cli.Client `inject:"#cli.Client"`
}

func (inst *TestPoint) _Impl() application.LifeRegistry {
	return inst
}

// GetLifeRegistration ...
func (inst *TestPoint) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnStart: inst.doTest,
	}
}

func (inst *TestPoint) doTest() error {
	if !inst.Enabled {
		return nil
	}
	cmd := "pwd"
	return inst.Client.RunCCA(nil, cmd, nil)
}
