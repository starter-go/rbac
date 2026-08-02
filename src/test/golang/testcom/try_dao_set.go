package testcom

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/units"
)

type TryDaoSet struct {

	//starter:component

	_as func(units.Unit) //starter:as(".")

	DaoSetService rbac.DaoSetService //starter:inject("#")
	AuthService   rbac.AuthService   //starter:inject("#")

	// others

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
func (inst *TryDaoSet) ListRegistrations(list []*units.Registration) []*units.Registration {

	u2 := &units.Registration{
		Name:     "try-dao-set",
		Enabled:  true,
		Do:       inst.runTryDaoSet,
		Priority: -11,
	}

	list = append(list, u2)
	return list
}

func (inst *TryDaoSet) runTryDaoSet(cc context.Context) error {

	var holder rbac.DaoSetHolder
	ser := inst.DaoSetService
	ds, err := holder.Get(ser)
	if err != nil {
		return err
	}

	role := new(rbac.RoleEntity)

	ds.Roles.Insert(nil, role)

	return nil
}

func (inst *TryDaoSet) _impl() units.Unit {
	return inst
}
