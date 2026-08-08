package unsupported

import (
	"github.com/starter-go/rbac/api/classes/authentications"
	"gorm.io/gorm"
)

type UnsupportedAuthenticationDAO struct {

	//starter:component

	_as func(IAuthenticationDAO) //starter:as("#")

}

// Delete implements [IAuthenticationDAO].
func (inst *UnsupportedAuthenticationDAO) Delete(db *gorm.DB, id authentications.ID) error {
	panic("unimplemented")
}

// Find implements [IAuthenticationDAO].
func (inst *UnsupportedAuthenticationDAO) Find(db *gorm.DB, id authentications.ID) (*authentications.Entity, error) {
	panic("unimplemented")
}

// GetDB implements [IAuthenticationDAO].
func (inst *UnsupportedAuthenticationDAO) GetDB(old *gorm.DB) *gorm.DB {
	panic("unimplemented")
}

// Insert implements [IAuthenticationDAO].
func (inst *UnsupportedAuthenticationDAO) Insert(db *gorm.DB, item *authentications.Entity) (*authentications.Entity, error) {
	panic("unimplemented")
}

// Query implements [IAuthenticationDAO].
func (inst *UnsupportedAuthenticationDAO) Query(db *gorm.DB, q *authentications.Query) ([]*authentications.Entity, error) {
	panic("unimplemented")
}

// Update implements [IAuthenticationDAO].
func (inst *UnsupportedAuthenticationDAO) Update(db *gorm.DB, id authentications.ID, callback func(old *authentications.Entity) error) (*authentications.Entity, error) {
	panic("unimplemented")
}

func (inst *UnsupportedAuthenticationDAO) _impl() IAuthenticationDAO {
	return inst
}
