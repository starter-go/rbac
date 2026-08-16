package main4rbac
import (
    p24287f458 "github.com/starter-go/rbac"
    pa89018078 "github.com/starter-go/rbac/src/main/golang/librbac/idaoset/agent"
    p9abb4d655 "github.com/starter-go/rbac/src/main/golang/librbac/idaoset/mem"
     "github.com/starter-go/application"
)

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

	
    com.DaoProviderList = inst.getDaoProviderList(ie)
    com.DaoSelector = inst.getDaoSelector(ie)


    return nil
}


func (inst*pa89018078d_agent_AuthentDaoAgent) getDaoProviderList(ie application.InjectionExt)[]p24287f458.AuthenticationDAO{
    dst := make([]p24287f458.AuthenticationDAO, 0)
    src := ie.ListComponents(".class-24287f4589fe5add27fb48a88d706565-AuthenticationDAO")
    for _, item1 := range src {
        item2 := item1.(p24287f458.AuthenticationDAO)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*pa89018078d_agent_AuthentDaoAgent) getDaoSelector(ie application.InjectionExt)string{
    return ie.GetString("${daoset.rbac.selector}")
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

	
    com.DaoProviderList = inst.getDaoProviderList(ie)
    com.DaoSelector = inst.getDaoSelector(ie)


    return nil
}


func (inst*pa89018078d_agent_PermissionDaoAgent) getDaoProviderList(ie application.InjectionExt)[]p24287f458.PermissionDAO{
    dst := make([]p24287f458.PermissionDAO, 0)
    src := ie.ListComponents(".class-24287f4589fe5add27fb48a88d706565-PermissionDAO")
    for _, item1 := range src {
        item2 := item1.(p24287f458.PermissionDAO)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*pa89018078d_agent_PermissionDaoAgent) getDaoSelector(ie application.InjectionExt)string{
    return ie.GetString("${daoset.rbac.selector}")
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

	
    com.DaoProviderList = inst.getDaoProviderList(ie)
    com.DaoSelector = inst.getDaoSelector(ie)


    return nil
}


func (inst*pa89018078d_agent_RoleDaoAgent) getDaoProviderList(ie application.InjectionExt)[]p24287f458.RoleDAO{
    dst := make([]p24287f458.RoleDAO, 0)
    src := ie.ListComponents(".class-24287f4589fe5add27fb48a88d706565-RoleDAO")
    for _, item1 := range src {
        item2 := item1.(p24287f458.RoleDAO)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*pa89018078d_agent_RoleDaoAgent) getDaoSelector(ie application.InjectionExt)string{
    return ie.GetString("${daoset.rbac.selector}")
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

	
    com.DaoProviderList = inst.getDaoProviderList(ie)
    com.DaoSelector = inst.getDaoSelector(ie)


    return nil
}


func (inst*pa89018078d_agent_SessionDaoAgent) getDaoProviderList(ie application.InjectionExt)[]p24287f458.SessionDAO{
    dst := make([]p24287f458.SessionDAO, 0)
    src := ie.ListComponents(".class-24287f4589fe5add27fb48a88d706565-SessionDAO")
    for _, item1 := range src {
        item2 := item1.(p24287f458.SessionDAO)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*pa89018078d_agent_SessionDaoAgent) getDaoSelector(ie application.InjectionExt)string{
    return ie.GetString("${daoset.rbac.selector}")
}



// type pa89018078.TableDaoAgent in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/agent
//
// id:com-a89018078de55056-agent-TableDaoAgent
// class:
// alias:alias-24287f4589fe5add27fb48a88d706565-TableDAO
// scope:singleton
//
type pa89018078d_agent_TableDaoAgent struct {
}

func (inst* pa89018078d_agent_TableDaoAgent) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a89018078de55056-agent-TableDaoAgent"
	r.Classes = ""
	r.Aliases = "alias-24287f4589fe5add27fb48a88d706565-TableDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa89018078d_agent_TableDaoAgent) new() any {
    return &pa89018078.TableDaoAgent{}
}

func (inst* pa89018078d_agent_TableDaoAgent) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa89018078.TableDaoAgent)
	nop(ie, com)

	
    com.DaoProviderList = inst.getDaoProviderList(ie)
    com.DaoSelector = inst.getDaoSelector(ie)


    return nil
}


func (inst*pa89018078d_agent_TableDaoAgent) getDaoProviderList(ie application.InjectionExt)[]p24287f458.TableDAO{
    dst := make([]p24287f458.TableDAO, 0)
    src := ie.ListComponents(".class-24287f4589fe5add27fb48a88d706565-TableDAO")
    for _, item1 := range src {
        item2 := item1.(p24287f458.TableDAO)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*pa89018078d_agent_TableDaoAgent) getDaoSelector(ie application.InjectionExt)string{
    return ie.GetString("${daoset.rbac.selector}")
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

	
    com.DaoProviderList = inst.getDaoProviderList(ie)
    com.DaoSelector = inst.getDaoSelector(ie)


    return nil
}


func (inst*pa89018078d_agent_UserDaoAgent) getDaoProviderList(ie application.InjectionExt)[]p24287f458.UserDAO{
    dst := make([]p24287f458.UserDAO, 0)
    src := ie.ListComponents(".class-24287f4589fe5add27fb48a88d706565-UserDAO")
    for _, item1 := range src {
        item2 := item1.(p24287f458.UserDAO)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*pa89018078d_agent_UserDaoAgent) getDaoSelector(ie application.InjectionExt)string{
    return ie.GetString("${daoset.rbac.selector}")
}



// type p9abb4d655.MemoryAuthenticationDAO in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/mem
//
// id:com-9abb4d6558cb6ecb-mem-MemoryAuthenticationDAO
// class:class-24287f4589fe5add27fb48a88d706565-AuthenticationDAO
// alias:
// scope:singleton
//
type p9abb4d6558_mem_MemoryAuthenticationDAO struct {
}

func (inst* p9abb4d6558_mem_MemoryAuthenticationDAO) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9abb4d6558cb6ecb-mem-MemoryAuthenticationDAO"
	r.Classes = "class-24287f4589fe5add27fb48a88d706565-AuthenticationDAO"
	r.Aliases = ""
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

	
    com.ConfigEnabled = inst.getConfigEnabled(ie)
    com.ConfigPriority = inst.getConfigPriority(ie)
    com.ConfigClass = inst.getConfigClass(ie)


    return nil
}


func (inst*p9abb4d6558_mem_MemoryAuthenticationDAO) getConfigEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${rbac-data-group.memory.enabled}")
}


func (inst*p9abb4d6558_mem_MemoryAuthenticationDAO) getConfigPriority(ie application.InjectionExt)int{
    return ie.GetInt("${rbac-data-group.memory.priority}")
}


func (inst*p9abb4d6558_mem_MemoryAuthenticationDAO) getConfigClass(ie application.InjectionExt)string{
    return ie.GetString("${rbac-data-group.memory.class}")
}



// type p9abb4d655.MemoryPermissionDao in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/mem
//
// id:com-9abb4d6558cb6ecb-mem-MemoryPermissionDao
// class:class-24287f4589fe5add27fb48a88d706565-PermissionDAO
// alias:
// scope:singleton
//
type p9abb4d6558_mem_MemoryPermissionDao struct {
}

func (inst* p9abb4d6558_mem_MemoryPermissionDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9abb4d6558cb6ecb-mem-MemoryPermissionDao"
	r.Classes = "class-24287f4589fe5add27fb48a88d706565-PermissionDAO"
	r.Aliases = ""
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

	
    com.ConfigEnabled = inst.getConfigEnabled(ie)
    com.ConfigPriority = inst.getConfigPriority(ie)
    com.ConfigClass = inst.getConfigClass(ie)


    return nil
}


func (inst*p9abb4d6558_mem_MemoryPermissionDao) getConfigEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${rbac-data-group.memory.enabled}")
}


func (inst*p9abb4d6558_mem_MemoryPermissionDao) getConfigPriority(ie application.InjectionExt)int{
    return ie.GetInt("${rbac-data-group.memory.priority}")
}


func (inst*p9abb4d6558_mem_MemoryPermissionDao) getConfigClass(ie application.InjectionExt)string{
    return ie.GetString("${rbac-data-group.memory.class}")
}



// type p9abb4d655.MemoryRoleDao in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/mem
//
// id:com-9abb4d6558cb6ecb-mem-MemoryRoleDao
// class:class-24287f4589fe5add27fb48a88d706565-RoleDAO
// alias:
// scope:singleton
//
type p9abb4d6558_mem_MemoryRoleDao struct {
}

func (inst* p9abb4d6558_mem_MemoryRoleDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9abb4d6558cb6ecb-mem-MemoryRoleDao"
	r.Classes = "class-24287f4589fe5add27fb48a88d706565-RoleDAO"
	r.Aliases = ""
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

	
    com.ConfigEnabled = inst.getConfigEnabled(ie)
    com.ConfigPriority = inst.getConfigPriority(ie)
    com.ConfigClass = inst.getConfigClass(ie)


    return nil
}


func (inst*p9abb4d6558_mem_MemoryRoleDao) getConfigEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${rbac-data-group.memory.enabled}")
}


func (inst*p9abb4d6558_mem_MemoryRoleDao) getConfigPriority(ie application.InjectionExt)int{
    return ie.GetInt("${rbac-data-group.memory.priority}")
}


func (inst*p9abb4d6558_mem_MemoryRoleDao) getConfigClass(ie application.InjectionExt)string{
    return ie.GetString("${rbac-data-group.memory.class}")
}



// type p9abb4d655.MemorySessionDao in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/mem
//
// id:com-9abb4d6558cb6ecb-mem-MemorySessionDao
// class:class-24287f4589fe5add27fb48a88d706565-SessionDAO
// alias:
// scope:singleton
//
type p9abb4d6558_mem_MemorySessionDao struct {
}

func (inst* p9abb4d6558_mem_MemorySessionDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9abb4d6558cb6ecb-mem-MemorySessionDao"
	r.Classes = "class-24287f4589fe5add27fb48a88d706565-SessionDAO"
	r.Aliases = ""
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

	
    com.ConfigEnabled = inst.getConfigEnabled(ie)
    com.ConfigPriority = inst.getConfigPriority(ie)
    com.ConfigClass = inst.getConfigClass(ie)


    return nil
}


func (inst*p9abb4d6558_mem_MemorySessionDao) getConfigEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${rbac-data-group.memory.enabled}")
}


func (inst*p9abb4d6558_mem_MemorySessionDao) getConfigPriority(ie application.InjectionExt)int{
    return ie.GetInt("${rbac-data-group.memory.priority}")
}


func (inst*p9abb4d6558_mem_MemorySessionDao) getConfigClass(ie application.InjectionExt)string{
    return ie.GetString("${rbac-data-group.memory.class}")
}



// type p9abb4d655.MemoryTableDao in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/mem
//
// id:com-9abb4d6558cb6ecb-mem-MemoryTableDao
// class:class-24287f4589fe5add27fb48a88d706565-TableDAO
// alias:
// scope:singleton
//
type p9abb4d6558_mem_MemoryTableDao struct {
}

func (inst* p9abb4d6558_mem_MemoryTableDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9abb4d6558cb6ecb-mem-MemoryTableDao"
	r.Classes = "class-24287f4589fe5add27fb48a88d706565-TableDAO"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9abb4d6558_mem_MemoryTableDao) new() any {
    return &p9abb4d655.MemoryTableDao{}
}

func (inst* p9abb4d6558_mem_MemoryTableDao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9abb4d655.MemoryTableDao)
	nop(ie, com)

	
    com.ConfigEnabled = inst.getConfigEnabled(ie)
    com.ConfigPriority = inst.getConfigPriority(ie)
    com.ConfigClass = inst.getConfigClass(ie)


    return nil
}


func (inst*p9abb4d6558_mem_MemoryTableDao) getConfigEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${rbac-data-group.memory.enabled}")
}


func (inst*p9abb4d6558_mem_MemoryTableDao) getConfigPriority(ie application.InjectionExt)int{
    return ie.GetInt("${rbac-data-group.memory.priority}")
}


func (inst*p9abb4d6558_mem_MemoryTableDao) getConfigClass(ie application.InjectionExt)string{
    return ie.GetString("${rbac-data-group.memory.class}")
}



// type p9abb4d655.MemoryUserDao in package:github.com/starter-go/rbac/src/main/golang/librbac/idaoset/mem
//
// id:com-9abb4d6558cb6ecb-mem-MemoryUserDao
// class:class-24287f4589fe5add27fb48a88d706565-UserDAO
// alias:
// scope:singleton
//
type p9abb4d6558_mem_MemoryUserDao struct {
}

func (inst* p9abb4d6558_mem_MemoryUserDao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9abb4d6558cb6ecb-mem-MemoryUserDao"
	r.Classes = "class-24287f4589fe5add27fb48a88d706565-UserDAO"
	r.Aliases = ""
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

	
    com.ConfigEnabled = inst.getConfigEnabled(ie)
    com.ConfigPriority = inst.getConfigPriority(ie)
    com.ConfigClass = inst.getConfigClass(ie)


    return nil
}


func (inst*p9abb4d6558_mem_MemoryUserDao) getConfigEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${rbac-data-group.memory.enabled}")
}


func (inst*p9abb4d6558_mem_MemoryUserDao) getConfigPriority(ie application.InjectionExt)int{
    return ie.GetInt("${rbac-data-group.memory.priority}")
}


func (inst*p9abb4d6558_mem_MemoryUserDao) getConfigClass(ie application.InjectionExt)string{
    return ie.GetString("${rbac-data-group.memory.class}")
}


