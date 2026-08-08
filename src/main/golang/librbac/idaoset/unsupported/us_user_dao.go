package unsupported

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/users"
	"gorm.io/gorm"
)

type UnsupportedUserDao struct {

	//starter:component

	_as func(IUserDao) //starter:as("#")

}

// FindByEmail implements [users.UserDAO].
func (inst *UnsupportedUserDao) FindByEmail(db *gorm.DB, addr users.EmailAddress) (*users.Entity, error) {
	panic("unimplemented")
}

// FindByName implements [users.UserDAO].
func (inst *UnsupportedUserDao) FindByName(db *gorm.DB, name users.UserName) (*users.Entity, error) {
	panic("unimplemented")
}

// FindByPhone implements [users.UserDAO].
func (inst *UnsupportedUserDao) FindByPhone(db *gorm.DB, num users.PhoneNumber) (*users.Entity, error) {
	panic("unimplemented")
}

// Delete implements [users.UserDAO].
func (inst *UnsupportedUserDao) Delete(db *gorm.DB, id users.ID) error {
	panic("unimplemented")
}

// Find implements [users.UserDAO].
func (inst *UnsupportedUserDao) Find(db *gorm.DB, id users.ID) (*users.Entity, error) {
	panic("unimplemented")
}

// GetDB implements [users.UserDAO].
func (inst *UnsupportedUserDao) GetDB(old *gorm.DB) *gorm.DB {
	panic("unimplemented")
}

// Insert implements [users.UserDAO].
func (inst *UnsupportedUserDao) Insert(db *gorm.DB, item *users.Entity) (*users.Entity, error) {
	panic("unimplemented")
}

// Query implements [users.UserDAO].
func (inst *UnsupportedUserDao) Query(db *gorm.DB, q *users.Query) ([]*users.Entity, error) {
	panic("unimplemented")
}

// Update implements [users.UserDAO].
func (inst *UnsupportedUserDao) Update(db *gorm.DB, id users.ID, callback func(old *users.Entity) error) (*users.Entity, error) {
	panic("unimplemented")
}

func (inst *UnsupportedUserDao) _impl() rbac.UserDAO {
	return inst
}
