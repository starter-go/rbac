package mem

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"gorm.io/gorm"
)

type IAuthenticationDAO interface {
	rbac.AuthenticationDAO
}

type IPermissionDAO interface {
	rbac.PermissionDAO
}

type IRoleDAO interface {
	rbac.RoleDAO
}

type ISessionDAO interface {
	rbac.SessionDAO
}

type ITableDAO interface {
	rbac.TableDAO
}

type IUserDao interface {
	rbac.UserDAO
}

type IMemoryEngine interface {
	NewLS() *MemoryLS

	InitTable(model any) error

	NextUUID() lang.UUID

	GetDB(old *gorm.DB) *gorm.DB
}
