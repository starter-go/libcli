package gen4libclitest

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportConfig ...
func ExportConfig(cb application.ComponentRegistry) error {
	return registerComponents(cb)
}
