package tables

import "gorm.io/gorm"

type DAO interface {

	// db

	GetDB(older *gorm.DB) *gorm.DB

	// fetch

	Find(db *gorm.DB, id TableID) (*TableEntity, error)

	Query(db *gorm.DB, q *TableQuery) ([]*TableEntity, error)

	// modify

	Insert(db *gorm.DB, item *TableEntity) (*TableEntity, error)

	Update(db *gorm.DB, id TableID, callback func(item *TableEntity) error) (*TableEntity, error)

	Delete(db *gorm.DB, id TableID) error
}
