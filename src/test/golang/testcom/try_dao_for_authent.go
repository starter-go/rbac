package testcom

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/lib/classes/authentications"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

type TryDaoForAuthent struct {

	//starter:component

	_as func(units.Unit) //starter:as(".")

	Dao rbac.AuthenticationDAO //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *TryDaoForAuthent) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:    "try-authent-dao",
		Enabled: true,
		Do:      inst.runTryAuthentDao,
	}

	list = append(list, u1)
	return list
}

func (inst *TryDaoForAuthent) runTryAuthentDao(cc context.Context) error {

	dao := inst.Dao
	it1 := new(rbac.AuthenticationEntity)
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

	it4, err := dao.Update(db, id, func(old *authentications.Entity) error {

		old.CommonName = "foo"

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

func (inst *TryDaoForAuthent) _impl() units.Unit {
	return inst
}
