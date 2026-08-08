package unsupported

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/permissions"
	"gorm.io/gorm"
)

type UnsupportedPermissionDao struct {

	//starter:component

	_as func(IPermissionDAO) //starter:as("#")

}

// Delete implements [permissions.DAO].
func (inst *UnsupportedPermissionDao) Delete(db *gorm.DB, id permissions.ID) error {
	panic("unimplemented")
}

// Find implements [permissions.DAO].
func (inst *UnsupportedPermissionDao) Find(db *gorm.DB, id permissions.ID) (*permissions.Entity, error) {
	panic("unimplemented")
}

// GetDB implements [permissions.DAO].
func (inst *UnsupportedPermissionDao) GetDB(old *gorm.DB) *gorm.DB {
	panic("unimplemented")
}

// Insert implements [permissions.DAO].
func (inst *UnsupportedPermissionDao) Insert(db *gorm.DB, item *permissions.Entity) (*permissions.Entity, error) {
	panic("unimplemented")
}

// Query implements [permissions.DAO].
func (inst *UnsupportedPermissionDao) Query(db *gorm.DB, q *permissions.Query) ([]*permissions.Entity, error) {
	panic("unimplemented")
}

// Update implements [permissions.DAO].
func (inst *UnsupportedPermissionDao) Update(db *gorm.DB, id permissions.ID, callback func(old *permissions.Entity) error) (*permissions.Entity, error) {
	panic("unimplemented")
}

func (inst *UnsupportedPermissionDao) _impl() rbac.PermissionDAO {
	return inst
}
