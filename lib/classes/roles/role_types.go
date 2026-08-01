package roles

import (
	"github.com/starter-go/rbac/lib/dxo"
)

type ID = dxo.RoleID

type Name = dxo.RoleName

type Pagination = dxo.Pagination

////////////////////////////////////////////////////////////////////////////////

// Role_DTO 表示 Role 的 REST 网络对象
type DTO struct {
	ID ID `json:"id"`

	dxo.BaseDTO

	Name        Name   `json:"name"`
	Description string `json:"description"`
}

type Entity struct {
	ID dxo.RoleID

	dxo.BaseEntity

	Name dxo.RoleName `gorm:"unique"`

	Description string
}

type VO struct {
	dxo.BaseVO

	Items []*DTO `json:"roles"`
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamer dxo.RbacTableNamer

func (Entity) TableName() string {
	return theTableNamer.GetFullTableName("roles")
}

////////////////////////////////////////////////////////////////////////////////
