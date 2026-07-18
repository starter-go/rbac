package rbac

import (
	"context"
)

// PermissionDTO 表示 Permission 的 REST 网络对象
type PermissionDTO struct {
	ID PermissionID `json:"id"`

	BaseDTO

	Enabled bool         `json:"enabled"`
	Method  string       `json:"method"`
	Path    string       `json:"path"`
	Roles   RoleNameList `json:"roles"`
}

type PermissionEntity struct {
	ID PermissionID

	Entity

	URI string `gorm:"unique"`

	Enabled bool
	Method  string
	Path    string
	Roles   RoleNameList
}

// PermissionQuery 查询参数
type PermissionQuery struct {
	All        bool // 查询全部条目
	Pagination Pagination
	Want       *PermissionEntity
}

// PermissionService 是针对 PermissionDTO 的服务
type PermissionService interface {
	Insert(c context.Context, o *PermissionDTO) (*PermissionDTO, error)
	Update(c context.Context, id PermissionID, o *PermissionDTO) (*PermissionDTO, error)
	Delete(c context.Context, id PermissionID) error

	Find(c context.Context, id PermissionID) (*PermissionDTO, error)
	List(c context.Context, q *PermissionQuery) ([]*PermissionDTO, error)
	ListAll(c context.Context) ([]*PermissionDTO, error)

	GetCache() PermissionCache
}

// PermissionCache 是一个带缓存的 Permission 查询接口
type PermissionCache interface {
	Clear()
	Find(c context.Context, want *PermissionDTO) (*PermissionDTO, error)
}

type PermissionChecker interface {

	// 添加用户具有的角色
	AddRolesHad(roles ...RoleName) PermissionChecker

	// 添加可接受访问的用户角色
	AddRolesAccepted(roles ...RoleName) PermissionChecker

	// 接受匿名者访问
	AcceptAnonymous() PermissionChecker

	// 检查目标对象的访问权限
	CheckObject(o *BaseDTO) error

	// 完成最终的检查
	Check() error
}
