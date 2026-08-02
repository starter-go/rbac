package agent

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/lib/classes/permissions"
	"gorm.io/gorm"
)

type PermissionDaoAgent struct {

	//starter:component

	_as func(rbac.PermissionDAO) //starter:as("#")

	Serivce rbac.DaoSetService //starter:inject("#")

	holder rbac.DaoSetHolder
}

func (inst *PermissionDaoAgent) target() rbac.PermissionDAO {
	ser := inst.Serivce
	tar, err := inst.holder.Get(ser)
	if err != nil {
		panic(err)
	}
	return tar.Permissions
}

// Delete implements [permissions.UserDAO].
func (inst *PermissionDaoAgent) Delete(db *gorm.DB, id permissions.ID) error {
	return inst.target().Delete(db, id)
}

// Find implements [permissions.UserDAO].
func (inst *PermissionDaoAgent) Find(db *gorm.DB, id permissions.ID) (*permissions.Entity, error) {
	return inst.target().Find(db, id)
}

// GetDB implements [permissions.UserDAO].
func (inst *PermissionDaoAgent) GetDB(old *gorm.DB) *gorm.DB {
	return inst.target().GetDB(old)
}

// Insert implements [permissions.UserDAO].
func (inst *PermissionDaoAgent) Insert(db *gorm.DB, item *permissions.Entity) (*permissions.Entity, error) {
	return inst.target().Insert(db, item)
}

// Query implements [permissions.UserDAO].
func (inst *PermissionDaoAgent) Query(db *gorm.DB, q *permissions.Query) ([]*permissions.Entity, error) {
	return inst.target().Query(db, q)
}

// Update implements [permissions.UserDAO].
func (inst *PermissionDaoAgent) Update(db *gorm.DB, id permissions.ID, callback func(old *permissions.Entity) error) (*permissions.Entity, error) {
	return inst.target().Update(db, id, callback)
}

func (inst *PermissionDaoAgent) _impl() rbac.PermissionDAO {
	return inst
}
