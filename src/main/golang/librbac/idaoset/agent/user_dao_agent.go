package agent

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/lib/classes/users"
	"gorm.io/gorm"
)

type UserDaoAgent struct {

	//starter:component

	_as func(rbac.UserDAO) //starter:as("#")

	Serivce rbac.DaoSetService //starter:inject("#")

	holder rbac.DaoSetHolder
}

func (inst *UserDaoAgent) target() rbac.UserDAO {
	ser := inst.Serivce
	tar, err := inst.holder.Get(ser)
	if err != nil {
		panic(err)
	}
	return tar.Users
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
