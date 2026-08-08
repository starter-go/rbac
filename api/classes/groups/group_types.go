package groups

import (
	"github.com/starter-go/rbac/api/dxo"
)

////////////////////////////////////////////////////////////////////////////////

type GroupID = dxo.GroupID

type GroupName = dxo.GroupName

////////////////////////////////////////////////////////////////////////////////

type ID = GroupID

type Name = GroupName

type Pagination = dxo.Pagination

type VO = GroupVO

type DTO = GroupDTO

type Entity = GroupEntity

////////////////////////////////////////////////////////////////////////////////

// GroupDTO 表示 Group 的 REST 网络对象
type GroupDTO struct {
	ID GroupID `json:"id"`

	dxo.BaseDTO

	Name        GroupName `json:"name"`
	Label       string    `json:"label"`
	Description string    `json:"description"`
}

type GroupEntity struct {
	ID GroupID

	dxo.BaseEntity

	Name        GroupName `gorm:"unique"`
	Label       string
	Description string
}

type GroupVO struct {
	dxo.BaseVO

	Items []*GroupDTO `json:"groups"`
}

////////////////////////////////////////////////////////////////////////////////

var theTableNamer dxo.RbacTableNamer

func (GroupEntity) TableName() string {
	return theTableNamer.GetFullTableName("groups")
}

////////////////////////////////////////////////////////////////////////////////
