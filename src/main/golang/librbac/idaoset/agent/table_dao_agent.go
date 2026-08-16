package agent

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/v0/libdao"
	"github.com/starter-go/v0/libdao/api/libdaoapi"

	"github.com/starter-go/rbac/api/classes/tables"

	"gorm.io/gorm"
)

type TableDaoAgent struct {

	//starter:component

	_as func(rbac.TableDAO) //starter:as("#")

	DaoProviderList []rbac.TableDAO                 //starter:inject(".")
	DaoSelector     string                          //starter:inject("${daoset.rbac.selector}")
	holder          libdao.DaoHolder[rbac.TableDAO] // cache for selected-dao

}

// GetRegistration implements [tables.DAO].
func (inst *TableDaoAgent) GetRegistration() *libdaoapi.DaoRegistration {
	panic("unimplemented")
}

func (inst *TableDaoAgent) target() rbac.TableDAO {

	sel := inst.DaoSelector
	all := inst.DaoProviderList
	return inst.holder.Select(sel, all)

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
