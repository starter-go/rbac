package permissions

import (
	"context"

	"github.com/starter-go/rbac/api/dxo"
)

// PermissionCache 是一个带缓存的 Permission 查询接口
type PermissionCache interface {
	Clear()
	Find(c context.Context, want *DTO) (*DTO, error)
}

type PermissionChecker interface {

	// 添加用户具有的角色
	AddRolesHad(roles ...dxo.RoleName) PermissionChecker

	// 添加可接受访问的用户角色
	AddRolesAccepted(roles ...dxo.RoleName) PermissionChecker

	// 接受匿名者访问
	AcceptAnonymous() PermissionChecker

	// 检查目标对象的访问权限
	CheckObject(o dxo.DTORef) error

	// 完成最终的检查
	Check() error
}
