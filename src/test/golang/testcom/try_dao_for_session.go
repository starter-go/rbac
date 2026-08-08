package testcom

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/sessions"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

type TryDaoForSession struct {

	//starter:component

	_as func(units.Unit) //starter:as(".")

	Dao rbac.SessionDAO //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *TryDaoForSession) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:    "try-session-dao",
		Enabled: true,
		Do:      inst.runTrySessionDao,
	}

	list = append(list, u1)
	return list
}

func (inst *TryDaoForSession) runTrySessionDao(cc context.Context) error {

	dao := inst.Dao
	it1 := new(rbac.SessionEntity)
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

	it4, err := dao.Update(db, id, func(old *sessions.Entity) error {

		old.Authenticated = false

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

func (inst *TryDaoForSession) _impl() units.Unit {
	return inst
}
