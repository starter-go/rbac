package authentications

// authentications 包含提供认证方案的 API

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/lib/dxo"
)

////////////////////////////////////////////////////////////////////////////////

type ID = dxo.AuthenticationID

type Pagination = dxo.Pagination

////////////////////////////////////////////////////////////////////////////////

// authentications.Entity 表示一个具体的认证方案实例
type Entity struct {
	ID ID

	dxo.BaseEntity

	Mechanism string

	CommonName string

	URI lang.URI `gorm:"unique"`

	Sum lang.Hex

	Salt lang.Base64

	Enabled bool

	Locked bool
}

type DTO struct {
	ID ID

	dxo.BaseDTO
}

type VO struct {
	dxo.BaseVO

	Items []*DTO `json:"authentications"`
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamer dxo.RbacTableNamer

func (Entity) TableName() string {
	return theTableNamer.GetFullTableName("authentications")
}

////////////////////////////////////////////////////////////////////////////////
