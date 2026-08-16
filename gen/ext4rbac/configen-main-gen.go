package ext4rbac

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p52f05ffa69_iauth_PasswordAuthenticator{})
    inst.register(&p6a368d1706_iservices_CheckerServiceImpl{})
    inst.register(&p6a368d1706_iservices_TableServiceImpl{})
    inst.register(&p6a368d1706_iservices_UserServiceImpl{})


    return nil
}
