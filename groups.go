package rbac

import (
	"context"
	"strconv"
)

// GroupID 是通用的用户组标识符
type GroupID int64

// GroupName 表示用户分组名称
type GroupName string

// GroupDTO 表示 Group 的 REST 网络对象
type GroupDTO struct {
	ID GroupID `json:"id"`

	BaseDTO

	Name        GroupName `json:"name"`
	Label       string    `json:"label"`
	Description string    `json:"description"`

	Roles   RoleNameList `json:"roles"`
	Enabled bool         `json:"enabled"`
}

// GroupQuery 是 Group 的查询参数
type GroupQuery struct {
	Conditions Conditions
	Pagination Pagination
	All        bool // 查询全部条目
}

// GroupService 是针对 GroupDTO 的服务
type GroupService interface {
	Insert(c context.Context, o *GroupDTO) (*GroupDTO, error)
	Update(c context.Context, id GroupID, o *GroupDTO) (*GroupDTO, error)
	Delete(c context.Context, id GroupID) error

	Find(c context.Context, id GroupID) (*GroupDTO, error)
	List(c context.Context, q *GroupQuery) ([]*GroupDTO, error)
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
