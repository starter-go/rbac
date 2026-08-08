package main4rbac

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

    
    inst.register(&p6051122a59_unsupported_UnsupportedAuthenticationDAO{})
    inst.register(&p6051122a59_unsupported_UnsupportedDaoSetProvider{})
    inst.register(&p6051122a59_unsupported_UnsupportedPermissionDao{})
    inst.register(&p6051122a59_unsupported_UnsupportedRoleDao{})
    inst.register(&p6051122a59_unsupported_UnsupportedSessionDao{})
    inst.register(&p6051122a59_unsupported_UnsupportedUserDao{})
    inst.register(&p8b617a3f2c_core_RbacDaoSetLoader{})
    inst.register(&p8b617a3f2c_core_RbacDaoSetServiceImpl{})
    inst.register(&p9abb4d6558_mem_MemoryAuthenticationDAO{})
    inst.register(&p9abb4d6558_mem_MemoryDaoSetProvider{})
    inst.register(&p9abb4d6558_mem_MemoryEngineFacade{})
    inst.register(&p9abb4d6558_mem_MemoryPermissionDao{})
    inst.register(&p9abb4d6558_mem_MemoryRoleDao{})
    inst.register(&p9abb4d6558_mem_MemorySessionDao{})
    inst.register(&p9abb4d6558_mem_MemoryTableDao{})
    inst.register(&p9abb4d6558_mem_MemoryUserDao{})
    inst.register(&pa89018078d_agent_AuthentDaoAgent{})
    inst.register(&pa89018078d_agent_PermissionDaoAgent{})
    inst.register(&pa89018078d_agent_RoleDaoAgent{})
    inst.register(&pa89018078d_agent_SessionDaoAgent{})
    inst.register(&pa89018078d_agent_TableDaoAgent{})
    inst.register(&pa89018078d_agent_UserDaoAgent{})


    return nil
}
