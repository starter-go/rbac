package rbac

import (
	"context"
	"strconv"

	"gorm.io/gorm"
)

////////////////////////////////////////////////////////////////////////////////

type ResourceGroupID = GroupID

type ResourceGroupName = GroupName

type ResourceGroupDTO = GroupDTO

type ResourceGroupEntity = GroupEntity

type ResourceGroupService = GroupService

type ResourceGroupDAO = GroupDAO

type ResourceGroupQuery = GroupQuery

////////////////////////////////////////////////////////////////////////////////

// GroupName 表示资源分组名称
type GroupName string

// GroupDTO 表示 Group 的 REST 网络对象
type GroupDTO struct {
	ID GroupID `json:"id"`

	DTO

	Name        GroupName `json:"name"`
	Label       string    `json:"label"`
	Description string    `json:"description"`

	// Enabled bool `json:"enabled"`
	// Roles   RoleNameList `json:"roles"`

}

type GroupEntity struct {
	ID GroupID

	Entity

	Name        GroupName `gorm:"unique"`
	Label       string
	Description string

	// Roles   RoleNameList
	// Enabled bool

}

// GroupQuery 是 Group 的查询参数
type GroupQuery struct {
	// Conditions Conditions

	Pagination Pagination
	All        bool // 查询全部条目
	Want       *GroupEntity
}

////////////////////////////////////////////////////////////////////////////////

// GroupService 是针对 GroupDTO 的服务
type GroupService interface {

	//fetch

	Find(c context.Context, id GroupID) (*GroupDTO, error)

	Query(c context.Context, q *GroupQuery) ([]*GroupDTO, error)

	// edit

	Insert(c context.Context, o *GroupDTO) (*GroupDTO, error)

	Update(c context.Context, id GroupID, o *GroupDTO) (*GroupDTO, error)

	Delete(c context.Context, id GroupID) error
}

// GroupDAO 是针对 GroupEntity 的 DAO
type GroupDAO interface {

	// fetch

	Find(db *gorm.DB, id GroupID) (*GroupEntity, error)

	Query(db *gorm.DB, q *GroupQuery) ([]*GroupEntity, error)

	// edit

	Insert(db *gorm.DB, item *GroupEntity) (*GroupEntity, error)

	Update(db *gorm.DB, id GroupID, callback func(older *GroupEntity) error) (*GroupEntity, error)

	Delete(db *gorm.DB, id GroupID) error
}

////////////////////////////////////////////////////////////////////////////////

func (id GroupID) String() string {
	n := int64(id)
	return strconv.FormatInt(n, 10)
}

// ParseGroupID 把字符串解析为 GroupID
func ParseGroupID(s string) (GroupID, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return GroupID(n), nil
}

////////////////////////////////////////////////////////////////////////////////

func (name GroupName) String() string {
	return string(name)
}

////////////////////////////////////////////////////////////////////////////////
