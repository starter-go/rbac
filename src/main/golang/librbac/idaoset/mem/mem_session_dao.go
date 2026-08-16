package mem

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/sessions"
	"github.com/starter-go/v0/libdao/api/libdaoapi"
	"gorm.io/gorm"
)

type MemorySessionDao struct {

	//starter:component

	_as func(rbac.SessionDAO) //starter:as(".")

}

// Delete implements [sessions.DAO].
func (inst *MemorySessionDao) Delete(db *gorm.DB, id sessions.ID) error {
	panic("unimplemented")
}

// Find implements [sessions.DAO].
func (inst *MemorySessionDao) Find(db *gorm.DB, id sessions.ID) (*sessions.Entity, error) {
	panic("unimplemented")
}

// GetDB implements [sessions.DAO].
func (inst *MemorySessionDao) GetDB(old *gorm.DB) *gorm.DB {
	panic("unimplemented")
}

// Insert implements [sessions.DAO].
func (inst *MemorySessionDao) Insert(db *gorm.DB, item *sessions.Entity) (*sessions.Entity, error) {
	panic("unimplemented")
}

// Query implements [sessions.DAO].
func (inst *MemorySessionDao) Query(db *gorm.DB, q *sessions.Query) ([]*sessions.Entity, error) {
	panic("unimplemented")
}

// Update implements [sessions.DAO].
func (inst *MemorySessionDao) Update(db *gorm.DB, id sessions.ID, callback func(old *sessions.Entity) error) (*sessions.Entity, error) {
	panic("unimplemented")
}

// GetRegistration implements [ISessionDAO].
func (inst *MemorySessionDao) GetRegistration() *libdaoapi.DaoRegistration {
	panic("unimplemented")
}

func (inst *MemorySessionDao) _impl() rbac.SessionDAO {
	return inst
}
