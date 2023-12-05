package main

import (
	"os"

	"github.com/starter-go/libcli"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(os.Args)
	i.MainModule(libcli.ModuleTest())
	i.WithPanic(true).Run()
}
