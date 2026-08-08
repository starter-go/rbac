package permissions

import "gorm.io/gorm"

type DAO interface {

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
