// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cfg4libcli

import (
	cli0xf7c71e "bitwormhole.com/starter/cli"
	lib0x739556 "bitwormhole.com/starter/libcli/lib"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: cli.Client
	cominfobuilder.Next()
	cominfobuilder.ID("cli.Client").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComCliClientFacade{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: cli.ContextHolder
	cominfobuilder.Next()
	cominfobuilder.ID("cli.ContextHolder").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComCliContextHolderImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-lib0x739556.CliFiltersConfig
	cominfobuilder.Next()
	cominfobuilder.ID("com2-lib0x739556.CliFiltersConfig").Class("cli.FilterRegistry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComCliFiltersConfig{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-lib0x739556.CliHandlersConfig
	cominfobuilder.Next()
	cominfobuilder.ID("com3-lib0x739556.CliHandlersConfig").Class("cli.HandlerRegistry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComCliHandlersConfig{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: cli.Server
	cominfobuilder.Next()
	cominfobuilder.ID("cli.Server").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComCliServerFacade{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com5-lib0x739556.CliServiceImpl
	cominfobuilder.Next()
	cominfobuilder.ID("com5-lib0x739556.CliServiceImpl").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComCliServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com6-lib0x739556.TestPoint
	cominfobuilder.Next()
	cominfobuilder.ID("com6-lib0x739556.TestPoint").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTestPoint{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComCliClientFacade : the factory of component: cli.Client
type comFactory4pComCliClientFacade struct {

    mPrototype * lib0x739556.CliClientFacade

	
	mContextHolderSelector config.InjectionSelector

}

func (inst * comFactory4pComCliClientFacade) init() application.ComponentFactory {

	
	inst.mContextHolderSelector = config.NewInjectionSelector("#cli.ContextHolder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComCliClientFacade) newObject() * lib0x739556.CliClientFacade {
	return & lib0x739556.CliClientFacade {}
}

func (inst * comFactory4pComCliClientFacade) castObject(instance application.ComponentInstance) * lib0x739556.CliClientFacade {
	return instance.Get().(*lib0x739556.CliClientFacade)
}

func (inst * comFactory4pComCliClientFacade) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComCliClientFacade) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComCliClientFacade) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComCliClientFacade) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCliClientFacade) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCliClientFacade) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ContextHolder = inst.getterForFieldContextHolderSelector(context)
	return context.LastError()
}

//getterForFieldContextHolderSelector
func (inst * comFactory4pComCliClientFacade) getterForFieldContextHolderSelector (context application.InstanceContext) lib0x739556.CliContextHolder {

	o1 := inst.mContextHolderSelector.GetOne(context)
	o2, ok := o1.(lib0x739556.CliContextHolder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "cli.Client")
		eb.Set("field", "ContextHolder")
		eb.Set("type1", "?")
		eb.Set("type2", "lib0x739556.CliContextHolder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComCliContextHolderImpl : the factory of component: cli.ContextHolder
type comFactory4pComCliContextHolderImpl struct {

    mPrototype * lib0x739556.CliContextHolderImpl

	

}

func (inst * comFactory4pComCliContextHolderImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComCliContextHolderImpl) newObject() * lib0x739556.CliContextHolderImpl {
	return & lib0x739556.CliContextHolderImpl {}
}

func (inst * comFactory4pComCliContextHolderImpl) castObject(instance application.ComponentInstance) * lib0x739556.CliContextHolderImpl {
	return instance.Get().(*lib0x739556.CliContextHolderImpl)
}

func (inst * comFactory4pComCliContextHolderImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComCliContextHolderImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComCliContextHolderImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComCliContextHolderImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCliContextHolderImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCliContextHolderImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComCliFiltersConfig : the factory of component: com2-lib0x739556.CliFiltersConfig
type comFactory4pComCliFiltersConfig struct {

    mPrototype * lib0x739556.CliFiltersConfig

	
	mContextHolderSelector config.InjectionSelector

}

func (inst * comFactory4pComCliFiltersConfig) init() application.ComponentFactory {

	
	inst.mContextHolderSelector = config.NewInjectionSelector("#cli.ContextHolder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComCliFiltersConfig) newObject() * lib0x739556.CliFiltersConfig {
	return & lib0x739556.CliFiltersConfig {}
}

func (inst * comFactory4pComCliFiltersConfig) castObject(instance application.ComponentInstance) * lib0x739556.CliFiltersConfig {
	return instance.Get().(*lib0x739556.CliFiltersConfig)
}

func (inst * comFactory4pComCliFiltersConfig) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComCliFiltersConfig) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComCliFiltersConfig) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComCliFiltersConfig) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCliFiltersConfig) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCliFiltersConfig) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ContextHolder = inst.getterForFieldContextHolderSelector(context)
	return context.LastError()
}

//getterForFieldContextHolderSelector
func (inst * comFactory4pComCliFiltersConfig) getterForFieldContextHolderSelector (context application.InstanceContext) lib0x739556.CliContextHolder {

	o1 := inst.mContextHolderSelector.GetOne(context)
	o2, ok := o1.(lib0x739556.CliContextHolder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com2-lib0x739556.CliFiltersConfig")
		eb.Set("field", "ContextHolder")
		eb.Set("type1", "?")
		eb.Set("type2", "lib0x739556.CliContextHolder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComCliHandlersConfig : the factory of component: com3-lib0x739556.CliHandlersConfig
type comFactory4pComCliHandlersConfig struct {

    mPrototype * lib0x739556.CliHandlersConfig

	

}

func (inst * comFactory4pComCliHandlersConfig) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComCliHandlersConfig) newObject() * lib0x739556.CliHandlersConfig {
	return & lib0x739556.CliHandlersConfig {}
}

func (inst * comFactory4pComCliHandlersConfig) castObject(instance application.ComponentInstance) * lib0x739556.CliHandlersConfig {
	return instance.Get().(*lib0x739556.CliHandlersConfig)
}

func (inst * comFactory4pComCliHandlersConfig) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComCliHandlersConfig) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComCliHandlersConfig) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComCliHandlersConfig) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCliHandlersConfig) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCliHandlersConfig) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComCliServerFacade : the factory of component: cli.Server
type comFactory4pComCliServerFacade struct {

    mPrototype * lib0x739556.CliServerFacade

	
	mContextHolderSelector config.InjectionSelector

}

func (inst * comFactory4pComCliServerFacade) init() application.ComponentFactory {

	
	inst.mContextHolderSelector = config.NewInjectionSelector("#cli.ContextHolder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComCliServerFacade) newObject() * lib0x739556.CliServerFacade {
	return & lib0x739556.CliServerFacade {}
}

func (inst * comFactory4pComCliServerFacade) castObject(instance application.ComponentInstance) * lib0x739556.CliServerFacade {
	return instance.Get().(*lib0x739556.CliServerFacade)
}

func (inst * comFactory4pComCliServerFacade) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComCliServerFacade) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComCliServerFacade) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComCliServerFacade) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCliServerFacade) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCliServerFacade) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ContextHolder = inst.getterForFieldContextHolderSelector(context)
	return context.LastError()
}

//getterForFieldContextHolderSelector
func (inst * comFactory4pComCliServerFacade) getterForFieldContextHolderSelector (context application.InstanceContext) lib0x739556.CliContextHolder {

	o1 := inst.mContextHolderSelector.GetOne(context)
	o2, ok := o1.(lib0x739556.CliContextHolder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "cli.Server")
		eb.Set("field", "ContextHolder")
		eb.Set("type1", "?")
		eb.Set("type2", "lib0x739556.CliContextHolder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComCliServiceImpl : the factory of component: com5-lib0x739556.CliServiceImpl
type comFactory4pComCliServiceImpl struct {

    mPrototype * lib0x739556.CliServiceImpl

	
	mContextHolderSelector config.InjectionSelector
	mClientSelector config.InjectionSelector
	mServerSelector config.InjectionSelector
	mHandlersSelector config.InjectionSelector
	mFiltersSelector config.InjectionSelector

}

func (inst * comFactory4pComCliServiceImpl) init() application.ComponentFactory {

	
	inst.mContextHolderSelector = config.NewInjectionSelector("#cli.ContextHolder",nil)
	inst.mClientSelector = config.NewInjectionSelector("#cli.Client",nil)
	inst.mServerSelector = config.NewInjectionSelector("#cli.Server",nil)
	inst.mHandlersSelector = config.NewInjectionSelector(".cli.HandlerRegistry",nil)
	inst.mFiltersSelector = config.NewInjectionSelector(".cli.FilterRegistry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComCliServiceImpl) newObject() * lib0x739556.CliServiceImpl {
	return & lib0x739556.CliServiceImpl {}
}

func (inst * comFactory4pComCliServiceImpl) castObject(instance application.ComponentInstance) * lib0x739556.CliServiceImpl {
	return instance.Get().(*lib0x739556.CliServiceImpl)
}

func (inst * comFactory4pComCliServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComCliServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComCliServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComCliServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCliServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComCliServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ContextHolder = inst.getterForFieldContextHolderSelector(context)
	obj.Client = inst.getterForFieldClientSelector(context)
	obj.Server = inst.getterForFieldServerSelector(context)
	obj.Handlers = inst.getterForFieldHandlersSelector(context)
	obj.Filters = inst.getterForFieldFiltersSelector(context)
	return context.LastError()
}

//getterForFieldContextHolderSelector
func (inst * comFactory4pComCliServiceImpl) getterForFieldContextHolderSelector (context application.InstanceContext) lib0x739556.CliContextHolder {

	o1 := inst.mContextHolderSelector.GetOne(context)
	o2, ok := o1.(lib0x739556.CliContextHolder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com5-lib0x739556.CliServiceImpl")
		eb.Set("field", "ContextHolder")
		eb.Set("type1", "?")
		eb.Set("type2", "lib0x739556.CliContextHolder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldClientSelector
func (inst * comFactory4pComCliServiceImpl) getterForFieldClientSelector (context application.InstanceContext) cli0xf7c71e.Client {

	o1 := inst.mClientSelector.GetOne(context)
	o2, ok := o1.(cli0xf7c71e.Client)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com5-lib0x739556.CliServiceImpl")
		eb.Set("field", "Client")
		eb.Set("type1", "?")
		eb.Set("type2", "cli0xf7c71e.Client")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldServerSelector
func (inst * comFactory4pComCliServiceImpl) getterForFieldServerSelector (context application.InstanceContext) cli0xf7c71e.Server {

	o1 := inst.mServerSelector.GetOne(context)
	o2, ok := o1.(cli0xf7c71e.Server)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com5-lib0x739556.CliServiceImpl")
		eb.Set("field", "Server")
		eb.Set("type1", "?")
		eb.Set("type2", "cli0xf7c71e.Server")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldHandlersSelector
func (inst * comFactory4pComCliServiceImpl) getterForFieldHandlersSelector (context application.InstanceContext) []cli0xf7c71e.HandlerRegistry {
	list1 := inst.mHandlersSelector.GetList(context)
	list2 := make([]cli0xf7c71e.HandlerRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(cli0xf7c71e.HandlerRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

//getterForFieldFiltersSelector
func (inst * comFactory4pComCliServiceImpl) getterForFieldFiltersSelector (context application.InstanceContext) []cli0xf7c71e.FilterRegistry {
	list1 := inst.mFiltersSelector.GetList(context)
	list2 := make([]cli0xf7c71e.FilterRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(cli0xf7c71e.FilterRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTestPoint : the factory of component: com6-lib0x739556.TestPoint
type comFactory4pComTestPoint struct {

    mPrototype * lib0x739556.TestPoint

	
	mEnabledSelector config.InjectionSelector
	mClientSelector config.InjectionSelector

}

func (inst * comFactory4pComTestPoint) init() application.ComponentFactory {

	
	inst.mEnabledSelector = config.NewInjectionSelector("${cli.module.test.enabled}",nil)
	inst.mClientSelector = config.NewInjectionSelector("#cli.Client",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTestPoint) newObject() * lib0x739556.TestPoint {
	return & lib0x739556.TestPoint {}
}

func (inst * comFactory4pComTestPoint) castObject(instance application.ComponentInstance) * lib0x739556.TestPoint {
	return instance.Get().(*lib0x739556.TestPoint)
}

func (inst * comFactory4pComTestPoint) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTestPoint) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTestPoint) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTestPoint) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestPoint) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestPoint) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Enabled = inst.getterForFieldEnabledSelector(context)
	obj.Client = inst.getterForFieldClientSelector(context)
	return context.LastError()
}

//getterForFieldEnabledSelector
func (inst * comFactory4pComTestPoint) getterForFieldEnabledSelector (context application.InstanceContext) bool {
    return inst.mEnabledSelector.GetBool(context)
}

//getterForFieldClientSelector
func (inst * comFactory4pComTestPoint) getterForFieldClientSelector (context application.InstanceContext) cli0xf7c71e.Client {

	o1 := inst.mClientSelector.GetOne(context)
	o2, ok := o1.(cli0xf7c71e.Client)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com6-lib0x739556.TestPoint")
		eb.Set("field", "Client")
		eb.Set("type1", "?")
		eb.Set("type2", "cli0xf7c71e.Client")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}




