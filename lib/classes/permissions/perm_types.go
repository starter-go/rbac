package permissions

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/lib/dxo"
)

type ID = dxo.PermissionID

type Pagination = dxo.Pagination

////////////////////////////////////////////////////////////////////////////////

// Permission_Entity
type Entity struct {
	ID ID

	dxo.BaseEntity

	Enabled  bool
	Method   string
	Path     string
	Priority int
	Roles    dxo.RoleNameList

	URI lang.URI `gorm:"unique"`
}

// Permission_DTO 表示 Permission 的 REST 网络对象
type DTO struct {
	ID ID `json:"id"`

	dxo.DTO

	Enabled  bool             `json:"enabled"`
	Method   string           `json:"method"`
	Path     string           `json:"path"`
	Priority int              `json:"priority"`
	Roles    dxo.RoleNameList `json:"roles"`
	URI      lang.URI         `json:"uri"`
}

type VO struct {
	dxo.BaseVO

	Items []*DTO `json:"permissions"`
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamer dxo.RbacTableNamer

func (Entity) TableName() string {
	return theTableNamer.GetFullTableName("permissions")
}

////////////////////////////////////////////////////////////////////////////////
