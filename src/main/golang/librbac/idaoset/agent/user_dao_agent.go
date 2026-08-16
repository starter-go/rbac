package agent

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/users"
	"github.com/starter-go/v0/libdao"
	"github.com/starter-go/v0/libdao/api/libdaoapi"
	"gorm.io/gorm"
)

type UserDaoAgent struct {

	//starter:component

	_as func(rbac.UserDAO) //starter:as("#")

	DaoProviderList []rbac.UserDAO                 //starter:inject(".")
	DaoSelector     string                         //starter:inject("${daoset.rbac.selector}")
	holder          libdao.DaoHolder[rbac.UserDAO] // cache for selected-dao
}

// GetRegistration implements [users.UserDAO].
func (inst *UserDaoAgent) GetRegistration() *libdaoapi.DaoRegistration {
	return new(libdao.DaoRegistration)
}

func (inst *UserDaoAgent) target() rbac.UserDAO {
	sel := inst.DaoSelector
	all := inst.DaoProviderList
	return inst.holder.Select(sel, all)
}

// FindByEmail implements [users.UserDAO].
func (inst *UserDaoAgent) FindByEmail(db *gorm.DB, addr users.EmailAddress) (*users.Entity, error) {

	return inst.target().FindByEmail(db, addr)
}

// FindByName implements [users.UserDAO].
func (inst *UserDaoAgent) FindByName(db *gorm.DB, name users.UserName) (*users.Entity, error) {

	return inst.target().FindByName(db, name)
}

// FindByPhone implements [users.UserDAO].
func (inst *UserDaoAgent) FindByPhone(db *gorm.DB, num users.PhoneNumber) (*users.Entity, error) {

	return inst.target().FindByPhone(db, num)
}

// Delete implements [users.UserDAO].
func (inst *UserDaoAgent) Delete(db *gorm.DB, id users.ID) error {
	return inst.target().Delete(db, id)
}

// Find implements [users.UserDAO].
func (inst *UserDaoAgent) Find(db *gorm.DB, id users.ID) (*users.Entity, error) {
	return inst.target().Find(db, id)
}

// GetDB implements [users.UserDAO].
func (inst *UserDaoAgent) GetDB(old *gorm.DB) *gorm.DB {
	return inst.target().GetDB(old)
}

// Insert implements [users.UserDAO].
func (inst *UserDaoAgent) Insert(db *gorm.DB, item *users.Entity) (*users.Entity, error) {
	return inst.target().Insert(db, item)
}

// Query implements [users.UserDAO].
func (inst *UserDaoAgent) Query(db *gorm.DB, q *users.Query) ([]*users.Entity, error) {
	return inst.target().Query(db, q)
}

// Update implements [users.UserDAO].
func (inst *UserDaoAgent) Update(db *gorm.DB, id users.ID, callback func(old *users.Entity) error) (*users.Entity, error) {
	return inst.target().Update(db, id, callback)
}

func (inst *UserDaoAgent) _impl() rbac.UserDAO {
	return inst
}
