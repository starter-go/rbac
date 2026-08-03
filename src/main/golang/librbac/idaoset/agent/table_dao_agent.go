package agent

import (
	"github.com/starter-go/rbac"

	"github.com/starter-go/rbac/lib/classes/tables"

	"gorm.io/gorm"
)

type TableDaoAgent struct {

	//starter:component

	_as func(rbac.TableDAO) //starter:as("#")

	Serivce rbac.DaoSetService //starter:inject("#")

	holder rbac.DaoSetHolder
}

func (inst *TableDaoAgent) target() rbac.TableDAO {
	ser := inst.Serivce
	tar, err := inst.holder.Get(ser)
	if err != nil {
		panic(err)
	}
	return tar.Tables
}

// Delete implements [tables.UserDAO].
func (inst *TableDaoAgent) Delete(db *gorm.DB, id tables.ID) error {
	return inst.target().Delete(db, id)
}

// Find implements [tables.UserDAO].
func (inst *TableDaoAgent) Find(db *gorm.DB, id tables.ID) (*tables.Entity, error) {
	return inst.target().Find(db, id)
}

// GetDB implements [tables.UserDAO].
func (inst *TableDaoAgent) GetDB(old *gorm.DB) *gorm.DB {
	return inst.target().GetDB(old)
}

// Insert implements [tables.UserDAO].
func (inst *TableDaoAgent) Insert(db *gorm.DB, item *tables.Entity) (*tables.Entity, error) {
	return inst.target().Insert(db, item)
}

// Query implements [tables.UserDAO].
func (inst *TableDaoAgent) Query(db *gorm.DB, q *tables.Query) ([]*tables.Entity, error) {
	return inst.target().Query(db, q)
}

// Update implements [tables.UserDAO].
func (inst *TableDaoAgent) Update(db *gorm.DB, id tables.ID, callback func(old *tables.Entity) error) (*tables.Entity, error) {
	return inst.target().Update(db, id, callback)
}

func (inst *TableDaoAgent) _impl() rbac.TableDAO {
	return inst
}
