package tables

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/lib/dxo"
)

// TableXxx  是用于描述 DB中的表格的元数据

////////////////////////////////////////////////////////////////////////////////

type TableName = dxo.TableName

type TableID = dxo.TableID

type Pagination = dxo.Pagination

////////////////////////////////////////////////////////////////////////////////

type ID = TableID

type Name = TableName

type Entity = TableEntity

type DTO = TableDTO

type VO = TableVO

type Query = TableQuery

////////////////////////////////////////////////////////////////////////////////

type TableDTO struct {
	ID TableID

	dxo.BaseDTO

	Name TableName `json:"name"` // the table-name

	GroupName string
	GroupURI  lang.URI
	TableURI  lang.URI

	Label       string
	Description string
}

type TableEntity struct {
	ID TableID

	dxo.BaseEntity

	Name TableName `gorm:"unique"` // the table-name

	GroupName string
	GroupURI  lang.URI
	TableURI  lang.URI

	Label       string
	Description string
}

type TableVO struct {
	dxo.BaseVO

	Items []*DTO `json:"tables"`
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamer dxo.RbacTableNamer

func (TableEntity) TableName() string {
	return theTableNamer.GetFullTableName("tables")
}

////////////////////////////////////////////////////////////////////////////////
// EOF
