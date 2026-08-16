package agent

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/sessions"
	"github.com/starter-go/v0/libdao"
	"github.com/starter-go/v0/libdao/api/libdaoapi"
	"gorm.io/gorm"
)

type SessionDaoAgent struct {

	//starter:component

	_as func(rbac.SessionDAO) //starter:as("#")

	DaoProviderList []rbac.SessionDAO                 //starter:inject(".")
	DaoSelector     string                            //starter:inject("${daoset.rbac.selector}")
	holder          libdao.DaoHolder[rbac.SessionDAO] // cache for selected-dao

}

// GetRegistration implements [sessions.DAO].
func (inst *SessionDaoAgent) GetRegistration() *libdaoapi.DaoRegistration {
	return new(libdao.DaoRegistration)
}

func (inst *SessionDaoAgent) target() rbac.SessionDAO {

	sel := inst.DaoSelector
	all := inst.DaoProviderList
	return inst.holder.Select(sel, all)

}

// Delete implements [sessions.UserDAO].
func (inst *SessionDaoAgent) Delete(db *gorm.DB, id sessions.ID) error {
	return inst.target().Delete(db, id)
}

// Find implements [sessions.UserDAO].
func (inst *SessionDaoAgent) Find(db *gorm.DB, id sessions.ID) (*sessions.Entity, error) {
	return inst.target().Find(db, id)
}

// GetDB implements [sessions.UserDAO].
func (inst *SessionDaoAgent) GetDB(old *gorm.DB) *gorm.DB {
	return inst.target().GetDB(old)
}

// Insert implements [sessions.UserDAO].
func (inst *SessionDaoAgent) Insert(db *gorm.DB, item *sessions.Entity) (*sessions.Entity, error) {
	return inst.target().Insert(db, item)
}

// Query implements [sessions.UserDAO].
func (inst *SessionDaoAgent) Query(db *gorm.DB, q *sessions.Query) ([]*sessions.Entity, error) {
	return inst.target().Query(db, q)
}

// Update implements [sessions.UserDAO].
func (inst *SessionDaoAgent) Update(db *gorm.DB, id sessions.ID, callback func(old *sessions.Entity) error) (*sessions.Entity, error) {
	return inst.target().Update(db, id, callback)
}

func (inst *SessionDaoAgent) _impl() rbac.SessionDAO {
	return inst
}
