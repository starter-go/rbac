package unsupported

import (
	"github.com/starter-go/rbac/api/classes/sessions"
	"gorm.io/gorm"
)

type UnsupportedSessionDao struct {

	//starter:component

	_as func(ISessionDAO) //starter:as("#")

}

// Delete implements [ISessionDAO].
func (inst *UnsupportedSessionDao) Delete(db *gorm.DB, id sessions.ID) error {
	panic("unimplemented")
}

// Find implements [ISessionDAO].
func (inst *UnsupportedSessionDao) Find(db *gorm.DB, id sessions.ID) (*sessions.Entity, error) {
	panic("unimplemented")
}

// GetDB implements [ISessionDAO].
func (inst *UnsupportedSessionDao) GetDB(old *gorm.DB) *gorm.DB {
	panic("unimplemented")
}

// Insert implements [ISessionDAO].
func (inst *UnsupportedSessionDao) Insert(db *gorm.DB, item *sessions.Entity) (*sessions.Entity, error) {
	panic("unimplemented")
}

// Query implements [ISessionDAO].
func (inst *UnsupportedSessionDao) Query(db *gorm.DB, q *sessions.Query) ([]*sessions.Entity, error) {
	panic("unimplemented")
}

// Update implements [ISessionDAO].
func (inst *UnsupportedSessionDao) Update(db *gorm.DB, id sessions.ID, callback func(old *sessions.Entity) error) (*sessions.Entity, error) {
	panic("unimplemented")
}

func (inst *UnsupportedSessionDao) _impl() ISessionDAO {
	return inst
}
