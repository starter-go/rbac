package mem

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/authentications"
	"github.com/starter-go/v0/libdao/api/libdaoapi"
	"gorm.io/gorm"
)

type MemoryAuthenticationDAO struct {

	//starter:component

	_as func(rbac.AuthenticationDAO) //starter:as(".")

}

// Delete implements [authentications.DAO].
func (inst *MemoryAuthenticationDAO) Delete(db *gorm.DB, id authentications.ID) error {
	panic("unimplemented")
}

// Find implements [authentications.DAO].
func (inst *MemoryAuthenticationDAO) Find(db *gorm.DB, id authentications.ID) (*authentications.Entity, error) {
	panic("unimplemented")
}

// GetDB implements [authentications.DAO].
func (inst *MemoryAuthenticationDAO) GetDB(old *gorm.DB) *gorm.DB {
	panic("unimplemented")
}

// Insert implements [authentications.DAO].
func (inst *MemoryAuthenticationDAO) Insert(db *gorm.DB, item *authentications.Entity) (*authentications.Entity, error) {
	panic("unimplemented")
}

// Query implements [authentications.DAO].
func (inst *MemoryAuthenticationDAO) Query(db *gorm.DB, q *authentications.Query) ([]*authentications.Entity, error) {
	panic("unimplemented")
}

// Update implements [authentications.DAO].
func (inst *MemoryAuthenticationDAO) Update(db *gorm.DB, id authentications.ID, callback func(old *authentications.Entity) error) (*authentications.Entity, error) {
	panic("unimplemented")
}

// GetRegistration implements [IAuthenticationDAO].
func (inst *MemoryAuthenticationDAO) GetRegistration() *libdaoapi.DaoRegistration {
	panic("unimplemented")
}

func (inst *MemoryAuthenticationDAO) _impl() rbac.AuthenticationDAO {
	return inst
}
