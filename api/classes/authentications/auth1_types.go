package authentications

// authentications 包含提供认证方案的 API

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/api/dxo"
)

////////////////////////////////////////////////////////////////////////////////

type ID = dxo.AuthenticationID

type Pagination = dxo.Pagination

////////////////////////////////////////////////////////////////////////////////

// authentications.Entity 表示一个具体的认证方案实例
type Entity struct {
	ID ID

	dxo.BaseEntity

	CommonName string
	Domain     dxo.DomainName
	Mechanism  string

	URI lang.URI `gorm:"unique"`

	Enabled bool
	Locked  bool

	Secret lang.Hex
	Salt   lang.Base64
}

type DTO struct {
	ID ID

	dxo.BaseDTO

	CommonName string         `json:"name"`
	Domain     dxo.DomainName `json:"domain"`
	Mechanism  string         `json:"mechanism"`

	URI lang.URI `json:"uri"`

	Enabled bool `json:"enabled"`
	Locked  bool `json:"locked"`
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
