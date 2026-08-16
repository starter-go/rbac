package agent

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/roles"
	"github.com/starter-go/v0/libdao"
	"github.com/starter-go/v0/libdao/api/libdaoapi"
	"gorm.io/gorm"
)

type RoleDaoAgent struct {

	//starter:component

	_as func(rbac.RoleDAO) //starter:as("#")

	DaoProviderList []rbac.RoleDAO                 //starter:inject(".")
	DaoSelector     string                         //starter:inject("${daoset.rbac.selector}")
	holder          libdao.DaoHolder[rbac.RoleDAO] // cache for selected-dao

}

// GetRegistration implements [roles.DAO].
func (inst *RoleDaoAgent) GetRegistration() *libdaoapi.DaoRegistration {
	panic("unimplemented")
}

func (inst *RoleDaoAgent) target() rbac.RoleDAO {

	sel := inst.DaoSelector
	all := inst.DaoProviderList
	return inst.holder.Select(sel, all)

}

// Delete implements [roles.UserDAO].
func (inst *RoleDaoAgent) Delete(db *gorm.DB, id roles.ID) error {
	return inst.target().Delete(db, id)
}

// Find implements [roles.UserDAO].
func (inst *RoleDaoAgent) Find(db *gorm.DB, id roles.ID) (*roles.Entity, error) {
	return inst.target().Find(db, id)
}

// GetDB implements [roles.UserDAO].
func (inst *RoleDaoAgent) GetDB(old *gorm.DB) *gorm.DB {
	return inst.target().GetDB(old)
}

// Insert implements [roles.UserDAO].
func (inst *RoleDaoAgent) Insert(db *gorm.DB, item *roles.Entity) (*roles.Entity, error) {
	return inst.target().Insert(db, item)
}

// Query implements [roles.UserDAO].
func (inst *RoleDaoAgent) Query(db *gorm.DB, q *roles.Query) ([]*roles.Entity, error) {
	return inst.target().Query(db, q)
}

// Update implements [roles.UserDAO].
func (inst *RoleDaoAgent) Update(db *gorm.DB, id roles.ID, callback func(old *roles.Entity) error) (*roles.Entity, error) {
	return inst.target().Update(db, id, callback)
}

func (inst *RoleDaoAgent) _impl() rbac.RoleDAO {
	return inst
}
