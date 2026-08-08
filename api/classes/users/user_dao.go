package users

import "gorm.io/gorm"

type UserDAO interface {

	// db

	GetDB(old *gorm.DB) *gorm.DB

	// fetch

	Find(db *gorm.DB, id ID) (*Entity, error)

	Query(db *gorm.DB, q *Query) ([]*Entity, error)

	FindByEmail(db *gorm.DB, addr EmailAddress) (*Entity, error)

	FindByPhone(db *gorm.DB, num PhoneNumber) (*Entity, error)

	FindByName(db *gorm.DB, name UserName) (*Entity, error)

	// modify

	Insert(db *gorm.DB, item *Entity) (*Entity, error)

	Update(db *gorm.DB, id ID, callback func(old *Entity) error) (*Entity, error)

	Delete(db *gorm.DB, id ID) error
}
