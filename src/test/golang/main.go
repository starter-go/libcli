package main

import (
	"bitwormhole.com/starter/libcli"
	"github.com/bitwormhole/starter"
)

func main() {
	i := starter.InitApp()
	i.UseMain(libcli.Module())
	i.Run()
}
