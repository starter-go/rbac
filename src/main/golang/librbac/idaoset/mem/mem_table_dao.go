package mem

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/tables"
	"github.com/starter-go/v0/libdao"
	"github.com/starter-go/v0/libdao/api/libdaoapi"
	"gorm.io/gorm"
)

type MemoryTableDao struct {

	//starter:component

	_as func(rbac.TableDAO) //starter:as(".")

	ConfigEnabled  bool   //starter:inject("${rbac-data-group.memory.enabled}")
	ConfigPriority int    //starter:inject("${rbac-data-group.memory.priority}")
	ConfigClass    string //starter:inject("${rbac-data-group.memory.class}")

}

// Delete implements [tables.DAO].
func (inst *MemoryTableDao) Delete(db *gorm.DB, id tables.TableID) error {
	panic("unimplemented")
}

// Find implements [tables.DAO].
func (inst *MemoryTableDao) Find(db *gorm.DB, id tables.TableID) (*tables.TableEntity, error) {
	panic("unimplemented")
}

// GetDB implements [tables.DAO].
func (inst *MemoryTableDao) GetDB(old *gorm.DB) *gorm.DB {
	panic("unimplemented")
}

// Insert implements [tables.DAO].
func (inst *MemoryTableDao) Insert(db *gorm.DB, item *tables.TableEntity) (*tables.TableEntity, error) {
	panic("unimplemented")
}

// Query implements [tables.DAO].
func (inst *MemoryTableDao) Query(db *gorm.DB, q *tables.TableQuery) ([]*tables.TableEntity, error) {
	panic("unimplemented")
}

// Update implements [tables.DAO].
func (inst *MemoryTableDao) Update(db *gorm.DB, id tables.TableID, callback func(old *tables.TableEntity) error) (*tables.TableEntity, error) {
	panic("unimplemented")
}

// GetRegistration implements [ITableDAO].
func (inst *MemoryTableDao) GetRegistration() *libdaoapi.DaoRegistration {

	r1 := &libdao.DaoRegistration{
		DAO: inst,

		Priority: inst.ConfigPriority,
		Enabled:  inst.ConfigEnabled,
		Class:    inst.ConfigClass,

		Name: "MemoryTableDao",
		ID:   "mem-dao-for-tables",
	}
	return r1

}

func (inst *MemoryTableDao) _impl() rbac.TableDAO {
	return inst
}
