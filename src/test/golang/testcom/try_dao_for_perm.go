package testcom

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/permissions"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

type TryDaoForPerm struct {

	//starter:component

	_as func(units.Unit) //starter:as(".")

	PermissionDao rbac.PermissionDAO //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *TryDaoForPerm) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:    "try-perm-dao",
		Enabled: true,
		Do:      inst.runTryPermDao,
	}

	list = append(list, u1)
	return list
}

func (inst *TryDaoForPerm) runTryPermDao(cc context.Context) error {

	dao := inst.PermissionDao
	it1 := new(rbac.PermissionEntity)
	db := dao.GetDB(nil)

	it2, err := dao.Insert(db, it1)
	if err != nil {
		return err
	}

	id := it2.ID

	it3, err := dao.Find(db, id)
	if err != nil {
		return err
	}

	it4, err := dao.Update(db, id, func(old *permissions.Entity) error {

		old.Method = "GET"
		old.Path = "/a/b/c"

		return nil
	})
	if err != nil {
		return err
	}

	err = dao.Delete(db, id)
	if err != nil {
		return err
	}

	_, err = dao.Find(db, id)
	if err != nil {
		vlog.Warn(" (find after delete) error : %s", err.Error())
	}

	vlog.Debug("insert user[1].entity.id = %d", it1.ID)
	vlog.Debug("insert user[2].entity.id = %d", it2.ID)
	vlog.Debug("insert user[3].entity.id = %d", it3.ID)
	vlog.Debug("insert user[4].entity.id = %d", it4.ID)

	return nil
}

func (inst *TryDaoForPerm) _impl() units.Unit {
	return inst
}
