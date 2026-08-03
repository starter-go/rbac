package sessions

import (
	"github.com/starter-go/application/properties"
	"github.com/starter-go/rbac/lib/dxo"
)

////////////////////////////////////////////////////////////////////////////////

type SessionID = dxo.SessionID

type SessionUUID = dxo.SessionUUID

////////////////////////////////////////////////////////////////////////////////

type ID = SessionID

type UUID = SessionUUID

type Pagination = dxo.Pagination

type Entity = SessionEntity

type DTO = SessionDTO

type VO = SessionVO

////////////////////////////////////////////////////////////////////////////////

type SessionEntity struct {
	ID SessionID

	dxo.BaseEntity

	Properties properties.Text

	Authenticated bool // 是否已验证

}

// SessionDTO 表示会话信息
type SessionDTO struct {
	ID SessionID `json:"id"`

	dxo.BaseDTO

	Authenticated bool `json:"authenticated"` // 是否已验证

	dxo.Term

	dxo.UserInfo
}

type SessionVO struct {
	dxo.BaseVO

	Items []*SessionDTO `json:"sessions"`
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamer dxo.RbacTableNamer

func (SessionEntity) TableName() string {
	return theTableNamer.GetFullTableName("sessions")
}

////////////////////////////////////////////////////////////////////////////////
