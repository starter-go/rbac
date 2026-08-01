package groups

import "gorm.io/gorm"

// Group_DAO 是针对 GroupEntity 的 DAO
type DAO interface {

	// fetch

	Find(db *gorm.DB, id GroupID) (*GroupEntity, error)

	Query(db *gorm.DB, q *Query) ([]*GroupEntity, error)

	// edit

	Insert(db *gorm.DB, item *GroupEntity) (*GroupEntity, error)

	Update(db *gorm.DB, id GroupID, callback func(older *GroupEntity) error) (*GroupEntity, error)

	Delete(db *gorm.DB, id GroupID) error
}
