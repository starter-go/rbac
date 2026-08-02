package unsupported

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/lib/classes/roles"
	"gorm.io/gorm"
)

type UnsupportedRoleDao struct {

	//starter:component

	_as func(IRoleDAO) //starter:as("#")

}

// Delete implements [roles.DAO].
func (inst *UnsupportedRoleDao) Delete(db *gorm.DB, id roles.ID) error {
	panic("unimplemented")
}

// Find implements [roles.DAO].
func (inst *UnsupportedRoleDao) Find(db *gorm.DB, id roles.ID) (*roles.Entity, error) {
	panic("unimplemented")
}

// GetDB implements [roles.DAO].
func (inst *UnsupportedRoleDao) GetDB(old *gorm.DB) *gorm.DB {
	panic("unimplemented")
}

// Insert implements [roles.DAO].
func (inst *UnsupportedRoleDao) Insert(db *gorm.DB, item *roles.Entity) (*roles.Entity, error) {
	panic("unimplemented")
}

// Query implements [roles.DAO].
func (inst *UnsupportedRoleDao) Query(db *gorm.DB, q *roles.Query) ([]*roles.Entity, error) {
	panic("unimplemented")
}

// Update implements [roles.DAO].
func (inst *UnsupportedRoleDao) Update(db *gorm.DB, id roles.ID, callback func(old *roles.Entity) error) (*roles.Entity, error) {
	panic("unimplemented")
}

func (inst *UnsupportedRoleDao) _impl() rbac.RoleDAO {
	return inst
}
