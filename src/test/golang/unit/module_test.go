package unit

import (
	"testing"

	"bitwormhole.com/starter/libcli"
	"github.com/bitwormhole/starter"
)

func TestModule(t *testing.T) {
	mod := libcli.Module()
	i := starter.InitApp()
	i.UseMain(mod)
	i.Run()
}
