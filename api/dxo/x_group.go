package dxo

import "strconv"

////////////////////////////////////////////////////////////////////////////////

// type ResourceGroupID = GroupID

// type ResourceGroupName = GroupName

// type ResourceGroupDTO = GroupDTO

// type ResourceGroupEntity = GroupEntity

// type ResourceGroupService = GroupService

// type ResourceGroupDAO = GroupDAO

// type ResourceGroupQuery = GroupQuery

////////////////////////////////////////////////////////////////////////////////

// GroupID 是通用的资源组标识符
type GroupID int64

// GroupName 表示资源分组名称
type GroupName string

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
