package mem

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/tables"
	"github.com/starter-go/v0/libdao/api/libdaoapi"
	"gorm.io/gorm"
)

type MemoryTableDao struct {

	//starter:component

	_as func(rbac.TableDAO) //starter:as(".")

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
	panic("unimplemented")
}

func (inst *MemoryTableDao) _impl() rbac.TableDAO {
	return inst
}
