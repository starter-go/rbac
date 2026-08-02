package testcom

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/lib/classes/users"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

type TryDaoForUser struct {

	//starter:component

	_as func(units.Unit) //starter:as(".")

	UserDao rbac.UserDAO //starter:inject("#")
}

// ListRegistrations implements units.Unit.
func (inst *TryDaoForUser) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:    "try-user-dao",
		Enabled: true,
		Do:      inst.runTryUserDao,
	}

	list = append(list, u1)
	return list
}

func (inst *TryDaoForUser) runTryUserDao(cc context.Context) error {

	dao := inst.UserDao
	it1 := new(rbac.UserEntity)
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

	it4, err := dao.Update(db, id, func(old *users.Entity) error {

		old.Avatar = "http://"
		old.NickName = "foo"

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

func (inst *TryDaoForUser) _impl() units.Unit {
	return inst
}
