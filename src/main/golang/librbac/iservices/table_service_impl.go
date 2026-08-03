package iservices

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/lib/classes/tables"
)

type TableServiceImpl struct {

	//starter:component

	_as func(rbac.TableService) //starter:as("#")

	Dao rbac.TableDAO //starter:inject("#")
}

// Delete implements [tables.Service].
func (inst *TableServiceImpl) Delete(c context.Context, id tables.ID) error {
	db := inst.Dao.GetDB(nil)
	return inst.Dao.Delete(db, id)
}

// Find implements [tables.Service].
func (inst *TableServiceImpl) Find(c context.Context, id tables.ID) (*tables.DTO, error) {

	db := inst.Dao.GetDB(nil)
	it4 := new(rbac.TableDTO)

	it3, err := inst.Dao.Find(db, id)
	if err != nil {
		return nil, err
	}

	err = tables.ConvertE2D(it3, it4)
	return it4, err
}

func (inst *TableServiceImpl) Insert(c context.Context, it1 *tables.DTO) (*tables.DTO, error) {

	db := inst.Dao.GetDB(nil)
	it2 := new(rbac.TableEntity)
	it4 := new(rbac.TableDTO)

	err := tables.ConvertD2E(it1, it2)
	if err != nil {
		return nil, err
	}

	it3, err := inst.Dao.Insert(db, it2)
	if err != nil {
		return nil, err
	}

	err = tables.ConvertE2D(it3, it4)
	return it4, err
}

// List implements [tables.Service].
func (inst *TableServiceImpl) Query(c context.Context, q *tables.Query) ([]*tables.DTO, error) {
	panic("unimplemented")
}

// Update implements [tables.Service].
func (inst *TableServiceImpl) Update(c context.Context, id tables.ID, it1 *tables.DTO) (*tables.DTO, error) {

	db := inst.Dao.GetDB(nil)
	it2 := new(tables.Entity)

	err := tables.ConvertD2E(it1, it2)
	if err != nil {
		return nil, err
	}

	it3, err := inst.Dao.Update(db, id, func(old *tables.Entity) error {

		return inst.innerUpdateItem(it2, old)
	})
	if err != nil {
		return nil, err
	}

	it4 := new(tables.DTO)
	err = tables.ConvertE2D(it3, it4)
	return it4, err
}

func (inst *TableServiceImpl) innerUpdateItem(src, dst *tables.Entity) error {

	up := new(rbac.EntityUpdater)

	up.UpdateString(&src.Label, &dst.Label)
	up.UpdateString(&src.Description, &dst.Description)
	up.UpdateTableName(&src.Name, &dst.Name)

	up.UpdateString(&src.GroupName, &dst.GroupName)
	up.UpdateURI(&src.GroupURI, &dst.GroupURI)
	up.UpdateURI(&src.TableURI, &dst.TableURI)

	return nil
}

func (inst *TableServiceImpl) _impl() rbac.TableService {
	return inst
}
