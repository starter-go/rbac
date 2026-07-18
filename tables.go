package rbac

import (
	"context"

	"github.com/starter-go/base/lang"
	"gorm.io/gorm"
)

// TableXxx  是用于描述 DB中的表格的元数据

type TableName string

type TableDTO struct {
	ID TableID

	DTO

	Name TableName `json:"name"` // the table-name

	GroupName string
	GroupURI  lang.URI
	TableURI  lang.URI

	Label       string
	Description string
}

type TableEntity struct {
	ID TableID

	Entity

	Name TableName `gorm:"unique"` // the table-name

	GroupName string
	GroupURI  lang.URI
	TableURI  lang.URI

	Label       string
	Description string
}

type TableQuery struct {
	All        bool
	Pagination Pagination
	Want       *TableEntity
}

type TableService interface {

	// fetch

	Find(c context.Context, id TableID) (*TableDTO, error)

	Query(c context.Context, q *TableQuery) ([]*TableDTO, error)

	// modify

	Insert(c context.Context, item *TableDTO) (*TableDTO, error)

	Update(c context.Context, id TableID, item *TableDTO) (*TableDTO, error)

	Delete(c context.Context, id TableID) error
}

type TableDAO interface {
	GetDB(older *gorm.DB) *gorm.DB

	// fetch

	Find(db *gorm.DB, id TableID) (*TableEntity, error)

	Query(db *gorm.DB, q *TableQuery) ([]*TableEntity, error)

	// modify

	Insert(db *gorm.DB, item *TableEntity) (*TableEntity, error)

	Update(db *gorm.DB, id TableID, callback func(item *TableEntity) error) (*TableEntity, error)

	Delete(db *gorm.DB, id TableID) error
}
