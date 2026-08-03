package testcom

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/lib/classes/tables"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

type TryDaoForTable struct {

	//starter:component

	_as func(units.Unit) //starter:as(".")

	Dao rbac.TableDAO     //starter:inject("#")
	Ser rbac.TableService //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *TryDaoForTable) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:    "try-table-dao",
		Enabled: true,
		Do:      inst.runTryTableDao,
	}

	u2 := &units.Registration{
		Name:    "try-table-service",
		Enabled: true,
		Do:      inst.runTryTableService,
	}

	list = append(list, u1, u2)
	return list
}

func (inst *TryDaoForTable) runTryTableService(cc context.Context) error {

	return nil

}

func (inst *TryDaoForTable) runTryTableDao(cc context.Context) error {

	dao := inst.Dao
	it1 := new(rbac.TableEntity)
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

	it4, err := dao.Update(db, id, func(old *tables.Entity) error {

		old.Name = ""

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

func (inst *TryDaoForTable) _impl() units.Unit {
	return inst
}
