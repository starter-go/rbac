package unsupported

import "github.com/starter-go/rbac"

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

type IUserDao interface {
	rbac.UserDAO
}
