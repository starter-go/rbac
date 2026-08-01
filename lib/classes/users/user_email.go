package users

import (
	"context"

	"github.com/starter-go/rbac/lib/dxo"
	"gorm.io/gorm"
)

type EmailAddressID = dxo.EmailAddressID

type EmailAddress = dxo.EmailAddress

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

// EmailAddressQuery 查询参数
type EmailAddressQuery struct {
	All bool // 查询全部条目

	Pagination Pagination

	Want *EmailAddressEntity
}

////////////////////////////////////////////////////////////////////////////////

// EmailAddressService ...
type EmailAddressService interface {
	Insert(c context.Context, o *EmailAddressDTO) (*EmailAddressDTO, error)
	Update(c context.Context, id EmailAddressID, o *EmailAddressDTO) (*EmailAddressDTO, error)
	Delete(c context.Context, id EmailAddressID) error

	Find(c context.Context, id EmailAddressID) (*EmailAddressDTO, error)
	List(c context.Context, q *EmailAddressQuery) ([]*EmailAddressDTO, error)
}

////////////////////////////////////////////////////////////////////////////////

// EmailAddressDAO ...
type EmailAddressDAO interface {

	// edit

	Insert(db *gorm.DB, o *EmailAddressEntity) (*EmailAddressEntity, error)

	Update(db *gorm.DB, id EmailAddressID, callback func(old *EmailAddressEntity) error) (*EmailAddressEntity, error)

	Delete(db *gorm.DB, id EmailAddressID) error

	// query

	Find(db *gorm.DB, id EmailAddressID) (*EmailAddressEntity, error)

	List(db *gorm.DB, q *EmailAddressQuery) ([]*EmailAddressEntity, error)
}

////////////////////////////////////////////////////////////////////////////////
