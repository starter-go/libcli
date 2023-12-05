package libcli

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/cli/modules/cli"
	"github.com/starter-go/libcli/gen/gen4libcli"
	"github.com/starter-go/libcli/gen/gen4libclitest"
)

const (
	theModuleName     = "github.com/starter-go/libcli"
	theModuleVersion  = "v0.0.2"
	theModuleRevision = 2
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// Module 导出模块 [github.com/starter-go/libcli]
func Module() application.Module {

	mb := application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.Components(gen4libcli.ExportConfig)
	mb.EmbedResources(theModuleResFS, theModuleResPath)

	mb.Depend(cli.Module())

	return mb.Create()
}

// ModuleTest 导出模块 [github.com/starter-go/libcli#test]
func ModuleTest() application.Module {

	parent := Module()

	mb := application.ModuleBuilder{}
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.Components(gen4libclitest.ExportConfig)
	mb.EmbedResources(theModuleResFS, theModuleResPath)

	mb.Depend(parent)

	return mb.Create()
}
