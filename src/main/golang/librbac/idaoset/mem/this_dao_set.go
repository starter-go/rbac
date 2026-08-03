package mem

import "github.com/starter-go/rbac"

type MemoryDaoSetProvider struct {

	//starter:component

	_as func(rbac.DaoSetRegistry) //starter:as(".")

	ConfigEnabled  bool //starter:inject("${rbac-dao-set.ram.enabled}")
	ConfigPriority int  //starter:inject("${rbac-dao-set.ram.priority}")

	AuthenDao  IAuthenticationDAO //starter:inject("#")
	PermDao    IPermissionDAO     //starter:inject("#")
	RoleDao    IRoleDAO           //starter:inject("#")
	SessionDao ISessionDAO        //starter:inject("#")
	TableDao   ITableDAO          //starter:inject("#")
	UserDao    IUserDao           //starter:inject("#")

	Engine IMemoryEngine //starter:inject("#")

}

// Provide implements [rbac.DaoSetProvider].
func (inst *MemoryDaoSetProvider) Provide(dst *rbac.DaoSet) *rbac.DaoSet {

	if dst == nil {
		dst = new(rbac.DaoSet)
	}

	dst.Authentications = inst.AuthenDao
	dst.Permissions = inst.PermDao
	dst.Roles = inst.RoleDao
	dst.Sessions = inst.SessionDao
	dst.Tables = inst.TableDao
	dst.Users = inst.UserDao

	return dst
}

// Registration implements [rbac.DaoSetRegistry].
func (inst *MemoryDaoSetProvider) Registration() *rbac.DaoSetRegistration {

	eng := inst.Engine

	eng.InitTable(new(rbac.AuthenticationEntity))
	eng.InitTable(new(rbac.PermissionEntity))
	eng.InitTable(new(rbac.RoleEntity))
	eng.InitTable(new(rbac.SessionEntity))
	eng.InitTable(new(rbac.TableEntity))
	eng.InitTable(new(rbac.UserEntity))

	return &rbac.DaoSetRegistration{
		Priority: inst.ConfigPriority,
		Enabled:  inst.ConfigEnabled,
		Label:    "MemoryDaoSetProvider",
		Provider: inst,
	}
}

func (inst *MemoryDaoSetProvider) _impl() (rbac.DaoSetRegistry, rbac.DaoSetProvider) {
	return inst, inst
}
