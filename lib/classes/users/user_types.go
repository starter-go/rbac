package users

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/lib/dxo"
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

////////////////////////////////////////////////////////////////////////////////

type UserEntity struct {
	ID UserID

	dxo.BaseEntity

	Name  Name         `gorm:"unique"`
	Phone PhoneNumber  `gorm:"unique"`
	Email EmailAddress `gorm:"unique"`

	NickName string
	Avatar   string
	Language string
	Roles    dxo.RoleNameList
	Enabled  bool
	Locked   bool

	Password lang.Hex
	Salt     lang.Hex

	Use2FA bool
}

// User_DTO 表示 User 的 REST 网络对象
type UserDTO struct {
	ID UserID `json:"id"`

	dxo.BaseDTO

	Name     Name             `json:"name"`
	NickName string           `json:"nickname"`
	Avatar   string           `json:"avatar"`
	Phone    PhoneNumber      `json:"phone"`
	Email    EmailAddress     `json:"email"`
	Language string           `json:"language"`
	Roles    dxo.RoleNameList `json:"roles"`
	Enabled  bool             `json:"enabled"`
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
