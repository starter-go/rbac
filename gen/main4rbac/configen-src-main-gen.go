package main4rbac
import (
    p24287f458 "github.com/starter-go/rbac"
    p914b425f9 "github.com/starter-go/rbac/src/main/golang/librbac/iauthx"
    pa89018078 "github.com/starter-go/rbac/src/main/golang/librbac/idaoset/agent"
    p8b617a3f2 "github.com/starter-go/rbac/src/main/golang/librbac/idaoset/core"
    p9abb4d655 "github.com/starter-go/rbac/src/main/golang/librbac/idaoset/mem"
    p6051122a5 "github.com/starter-go/rbac/src/main/golang/librbac/idaoset/unsupported"
     "github.com/starter-go/application"
)

// type p914b425f9.AuthxServiceImpl in package:github.com/starter-go/rbac/src/main/golang/librbac/iauthx
//
// id:com-914b425f94c19b47-iauthx-AuthxServiceImpl
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-AuthService
// scope:singleton
//
type p914b425f94_iauthx_AuthxServiceImpl struct {
}

func (inst* p914b425f94_iauthx_AuthxServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-914b425f94c19b47-iauthx-AuthxServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-AuthService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p914b425f94_iauthx_AuthxServiceImpl) new() any {
    return &p914b425f9.AuthxServiceImpl{}
}

func (inst* p914b425f94_iauthx_AuthxServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p914b425f9.AuthxServiceImpl)
	nop(ie, com)

	


    return nil
}



// type pa89018078.AuthentDaoAgent in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/agent
//
// id:com-a89018078de55056-agent-AuthentDaoAgent
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-AuthenticationDAO
// scope:singleton
//
type pa89018078d_agent_AuthentDaoAgent struct {
}

func (inst* pa89018078d_agent_AuthentDaoAgent) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a89018078de55056-agent-AuthentDaoAgent"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-AuthenticationDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa89018078d_agent_AuthentDaoAgent) new() any {
    return &pa89018078.AuthentDaoAgent{}
}

func (inst* pa89018078d_agent_AuthentDaoAgent) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa89018078.AuthentDaoAgent)
	nop(ie, com)

	
    com.Serivce = inst.getSerivce(ie)


    return nil
}


func (inst*pa89018078d_agent_AuthentDaoAgent) getSerivce(ie application.InjectionExt)p24287f458.DaoSetService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-DaoSetService").(p24287f458.DaoSetService)
}



// type pa89018078.PermissionDaoAgent in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/agent
//
// id:com-a89018078de55056-agent-PermissionDaoAgent
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-PermissionDAO
// scope:singleton
//
type pa89018078d_agent_PermissionDaoAgent struct {
}

func (inst* pa89018078d_agent_PermissionDaoAgent) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a89018078de55056-agent-PermissionDaoAgent"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-PermissionDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa89018078d_agent_PermissionDaoAgent) new() any {
    return &pa89018078.PermissionDaoAgent{}
}

func (inst* pa89018078d_agent_PermissionDaoAgent) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa89018078.PermissionDaoAgent)
	nop(ie, com)

	
    com.Serivce = inst.getSerivce(ie)


    return nil
}


func (inst*pa89018078d_agent_PermissionDaoAgent) getSerivce(ie application.InjectionExt)p24287f458.DaoSetService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-DaoSetService").(p24287f458.DaoSetService)
}



// type pa89018078.RoleDaoAgent in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/agent
//
// id:com-a89018078de55056-agent-RoleDaoAgent
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-RoleDAO
// scope:singleton
//
type pa89018078d_agent_RoleDaoAgent struct {
}

func (inst* pa89018078d_agent_RoleDaoAgent) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a89018078de55056-agent-RoleDaoAgent"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-RoleDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa89018078d_agent_RoleDaoAgent) new() any {
    return &pa89018078.RoleDaoAgent{}
}

func (inst* pa89018078d_agent_RoleDaoAgent) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa89018078.RoleDaoAgent)
	nop(ie, com)

	
    com.Serivce = inst.getSerivce(ie)


    return nil
}


func (inst*pa89018078d_agent_RoleDaoAgent) getSerivce(ie application.InjectionExt)p24287f458.DaoSetService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-DaoSetService").(p24287f458.DaoSetService)
}



// type pa89018078.SessionDaoAgent in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/agent
//
// id:com-a89018078de55056-agent-SessionDaoAgent
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-SessionDAO
// scope:singleton
//
type pa89018078d_agent_SessionDaoAgent struct {
}

func (inst* pa89018078d_agent_SessionDaoAgent) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a89018078de55056-agent-SessionDaoAgent"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-SessionDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa89018078d_agent_SessionDaoAgent) new() any {
    return &pa89018078.SessionDaoAgent{}
}

func (inst* pa89018078d_agent_SessionDaoAgent) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa89018078.SessionDaoAgent)
	nop(ie, com)

	
    com.Serivce = inst.getSerivce(ie)


    return nil
}


func (inst*pa89018078d_agent_SessionDaoAgent) getSerivce(ie application.InjectionExt)p24287f458.DaoSetService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-DaoSetService").(p24287f458.DaoSetService)
}



// type pa89018078.UserDaoAgent in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/agent
//
// id:com-a89018078de55056-agent-UserDaoAgent
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-UserDAO
// scope:singleton
//
type pa89018078d_agent_UserDaoAgent struct {
}

func (inst* pa89018078d_agent_UserDaoAgent) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a89018078de55056-agent-UserDaoAgent"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-UserDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa89018078d_agent_UserDaoAgent) new() any {
    return &pa89018078.UserDaoAgent{}
}

func (inst* pa89018078d_agent_UserDaoAgent) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa89018078.UserDaoAgent)
	nop(ie, com)

	
    com.Serivce = inst.getSerivce(ie)


    return nil
}


func (inst*pa89018078d_agent_UserDaoAgent) getSerivce(ie application.InjectionExt)p24287f458.DaoSetService{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-DaoSetService").(p24287f458.DaoSetService)
}



// type p8b617a3f2.RbacDaoSetLoader in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/core
//
// id:com-8b617a3f2c4381ff-core-RbacDaoSetLoader
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-DaoSetLoader
// scope:singleton
//
type p8b617a3f2c_core_RbacDaoSetLoader struct {
}

func (inst* p8b617a3f2c_core_RbacDaoSetLoader) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8b617a3f2c4381ff-core-RbacDaoSetLoader"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-DaoSetLoader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8b617a3f2c_core_RbacDaoSetLoader) new() any {
    return &p8b617a3f2.RbacDaoSetLoader{}
}

func (inst* p8b617a3f2c_core_RbacDaoSetLoader) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8b617a3f2.RbacDaoSetLoader)
	nop(ie, com)

	
    com.DSRegList = inst.getDSRegList(ie)


    return nil
}


func (inst*p8b617a3f2c_core_RbacDaoSetLoader) getDSRegList(ie application.InjectionExt)[]p24287f458.DaoSetRegistry{
    dst := make([]p24287f458.DaoSetRegistry, 0)
    src := ie.ListComponents(".class-24287f4589fe5add27fb48a88d706565-DaoSetRegistry")
    for _, item1 := range src {
        item2 := item1.(p24287f458.DaoSetRegistry)
        dst = append(dst, item2)
    }
    return dst
}



// type p8b617a3f2.RbacDaoSetServiceImpl in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/core
//
// id:com-8b617a3f2c4381ff-core-RbacDaoSetServiceImpl
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-DaoSetService
// scope:singleton
//
type p8b617a3f2c_core_RbacDaoSetServiceImpl struct {
}

func (inst* p8b617a3f2c_core_RbacDaoSetServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-8b617a3f2c4381ff-core-RbacDaoSetServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-DaoSetService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p8b617a3f2c_core_RbacDaoSetServiceImpl) new() any {
    return &p8b617a3f2.RbacDaoSetServiceImpl{}
}

func (inst* p8b617a3f2c_core_RbacDaoSetServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p8b617a3f2.RbacDaoSetServiceImpl)
	nop(ie, com)

	
    com.Loader = inst.getLoader(ie)


    return nil
}


func (inst*p8b617a3f2c_core_RbacDaoSetServiceImpl) getLoader(ie application.InjectionExt)p24287f458.DaoSetLoader{
    return ie.GetComponent("#alias-24287f4589fe5add27fb48a88d706565-DaoSetLoader").(p24287f458.DaoSetLoader)
}



// type p9abb4d655.MemoryAuthenticationDAO in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/mem
//
// id:com-9abb4d6558cb6ecb-mem-MemoryAuthenticationDAO
// class:
// alias:alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IAuthenticationDAO
// scope:singleton
//
type p9abb4d6558_mem_MemoryAuthenticationDAO struct {
}

func (inst* p9abb4d6558_mem_MemoryAuthenticationDAO) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9abb4d6558cb6ecb-mem-MemoryAuthenticationDAO"
	r.Classes = ""
	r.Aliases = "alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IAuthenticationDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9abb4d6558_mem_MemoryAuthenticationDAO) new() any {
    return &p9abb4d655.MemoryAuthenticationDAO{}
}

func (inst* p9abb4d6558_mem_MemoryAuthenticationDAO) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9abb4d655.MemoryAuthenticationDAO)
	nop(ie, com)

	
    com.Engine = inst.getEngine(ie)


    return nil
}


func (inst*p9abb4d6558_mem_MemoryAuthenticationDAO) getEngine(ie application.InjectionExt)p9abb4d655.IMemoryEngine{
    return ie.GetComponent("#alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IMemoryEngine").(p9abb4d655.IMemoryEngine)
}



// type p9abb4d655.MemoryPermissionDao in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/mem
//
// id:com-9abb4d6558cb6ecb-mem-MemoryPermissionDao
// class:
// alias:alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IPermissionDAO
// scope:singleton
//
type p9abb4d6558_mem_MemoryPermissionDao struct {
}

func (inst* p9abb4d6558_mem_MemoryPermissionDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9abb4d6558cb6ecb-mem-MemoryPermissionDao"
	r.Classes = ""
	r.Aliases = "alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IPermissionDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9abb4d6558_mem_MemoryPermissionDao) new() any {
    return &p9abb4d655.MemoryPermissionDao{}
}

func (inst* p9abb4d6558_mem_MemoryPermissionDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9abb4d655.MemoryPermissionDao)
	nop(ie, com)

	
    com.Engine = inst.getEngine(ie)


    return nil
}


func (inst*p9abb4d6558_mem_MemoryPermissionDao) getEngine(ie application.InjectionExt)p9abb4d655.IMemoryEngine{
    return ie.GetComponent("#alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IMemoryEngine").(p9abb4d655.IMemoryEngine)
}



// type p9abb4d655.MemoryRoleDao in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/mem
//
// id:com-9abb4d6558cb6ecb-mem-MemoryRoleDao
// class:
// alias:alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IRoleDAO
// scope:singleton
//
type p9abb4d6558_mem_MemoryRoleDao struct {
}

func (inst* p9abb4d6558_mem_MemoryRoleDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9abb4d6558cb6ecb-mem-MemoryRoleDao"
	r.Classes = ""
	r.Aliases = "alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IRoleDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9abb4d6558_mem_MemoryRoleDao) new() any {
    return &p9abb4d655.MemoryRoleDao{}
}

func (inst* p9abb4d6558_mem_MemoryRoleDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9abb4d655.MemoryRoleDao)
	nop(ie, com)

	
    com.Engine = inst.getEngine(ie)


    return nil
}


func (inst*p9abb4d6558_mem_MemoryRoleDao) getEngine(ie application.InjectionExt)p9abb4d655.IMemoryEngine{
    return ie.GetComponent("#alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IMemoryEngine").(p9abb4d655.IMemoryEngine)
}



// type p9abb4d655.MemorySessionDao in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/mem
//
// id:com-9abb4d6558cb6ecb-mem-MemorySessionDao
// class:
// alias:alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-ISessionDAO
// scope:singleton
//
type p9abb4d6558_mem_MemorySessionDao struct {
}

func (inst* p9abb4d6558_mem_MemorySessionDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9abb4d6558cb6ecb-mem-MemorySessionDao"
	r.Classes = ""
	r.Aliases = "alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-ISessionDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9abb4d6558_mem_MemorySessionDao) new() any {
    return &p9abb4d655.MemorySessionDao{}
}

func (inst* p9abb4d6558_mem_MemorySessionDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9abb4d655.MemorySessionDao)
	nop(ie, com)

	
    com.Engine = inst.getEngine(ie)


    return nil
}


func (inst*p9abb4d6558_mem_MemorySessionDao) getEngine(ie application.InjectionExt)p9abb4d655.IMemoryEngine{
    return ie.GetComponent("#alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IMemoryEngine").(p9abb4d655.IMemoryEngine)
}



// type p9abb4d655.MemoryUserDao in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/mem
//
// id:com-9abb4d6558cb6ecb-mem-MemoryUserDao
// class:
// alias:alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IUserDao
// scope:singleton
//
type p9abb4d6558_mem_MemoryUserDao struct {
}

func (inst* p9abb4d6558_mem_MemoryUserDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9abb4d6558cb6ecb-mem-MemoryUserDao"
	r.Classes = ""
	r.Aliases = "alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IUserDao"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9abb4d6558_mem_MemoryUserDao) new() any {
    return &p9abb4d655.MemoryUserDao{}
}

func (inst* p9abb4d6558_mem_MemoryUserDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9abb4d655.MemoryUserDao)
	nop(ie, com)

	
    com.Engine = inst.getEngine(ie)


    return nil
}


func (inst*p9abb4d6558_mem_MemoryUserDao) getEngine(ie application.InjectionExt)p9abb4d655.IMemoryEngine{
    return ie.GetComponent("#alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IMemoryEngine").(p9abb4d655.IMemoryEngine)
}



// type p9abb4d655.MemoryEngineFacade in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/mem
//
// id:com-9abb4d6558cb6ecb-mem-MemoryEngineFacade
// class:
// alias:alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IMemoryEngine
// scope:singleton
//
type p9abb4d6558_mem_MemoryEngineFacade struct {
}

func (inst* p9abb4d6558_mem_MemoryEngineFacade) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9abb4d6558cb6ecb-mem-MemoryEngineFacade"
	r.Classes = ""
	r.Aliases = "alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IMemoryEngine"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9abb4d6558_mem_MemoryEngineFacade) new() any {
    return &p9abb4d655.MemoryEngineFacade{}
}

func (inst* p9abb4d6558_mem_MemoryEngineFacade) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9abb4d655.MemoryEngineFacade)
	nop(ie, com)

	


    return nil
}



// type p9abb4d655.MemoryDaoSetProvider in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/mem
//
// id:com-9abb4d6558cb6ecb-mem-MemoryDaoSetProvider
// class:class-24287f4589fe5add27fb48a88d706565-DaoSetRegistry
// alias:
// scope:singleton
//
type p9abb4d6558_mem_MemoryDaoSetProvider struct {
}

func (inst* p9abb4d6558_mem_MemoryDaoSetProvider) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9abb4d6558cb6ecb-mem-MemoryDaoSetProvider"
	r.Classes = "class-24287f4589fe5add27fb48a88d706565-DaoSetRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9abb4d6558_mem_MemoryDaoSetProvider) new() any {
    return &p9abb4d655.MemoryDaoSetProvider{}
}

func (inst* p9abb4d6558_mem_MemoryDaoSetProvider) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9abb4d655.MemoryDaoSetProvider)
	nop(ie, com)

	
    com.AuthenDao = inst.getAuthenDao(ie)
    com.PermDao = inst.getPermDao(ie)
    com.RoleDao = inst.getRoleDao(ie)
    com.SessionDao = inst.getSessionDao(ie)
    com.UserDao = inst.getUserDao(ie)
    com.Engine = inst.getEngine(ie)


    return nil
}


func (inst*p9abb4d6558_mem_MemoryDaoSetProvider) getAuthenDao(ie application.InjectionExt)p9abb4d655.IAuthenticationDAO{
    return ie.GetComponent("#alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IAuthenticationDAO").(p9abb4d655.IAuthenticationDAO)
}


func (inst*p9abb4d6558_mem_MemoryDaoSetProvider) getPermDao(ie application.InjectionExt)p9abb4d655.IPermissionDAO{
    return ie.GetComponent("#alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IPermissionDAO").(p9abb4d655.IPermissionDAO)
}


func (inst*p9abb4d6558_mem_MemoryDaoSetProvider) getRoleDao(ie application.InjectionExt)p9abb4d655.IRoleDAO{
    return ie.GetComponent("#alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IRoleDAO").(p9abb4d655.IRoleDAO)
}


func (inst*p9abb4d6558_mem_MemoryDaoSetProvider) getSessionDao(ie application.InjectionExt)p9abb4d655.ISessionDAO{
    return ie.GetComponent("#alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-ISessionDAO").(p9abb4d655.ISessionDAO)
}


func (inst*p9abb4d6558_mem_MemoryDaoSetProvider) getUserDao(ie application.InjectionExt)p9abb4d655.IUserDao{
    return ie.GetComponent("#alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IUserDao").(p9abb4d655.IUserDao)
}


func (inst*p9abb4d6558_mem_MemoryDaoSetProvider) getEngine(ie application.InjectionExt)p9abb4d655.IMemoryEngine{
    return ie.GetComponent("#alias-9abb4d6558cb6ecb9c8f6bb5a2dd0179-IMemoryEngine").(p9abb4d655.IMemoryEngine)
}



// type p6051122a5.UnsupportedDaoSetProvider in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/unsupported
//
// id:com-6051122a5990b446-unsupported-UnsupportedDaoSetProvider
// class:class-24287f4589fe5add27fb48a88d706565-DaoSetRegistry
// alias:
// scope:singleton
//
type p6051122a59_unsupported_UnsupportedDaoSetProvider struct {
}

func (inst* p6051122a59_unsupported_UnsupportedDaoSetProvider) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6051122a5990b446-unsupported-UnsupportedDaoSetProvider"
	r.Classes = "class-24287f4589fe5add27fb48a88d706565-DaoSetRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6051122a59_unsupported_UnsupportedDaoSetProvider) new() any {
    return &p6051122a5.UnsupportedDaoSetProvider{}
}

func (inst* p6051122a59_unsupported_UnsupportedDaoSetProvider) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6051122a5.UnsupportedDaoSetProvider)
	nop(ie, com)

	
    com.AuthenDao = inst.getAuthenDao(ie)
    com.PermDao = inst.getPermDao(ie)
    com.RoleDao = inst.getRoleDao(ie)
    com.SessionDao = inst.getSessionDao(ie)
    com.UserDao = inst.getUserDao(ie)


    return nil
}


func (inst*p6051122a59_unsupported_UnsupportedDaoSetProvider) getAuthenDao(ie application.InjectionExt)p6051122a5.IAuthenticationDAO{
    return ie.GetComponent("#alias-6051122a5990b446e1c6fd50b9ff77ac-IAuthenticationDAO").(p6051122a5.IAuthenticationDAO)
}


func (inst*p6051122a59_unsupported_UnsupportedDaoSetProvider) getPermDao(ie application.InjectionExt)p6051122a5.IPermissionDAO{
    return ie.GetComponent("#alias-6051122a5990b446e1c6fd50b9ff77ac-IPermissionDAO").(p6051122a5.IPermissionDAO)
}


func (inst*p6051122a59_unsupported_UnsupportedDaoSetProvider) getRoleDao(ie application.InjectionExt)p6051122a5.IRoleDAO{
    return ie.GetComponent("#alias-6051122a5990b446e1c6fd50b9ff77ac-IRoleDAO").(p6051122a5.IRoleDAO)
}


func (inst*p6051122a59_unsupported_UnsupportedDaoSetProvider) getSessionDao(ie application.InjectionExt)p6051122a5.ISessionDAO{
    return ie.GetComponent("#alias-6051122a5990b446e1c6fd50b9ff77ac-ISessionDAO").(p6051122a5.ISessionDAO)
}


func (inst*p6051122a59_unsupported_UnsupportedDaoSetProvider) getUserDao(ie application.InjectionExt)p6051122a5.IUserDao{
    return ie.GetComponent("#alias-6051122a5990b446e1c6fd50b9ff77ac-IUserDao").(p6051122a5.IUserDao)
}



// type p6051122a5.UnsupportedAuthenticationDAO in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/unsupported
//
// id:com-6051122a5990b446-unsupported-UnsupportedAuthenticationDAO
// class:
// alias:alias-6051122a5990b446e1c6fd50b9ff77ac-IAuthenticationDAO
// scope:singleton
//
type p6051122a59_unsupported_UnsupportedAuthenticationDAO struct {
}

func (inst* p6051122a59_unsupported_UnsupportedAuthenticationDAO) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6051122a5990b446-unsupported-UnsupportedAuthenticationDAO"
	r.Classes = ""
	r.Aliases = "alias-6051122a5990b446e1c6fd50b9ff77ac-IAuthenticationDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6051122a59_unsupported_UnsupportedAuthenticationDAO) new() any {
    return &p6051122a5.UnsupportedAuthenticationDAO{}
}

func (inst* p6051122a59_unsupported_UnsupportedAuthenticationDAO) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6051122a5.UnsupportedAuthenticationDAO)
	nop(ie, com)

	


    return nil
}



// type p6051122a5.UnsupportedPermissionDao in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/unsupported
//
// id:com-6051122a5990b446-unsupported-UnsupportedPermissionDao
// class:
// alias:alias-6051122a5990b446e1c6fd50b9ff77ac-IPermissionDAO
// scope:singleton
//
type p6051122a59_unsupported_UnsupportedPermissionDao struct {
}

func (inst* p6051122a59_unsupported_UnsupportedPermissionDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6051122a5990b446-unsupported-UnsupportedPermissionDao"
	r.Classes = ""
	r.Aliases = "alias-6051122a5990b446e1c6fd50b9ff77ac-IPermissionDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6051122a59_unsupported_UnsupportedPermissionDao) new() any {
    return &p6051122a5.UnsupportedPermissionDao{}
}

func (inst* p6051122a59_unsupported_UnsupportedPermissionDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6051122a5.UnsupportedPermissionDao)
	nop(ie, com)

	


    return nil
}



// type p6051122a5.UnsupportedRoleDao in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/unsupported
//
// id:com-6051122a5990b446-unsupported-UnsupportedRoleDao
// class:
// alias:alias-6051122a5990b446e1c6fd50b9ff77ac-IRoleDAO
// scope:singleton
//
type p6051122a59_unsupported_UnsupportedRoleDao struct {
}

func (inst* p6051122a59_unsupported_UnsupportedRoleDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6051122a5990b446-unsupported-UnsupportedRoleDao"
	r.Classes = ""
	r.Aliases = "alias-6051122a5990b446e1c6fd50b9ff77ac-IRoleDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6051122a59_unsupported_UnsupportedRoleDao) new() any {
    return &p6051122a5.UnsupportedRoleDao{}
}

func (inst* p6051122a59_unsupported_UnsupportedRoleDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6051122a5.UnsupportedRoleDao)
	nop(ie, com)

	


    return nil
}



// type p6051122a5.UnsupportedSessionDao in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/unsupported
//
// id:com-6051122a5990b446-unsupported-UnsupportedSessionDao
// class:
// alias:alias-6051122a5990b446e1c6fd50b9ff77ac-ISessionDAO
// scope:singleton
//
type p6051122a59_unsupported_UnsupportedSessionDao struct {
}

func (inst* p6051122a59_unsupported_UnsupportedSessionDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6051122a5990b446-unsupported-UnsupportedSessionDao"
	r.Classes = ""
	r.Aliases = "alias-6051122a5990b446e1c6fd50b9ff77ac-ISessionDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6051122a59_unsupported_UnsupportedSessionDao) new() any {
    return &p6051122a5.UnsupportedSessionDao{}
}

func (inst* p6051122a59_unsupported_UnsupportedSessionDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6051122a5.UnsupportedSessionDao)
	nop(ie, com)

	


    return nil
}



// type p6051122a5.UnsupportedUserDao in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/unsupported
//
// id:com-6051122a5990b446-unsupported-UnsupportedUserDao
// class:
// alias:alias-6051122a5990b446e1c6fd50b9ff77ac-IUserDao
// scope:singleton
//
type p6051122a59_unsupported_UnsupportedUserDao struct {
}

func (inst* p6051122a59_unsupported_UnsupportedUserDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6051122a5990b446-unsupported-UnsupportedUserDao"
	r.Classes = ""
	r.Aliases = "alias-6051122a5990b446e1c6fd50b9ff77ac-IUserDao"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6051122a59_unsupported_UnsupportedUserDao) new() any {
    return &p6051122a5.UnsupportedUserDao{}
}

func (inst* p6051122a59_unsupported_UnsupportedUserDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6051122a5.UnsupportedUserDao)
	nop(ie, com)

	


    return nil
}


