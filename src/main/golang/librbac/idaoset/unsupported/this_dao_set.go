package unsupported

import "github.com/starter-go/rbac"

type UnsupportedDaoSetProvider struct {

	//starter:component

	_as func(rbac.DaoSetRegistry) //starter:as(".")

	ConfigEnabled  bool //starter:inject("${rbac-dao-set.unsupported.enabled}")
	ConfigPriority int  //starter:inject("${rbac-dao-set.unsupported.priority}")

	AuthenDao  IAuthenticationDAO //starter:inject("#")
	PermDao    IPermissionDAO     //starter:inject("#")
	RoleDao    IRoleDAO           //starter:inject("#")
	SessionDao ISessionDAO        //starter:inject("#")
	UserDao    IUserDao           //starter:inject("#")

}

// Provide implements [rbac.DaoSetProvider].
func (inst *UnsupportedDaoSetProvider) Provide(dst *rbac.DaoSet) *rbac.DaoSet {

	if dst == nil {
		dst = new(rbac.DaoSet)
	}

	dst.Authentications = inst.AuthenDao
	dst.Permissions = inst.PermDao
	dst.Roles = inst.RoleDao
	dst.Sessions = inst.SessionDao
	dst.Users = inst.UserDao

	return dst
}

// Registration implements [rbac.DaoSetRegistry].
func (inst *UnsupportedDaoSetProvider) Registration() *rbac.DaoSetRegistration {
	return &rbac.DaoSetRegistration{
		Priority: inst.ConfigPriority,
		Enabled:  inst.ConfigEnabled,
		Label:    "Unsupported_Dao_Set_Provider",
		Provider: inst,
	}
}

func (inst *UnsupportedDaoSetProvider) _impl() (rbac.DaoSetRegistry, rbac.DaoSetProvider) {
	return inst, inst
}
