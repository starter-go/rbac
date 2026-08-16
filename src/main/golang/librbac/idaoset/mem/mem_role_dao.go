package mem

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/roles"
	"github.com/starter-go/v0/libdao"
	"github.com/starter-go/v0/libdao/api/libdaoapi"
	"gorm.io/gorm"
)

type MemoryRoleDao struct {

	//starter:component

	_as func(rbac.RoleDAO) //starter:as(".")

	ConfigEnabled  bool   //starter:inject("${rbac-data-group.memory.enabled}")
	ConfigPriority int    //starter:inject("${rbac-data-group.memory.priority}")
	ConfigClass    string //starter:inject("${rbac-data-group.memory.class}")

}

// Delete implements [roles.DAO].
func (inst *MemoryRoleDao) Delete(db *gorm.DB, id roles.ID) error {
	panic("unimplemented")
}

// Find implements [roles.DAO].
func (inst *MemoryRoleDao) Find(db *gorm.DB, id roles.ID) (*roles.Entity, error) {
	panic("unimplemented")
}

// GetDB implements [roles.DAO].
func (inst *MemoryRoleDao) GetDB(old *gorm.DB) *gorm.DB {
	panic("unimplemented")
}

// Insert implements [roles.DAO].
func (inst *MemoryRoleDao) Insert(db *gorm.DB, item *roles.Entity) (*roles.Entity, error) {
	panic("unimplemented")
}

// Query implements [roles.DAO].
func (inst *MemoryRoleDao) Query(db *gorm.DB, q *roles.Query) ([]*roles.Entity, error) {
	panic("unimplemented")
}

// Update implements [roles.DAO].
func (inst *MemoryRoleDao) Update(db *gorm.DB, id roles.ID, callback func(old *roles.Entity) error) (*roles.Entity, error) {
	panic("unimplemented")
}

// GetRegistration implements [roles.DAO].
func (inst *MemoryRoleDao) GetRegistration() *libdaoapi.DaoRegistration {

	r1 := &libdao.DaoRegistration{
		DAO: inst,

		Priority: inst.ConfigPriority,
		Enabled:  inst.ConfigEnabled,
		Class:    inst.ConfigClass,

		Name: "MemoryRoleDao",
		ID:   "mem-dao-for-roles",
	}
	return r1

}

func (inst *MemoryRoleDao) _impl() rbac.RoleDAO {
	return inst
}
