package agent

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/lib/classes/authentications"
	"gorm.io/gorm"
)

type AuthentDaoAgent struct {

	//starter:component

	_as func(rbac.AuthenticationDAO) //starter:as("#")

	Serivce rbac.DaoSetService //starter:inject("#")

	holder rbac.DaoSetHolder
}

func (inst *AuthentDaoAgent) target() rbac.AuthenticationDAO {
	ser := inst.Serivce
	tar, err := inst.holder.Get(ser)
	if err != nil {
		panic(err)
	}
	return tar.Authentications
}

// Delete implements [authentications.UserDAO].
func (inst *AuthentDaoAgent) Delete(db *gorm.DB, id authentications.ID) error {
	return inst.target().Delete(db, id)
}

// Find implements [authentications.UserDAO].
func (inst *AuthentDaoAgent) Find(db *gorm.DB, id authentications.ID) (*authentications.Entity, error) {
	return inst.target().Find(db, id)
}

// GetDB implements [authentications.UserDAO].
func (inst *AuthentDaoAgent) GetDB(old *gorm.DB) *gorm.DB {
	return inst.target().GetDB(old)
}

// Insert implements [authentications.UserDAO].
func (inst *AuthentDaoAgent) Insert(db *gorm.DB, item *authentications.Entity) (*authentications.Entity, error) {
	return inst.target().Insert(db, item)
}

// Query implements [authentications.UserDAO].
func (inst *AuthentDaoAgent) Query(db *gorm.DB, q *authentications.Query) ([]*authentications.Entity, error) {
	return inst.target().Query(db, q)
}

// Update implements [authentications.UserDAO].
func (inst *AuthentDaoAgent) Update(db *gorm.DB, id authentications.ID, callback func(old *authentications.Entity) error) (*authentications.Entity, error) {
	return inst.target().Update(db, id, callback)
}

func (inst *AuthentDaoAgent) _impl() rbac.AuthenticationDAO {
	return inst
}
