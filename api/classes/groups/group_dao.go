package groups

import (
	"github.com/starter-go/v0/libdao"
	"gorm.io/gorm"
)

// Group_DAO 是针对 GroupEntity 的 DAO
type DAO interface {

	// extends

	libdao.DAO

	// db

	GetDB(old *gorm.DB) *gorm.DB

	// fetch

	Find(db *gorm.DB, id ID) (*Entity, error)

	Query(db *gorm.DB, q *Query) ([]*Entity, error)

	// modify

	Insert(db *gorm.DB, item *Entity) (*Entity, error)

	Update(db *gorm.DB, id ID, callback func(old *Entity) error) (*Entity, error)

	Delete(db *gorm.DB, id ID) error
}
