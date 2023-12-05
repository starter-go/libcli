package gen4libclitest
import (
    p1336d65ed "github.com/starter-go/cli"
    pfd93c80ff "github.com/starter-go/libcli/src/test/golang/unit"
     "github.com/starter-go/application"
)

// type pfd93c80ff.Test1 in package:github.com/starter-go/libcli/src/test/golang/unit
//
// id:com-fd93c80ff30567a7-unit-Test1
// class:
// alias:
// scope:singleton
//
type pfd93c80ff3_unit_Test1 struct {
}

func (inst* pfd93c80ff3_unit_Test1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fd93c80ff30567a7-unit-Test1"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfd93c80ff3_unit_Test1) new() any {
    return &pfd93c80ff.Test1{}
}

func (inst* pfd93c80ff3_unit_Test1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfd93c80ff.Test1)
	nop(ie, com)

	
    com.CLI = inst.getCLI(ie)


    return nil
}


func (inst*pfd93c80ff3_unit_Test1) getCLI(ie application.InjectionExt)p1336d65ed.CLI{
    return ie.GetComponent("#alias-1336d65edeed550b78a5d5b61e92d726-CLI").(p1336d65ed.CLI)
}


