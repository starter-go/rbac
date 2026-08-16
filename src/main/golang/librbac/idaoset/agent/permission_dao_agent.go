package agent

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/permissions"
	"github.com/starter-go/v0/libdao"
	"github.com/starter-go/v0/libdao/api/libdaoapi"
	"gorm.io/gorm"
)

type PermissionDaoAgent struct {

	//starter:component

	_as func(rbac.PermissionDAO) //starter:as("#")

	DaoProviderList []rbac.PermissionDAO                 //starter:inject(".")
	DaoSelector     string                               //starter:inject("${daoset.rbac.selector}")
	holder          libdao.DaoHolder[rbac.PermissionDAO] // cache for selected-dao

}

// GetRegistration implements [permissions.DAO].
func (inst *PermissionDaoAgent) GetRegistration() *libdaoapi.DaoRegistration {
	return new(libdao.DaoRegistration)
}

func (inst *PermissionDaoAgent) target() rbac.PermissionDAO {

	sel := inst.DaoSelector
	all := inst.DaoProviderList
	return inst.holder.Select(sel, all)

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
