package emailaddresses

import "github.com/starter-go/rbac/api/dxo"

////////////////////////////////////////////////////////////////////////////////

type EmailAddressID = dxo.EmailAddressID

type EmailAddress = dxo.EmailAddress

type Pagination = dxo.Pagination

type Entity = EmailAddressEntity

type DTO = EmailAddressDTO

type VO = EmailAddressVO

type ID = dxo.EmailAddressID

////////////////////////////////////////////////////////////////////////////////

// EmailAddressEntity ...
type EmailAddressEntity struct {
	ID EmailAddressID

	dxo.BaseEntity

	Address EmailAddress `gorm:"unique"`
}

// EmailAddressDTO ...
type EmailAddressDTO struct {
	ID EmailAddressID `json:"id"`

	dxo.BaseDTO

	Address EmailAddress `json:"address"`
}

// EmailAddressVO ...
type EmailAddressVO struct {
	dxo.BaseVO

	Items []*EmailAddressDTO `gorm:"email_addresses"`
}

////////////////////////////////////////////////////////////////////////////////
// EOF
