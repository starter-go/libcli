package gen4libcli
import (
    p1336d65ed "github.com/starter-go/cli"
    p729331934 "github.com/starter-go/libcli/lib"
     "github.com/starter-go/application"
)

// type p729331934.CliClientFacade in package:github.com/starter-go/libcli/lib
//
// id:com-72933193420cd966-lib-CliClientFacade
// class:
// alias:alias-1336d65edeed550b78a5d5b61e92d726-Client
// scope:singleton
//
type p7293319342_lib_CliClientFacade struct {
}

func (inst* p7293319342_lib_CliClientFacade) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-72933193420cd966-lib-CliClientFacade"
	r.Classes = ""
	r.Aliases = "alias-1336d65edeed550b78a5d5b61e92d726-Client"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7293319342_lib_CliClientFacade) new() any {
    return &p729331934.CliClientFacade{}
}

func (inst* p7293319342_lib_CliClientFacade) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p729331934.CliClientFacade)
	nop(ie, com)

	
    com.ContextHolder = inst.getContextHolder(ie)


    return nil
}


func (inst*p7293319342_lib_CliClientFacade) getContextHolder(ie application.InjectionExt)p729331934.CliContextHolder{
    return ie.GetComponent("#alias-72933193420cd9662ad9c3dd31580c00-CliContextHolder").(p729331934.CliContextHolder)
}



// type p729331934.CliContextHolderImpl in package:github.com/starter-go/libcli/lib
//
// id:com-72933193420cd966-lib-CliContextHolderImpl
// class:
// alias:alias-72933193420cd9662ad9c3dd31580c00-CliContextHolder
// scope:singleton
//
type p7293319342_lib_CliContextHolderImpl struct {
}

func (inst* p7293319342_lib_CliContextHolderImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-72933193420cd966-lib-CliContextHolderImpl"
	r.Classes = ""
	r.Aliases = "alias-72933193420cd9662ad9c3dd31580c00-CliContextHolder"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7293319342_lib_CliContextHolderImpl) new() any {
    return &p729331934.CliContextHolderImpl{}
}

func (inst* p7293319342_lib_CliContextHolderImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p729331934.CliContextHolderImpl)
	nop(ie, com)

	


    return nil
}



// type p729331934.CliFiltersConfig in package:github.com/starter-go/libcli/lib
//
// id:com-72933193420cd966-lib-CliFiltersConfig
// class:class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry
// alias:
// scope:singleton
//
type p7293319342_lib_CliFiltersConfig struct {
}

func (inst* p7293319342_lib_CliFiltersConfig) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-72933193420cd966-lib-CliFiltersConfig"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7293319342_lib_CliFiltersConfig) new() any {
    return &p729331934.CliFiltersConfig{}
}

func (inst* p7293319342_lib_CliFiltersConfig) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p729331934.CliFiltersConfig)
	nop(ie, com)

	
    com.ContextHolder = inst.getContextHolder(ie)


    return nil
}


func (inst*p7293319342_lib_CliFiltersConfig) getContextHolder(ie application.InjectionExt)p729331934.CliContextHolder{
    return ie.GetComponent("#alias-72933193420cd9662ad9c3dd31580c00-CliContextHolder").(p729331934.CliContextHolder)
}



// type p729331934.CliHandlersConfig in package:github.com/starter-go/libcli/lib
//
// id:com-72933193420cd966-lib-CliHandlersConfig
// class:class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry
// alias:
// scope:singleton
//
type p7293319342_lib_CliHandlersConfig struct {
}

func (inst* p7293319342_lib_CliHandlersConfig) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-72933193420cd966-lib-CliHandlersConfig"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7293319342_lib_CliHandlersConfig) new() any {
    return &p729331934.CliHandlersConfig{}
}

func (inst* p7293319342_lib_CliHandlersConfig) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p729331934.CliHandlersConfig)
	nop(ie, com)

	


    return nil
}



// type p729331934.CliServerFacade in package:github.com/starter-go/libcli/lib
//
// id:com-72933193420cd966-lib-CliServerFacade
// class:
// alias:alias-1336d65edeed550b78a5d5b61e92d726-Server
// scope:singleton
//
type p7293319342_lib_CliServerFacade struct {
}

func (inst* p7293319342_lib_CliServerFacade) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-72933193420cd966-lib-CliServerFacade"
	r.Classes = ""
	r.Aliases = "alias-1336d65edeed550b78a5d5b61e92d726-Server"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7293319342_lib_CliServerFacade) new() any {
    return &p729331934.CliServerFacade{}
}

func (inst* p7293319342_lib_CliServerFacade) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p729331934.CliServerFacade)
	nop(ie, com)

	
    com.ContextHolder = inst.getContextHolder(ie)


    return nil
}


func (inst*p7293319342_lib_CliServerFacade) getContextHolder(ie application.InjectionExt)p729331934.CliContextHolder{
    return ie.GetComponent("#alias-72933193420cd9662ad9c3dd31580c00-CliContextHolder").(p729331934.CliContextHolder)
}



// type p729331934.CliServiceImpl in package:github.com/starter-go/libcli/lib
//
// id:com-72933193420cd966-lib-CliServiceImpl
// class:class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle
// alias:
// scope:singleton
//
type p7293319342_lib_CliServiceImpl struct {
}

func (inst* p7293319342_lib_CliServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-72933193420cd966-lib-CliServiceImpl"
	r.Classes = "class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7293319342_lib_CliServiceImpl) new() any {
    return &p729331934.CliServiceImpl{}
}

func (inst* p7293319342_lib_CliServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p729331934.CliServiceImpl)
	nop(ie, com)

	
    com.ContextHolder = inst.getContextHolder(ie)
    com.Client = inst.getClient(ie)
    com.Server = inst.getServer(ie)
    com.Handlers = inst.getHandlers(ie)
    com.Filters = inst.getFilters(ie)


    return nil
}


func (inst*p7293319342_lib_CliServiceImpl) getContextHolder(ie application.InjectionExt)p729331934.CliContextHolder{
    return ie.GetComponent("#alias-72933193420cd9662ad9c3dd31580c00-CliContextHolder").(p729331934.CliContextHolder)
}


func (inst*p7293319342_lib_CliServiceImpl) getClient(ie application.InjectionExt)p1336d65ed.Client{
    return ie.GetComponent("#alias-1336d65edeed550b78a5d5b61e92d726-Client").(p1336d65ed.Client)
}


func (inst*p7293319342_lib_CliServiceImpl) getServer(ie application.InjectionExt)p1336d65ed.Server{
    return ie.GetComponent("#alias-1336d65edeed550b78a5d5b61e92d726-Server").(p1336d65ed.Server)
}


func (inst*p7293319342_lib_CliServiceImpl) getHandlers(ie application.InjectionExt)[]p1336d65ed.HandlerRegistry{
    dst := make([]p1336d65ed.HandlerRegistry, 0)
    src := ie.ListComponents(".class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry")
    for _, item1 := range src {
        item2 := item1.(p1336d65ed.HandlerRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p7293319342_lib_CliServiceImpl) getFilters(ie application.InjectionExt)[]p1336d65ed.FilterRegistry{
    dst := make([]p1336d65ed.FilterRegistry, 0)
    src := ie.ListComponents(".class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry")
    for _, item1 := range src {
        item2 := item1.(p1336d65ed.FilterRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p729331934.TestPoint in package:github.com/starter-go/libcli/lib
//
// id:com-72933193420cd966-lib-TestPoint
// class:class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle
// alias:
// scope:singleton
//
type p7293319342_lib_TestPoint struct {
}

func (inst* p7293319342_lib_TestPoint) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-72933193420cd966-lib-TestPoint"
	r.Classes = "class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p7293319342_lib_TestPoint) new() any {
    return &p729331934.TestPoint{}
}

func (inst* p7293319342_lib_TestPoint) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p729331934.TestPoint)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.CLI = inst.getCLI(ie)


    return nil
}


func (inst*p7293319342_lib_TestPoint) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${cli.module.test.enabled}")
}


func (inst*p7293319342_lib_TestPoint) getCLI(ie application.InjectionExt)p1336d65ed.CLI{
    return ie.GetComponent("#alias-1336d65edeed550b78a5d5b61e92d726-CLI").(p1336d65ed.CLI)
}


