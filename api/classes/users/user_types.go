package users

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/api/dxo"
	"github.com/starter-go/rbac/api/localization"
)

////////////////////////////////////////////////////////////////////////////////

type UserID = dxo.UserID

type UserName = dxo.UserName

////////////////////////////////////////////////////////////////////////////////
// short aliases

type ID = UserID

type Name = UserName

type Entity = UserEntity

type DTO = UserDTO

type VO = UserVO

type Pagination = dxo.Pagination

type Query = UserQuery

type Service = UserService

type DAO = UserDAO

type PhoneNumberID = dxo.PhoneNumberID

type PhoneNumber = dxo.PhoneNumber

type EmailAddress = dxo.EmailAddress

type DomainName = dxo.DomainName

////////////////////////////////////////////////////////////////////////////////

type UserEntity struct {
	ID UserID

	dxo.BaseEntity

	Name  Name         `gorm:"unique"`
	Phone PhoneNumber  `gorm:"unique"`
	Email EmailAddress `gorm:"unique"`

	Nickname string
	Avatar   lang.URL
	Language localization.Locale
	Roles    dxo.RoleNameList
	Domain   DomainName

	Password lang.Hex
	Salt     lang.Base64

	Use2FA  bool
	Enabled bool
	Locked  bool
}

// User_DTO 表示 User 的 REST 网络对象
type UserDTO struct {
	ID UserID `json:"id"`

	dxo.BaseDTO

	Name     Name                `json:"name"`
	Nickname string              `json:"nickname"`
	Avatar   lang.URL            `json:"avatar"`
	Phone    PhoneNumber         `json:"phone"`
	Email    EmailAddress        `json:"email"`
	Language localization.Locale `json:"language"`
	Roles    dxo.RoleNameList    `json:"roles"`
	Domain   DomainName          `json:"domain"`

	Enabled bool `json:"enabled"`
	Locked  bool `json:"locked"`
	Use2FA  bool `json:"2fa"`
}

type UserVO struct {
	dxo.BaseVO

	Items []*UserDTO `json:"users"`
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamer dxo.RbacTableNamer

func (Entity) TableName() string {
	return theTableNamer.GetFullTableName("users")
}

////////////////////////////////////////////////////////////////////////////////
