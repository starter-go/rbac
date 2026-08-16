package mem

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/permissions"
	"github.com/starter-go/v0/libdao"
	"github.com/starter-go/v0/libdao/api/libdaoapi"
	"gorm.io/gorm"
)

type MemoryPermissionDao struct {

	//starter:component

	_as func(rbac.PermissionDAO) //starter:as(".")

	ConfigEnabled  bool   //starter:inject("${rbac-data-group.memory.enabled}")
	ConfigPriority int    //starter:inject("${rbac-data-group.memory.priority}")
	ConfigClass    string //starter:inject("${rbac-data-group.memory.class}")

}

// Delete implements [permissions.DAO].
func (inst *MemoryPermissionDao) Delete(db *gorm.DB, id permissions.ID) error {
	panic("unimplemented")
}

// Find implements [permissions.DAO].
func (inst *MemoryPermissionDao) Find(db *gorm.DB, id permissions.ID) (*permissions.Entity, error) {
	panic("unimplemented")
}

// GetDB implements [permissions.DAO].
func (inst *MemoryPermissionDao) GetDB(old *gorm.DB) *gorm.DB {
	panic("unimplemented")
}

// Insert implements [permissions.DAO].
func (inst *MemoryPermissionDao) Insert(db *gorm.DB, item *permissions.Entity) (*permissions.Entity, error) {
	panic("unimplemented")
}

// Query implements [permissions.DAO].
func (inst *MemoryPermissionDao) Query(db *gorm.DB, q *permissions.Query) ([]*permissions.Entity, error) {
	panic("unimplemented")
}

// Update implements [permissions.DAO].
func (inst *MemoryPermissionDao) Update(db *gorm.DB, id permissions.ID, callback func(old *permissions.Entity) error) (*permissions.Entity, error) {
	panic("unimplemented")
}

// GetRegistration implements [permissions.DAO].
func (inst *MemoryPermissionDao) GetRegistration() *libdaoapi.DaoRegistration {

	r1 := &libdao.DaoRegistration{
		DAO: inst,

		Priority: inst.ConfigPriority,
		Enabled:  inst.ConfigEnabled,
		Class:    inst.ConfigClass,

		Name: "MemoryPermissionDao",
		ID:   "mem-dao-for-permissions",
	}
	return r1

}

func (inst *MemoryPermissionDao) _impl() rbac.PermissionDAO {
	return inst
}
