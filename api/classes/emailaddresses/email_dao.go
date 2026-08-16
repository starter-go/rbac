package emailaddresses

import (
	"github.com/starter-go/v0/libdao"
	"gorm.io/gorm"
)

// EmailAddress_DAO ...
type DAO interface {

	// extends

	libdao.DAO

	// edit

	Insert(db *gorm.DB, o *Entity) (*Entity, error)

	Update(db *gorm.DB, id ID, callback func(old *Entity) error) (*Entity, error)

	Delete(db *gorm.DB, id ID) error

	// query

	Find(db *gorm.DB, id ID) (*Entity, error)

	List(db *gorm.DB, q *Query) ([]*Entity, error)
}
