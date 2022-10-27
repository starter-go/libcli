// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cfg4libcli

import (
	cli0xf7c71e "bitwormhole.com/starter/cli"
	lib0x739556 "bitwormhole.com/starter/libcli/lib"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComCliClientFacade struct {
	instance *lib0x739556.CliClientFacade
	 markup0x23084a.Component `id:"cli.Client"`
	ContextHolder lib0x739556.CliContextHolder `inject:"#cli.ContextHolder"`
}


type pComCliContextHolderImpl struct {
	instance *lib0x739556.CliContextHolderImpl
	 markup0x23084a.Component `id:"cli.ContextHolder"`
}


type pComCliFiltersConfig struct {
	instance *lib0x739556.CliFiltersConfig
	 markup0x23084a.Component `class:"cli.FilterRegistry"`
	ContextHolder lib0x739556.CliContextHolder `inject:"#cli.ContextHolder"`
}


type pComCliHandlersConfig struct {
	instance *lib0x739556.CliHandlersConfig
	 markup0x23084a.Component `class:"cli.HandlerRegistry"`
}


type pComCliServerFacade struct {
	instance *lib0x739556.CliServerFacade
	 markup0x23084a.Component `id:"cli.Server"`
	ContextHolder lib0x739556.CliContextHolder `inject:"#cli.ContextHolder"`
}


type pComCliServiceImpl struct {
	instance *lib0x739556.CliServiceImpl
	 markup0x23084a.Component `class:"life"`
	ContextHolder lib0x739556.CliContextHolder `inject:"#cli.ContextHolder"`
	Client cli0xf7c71e.Client `inject:"#cli.Client"`
	Server cli0xf7c71e.Server `inject:"#cli.Server"`
	Handlers []cli0xf7c71e.HandlerRegistry `inject:".cli.HandlerRegistry"`
	Filters []cli0xf7c71e.FilterRegistry `inject:".cli.FilterRegistry"`
}


type pComTestPoint struct {
	instance *lib0x739556.TestPoint
	 markup0x23084a.Component `class:"life"`
	Enabled bool `inject:"${cli.module.test.enabled}"`
	Client cli0xf7c71e.Client `inject:"#cli.Client"`
}

