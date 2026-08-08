package core4rbac
import (
    pa9adc578e "github.com/starter-go/rbac/api/classes/authx"
    p4026b8990 "github.com/starter-go/rbac/src/core/golang/corecom/iauthx"
     "github.com/starter-go/application"
)

// type p4026b8990.AuthxServiceImpl in package:github.com/starter-go/rbac/src/core/golang/corecom/iauthx
//
// id:com-4026b8990d646998-iauthx-AuthxServiceImpl
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-AuthService
// scope:singleton
//
type p4026b8990d_iauthx_AuthxServiceImpl struct {
}

func (inst* p4026b8990d_iauthx_AuthxServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4026b8990d646998-iauthx-AuthxServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-AuthService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4026b8990d_iauthx_AuthxServiceImpl) new() any {
    return &p4026b8990.AuthxServiceImpl{}
}

func (inst* p4026b8990d_iauthx_AuthxServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4026b8990.AuthxServiceImpl)
	nop(ie, com)

	
    com.RegistryList = inst.getRegistryList(ie)


    return nil
}


func (inst*p4026b8990d_iauthx_AuthxServiceImpl) getRegistryList(ie application.InjectionExt)[]pa9adc578e.Registry{
    dst := make([]pa9adc578e.Registry, 0)
    src := ie.ListComponents(".class-a9adc578e4b7f0bd1de2909d6ab953e5-Registry")
    for _, item1 := range src {
        item2 := item1.(pa9adc578e.Registry)
        dst = append(dst, item2)
    }
    return dst
}


