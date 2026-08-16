package mem

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/users"
	"github.com/starter-go/v0/libdao"
	"github.com/starter-go/v0/libdao/api/libdaoapi"
	"gorm.io/gorm"
)

type MemoryUserDao struct {

	//starter:component

	_as func(rbac.UserDAO) //starter:as(".")

	ConfigEnabled  bool   //starter:inject("${rbac-data-group.memory.enabled}")
	ConfigPriority int    //starter:inject("${rbac-data-group.memory.priority}")
	ConfigClass    string //starter:inject("${rbac-data-group.memory.class}")
}

// Delete implements [users.UserDAO].
func (inst *MemoryUserDao) Delete(db *gorm.DB, id users.ID) error {
	panic("unimplemented")
}

// Find implements [users.UserDAO].
func (inst *MemoryUserDao) Find(db *gorm.DB, id users.ID) (*users.Entity, error) {
	panic("unimplemented")
}

// FindByEmail implements [users.UserDAO].
func (inst *MemoryUserDao) FindByEmail(db *gorm.DB, addr users.EmailAddress) (*users.Entity, error) {
	panic("unimplemented")
}

// FindByName implements [users.UserDAO].
func (inst *MemoryUserDao) FindByName(db *gorm.DB, name users.UserName) (*users.Entity, error) {
	panic("unimplemented")
}

// FindByPhone implements [users.UserDAO].
func (inst *MemoryUserDao) FindByPhone(db *gorm.DB, num users.PhoneNumber) (*users.Entity, error) {
	panic("unimplemented")
}

// GetDB implements [users.UserDAO].
func (inst *MemoryUserDao) GetDB(old *gorm.DB) *gorm.DB {
	panic("unimplemented")
}

// Insert implements [users.UserDAO].
func (inst *MemoryUserDao) Insert(db *gorm.DB, item *users.Entity) (*users.Entity, error) {
	panic("unimplemented")
}

// Query implements [users.UserDAO].
func (inst *MemoryUserDao) Query(db *gorm.DB, q *users.Query) ([]*users.Entity, error) {
	panic("unimplemented")
}

// Update implements [users.UserDAO].
func (inst *MemoryUserDao) Update(db *gorm.DB, id users.ID, callback func(old *users.Entity) error) (*users.Entity, error) {
	panic("unimplemented")
}

// GetRegistration implements [users.UserDAO].
func (inst *MemoryUserDao) GetRegistration() *libdaoapi.DaoRegistration {
	r1 := &libdao.DaoRegistration{
		DAO: inst,

		Priority: inst.ConfigPriority,
		Enabled:  inst.ConfigEnabled,
		Class:    inst.ConfigClass,

		Name: "MemoryUserDao",
		ID:   "mem-dao-for-users",
	}
	return r1
}

func (inst *MemoryUserDao) _impl() rbac.UserDAO {
	return inst
}
