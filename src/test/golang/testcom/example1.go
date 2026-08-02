package testcom

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/units"
)

type Example1 struct {

	//starter:component

	AuthReg rbac.AuthRegistry
	AuthSer rbac.AuthService

	AuthenticationService rbac.AuthenticationService
	GroupService          rbac.GroupService
	PermissionService     rbac.PermissionService
	RegionService         rbac.RegionService
	RoleService           rbac.RoleService
	SessionService        rbac.SessionService
	TableService          rbac.TableService
	UserService           rbac.UserService

	AuthenticationDAO rbac.AuthenticationDAO
	GroupDao          rbac.GroupDAO
	PermissionDao     rbac.PermissionDAO
	RegionDao         rbac.RegionDAO
	RoleDao           rbac.RoleDAO
	SessionDao        rbac.SessionDAO
	TableDao          rbac.TableDAO
	UserDao           rbac.UserDAO
}

// ListRegistrations implements units.Unit.
func (inst *Example1) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{

		Name:    "Example1",
		Enabled: true,
		Do:      inst.run,
	}

	list = append(list, u1)
	return list
}

func (inst *Example1) run(cc context.Context) error {
	return nil
}

func (inst *Example1) _impl() units.Unit {
	return inst
}
