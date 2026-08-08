package agent

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/sessions"
	"gorm.io/gorm"
)

type SessionDaoAgent struct {

	//starter:component

	_as func(rbac.SessionDAO) //starter:as("#")

	Serivce rbac.DaoSetService //starter:inject("#")

	holder rbac.DaoSetHolder
}

func (inst *SessionDaoAgent) target() rbac.SessionDAO {
	ser := inst.Serivce
	tar, err := inst.holder.Get(ser)
	if err != nil {
		panic(err)
	}
	return tar.Sessions
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
