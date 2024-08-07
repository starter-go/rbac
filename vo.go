package rbac

////////////////////////////////////////////////////////////////////////////////

// VOGetter 是获取 BaseVO 的接口
type VOGetter interface {
	GetVO() *BaseVO
}

// GetVO 实现 BaseGetter
func (inst *BaseVO) GetVO() *BaseVO {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

// AuthVO ...
type AuthVO struct {
	BaseVO

	Auth []*AuthDTO `json:"auth"` // 用于验证的信息
}

// AvatarVO ...
type AvatarVO struct {
	BaseVO

	Avatars []*AvatarDTO `json:"avatars"`
}

// PermissionVO ...
type PermissionVO struct {
	BaseVO

	Permissions []*PermissionDTO `json:"permissions"`
}

// RoleVO ...
type RoleVO struct {
	BaseVO

	Roles []*RoleDTO `json:"roles"`
}

// SessionVO ...
type SessionVO struct {
	BaseVO

	Sessions []*SessionDTO `json:"sessions"`
}

// UserVO ...
type UserVO struct {
	BaseVO

	Users []*UserDTO `json:"users"`
}

// GroupVO ...
type GroupVO struct {
	BaseVO

	Groups []*GroupDTO `json:"groups"`
}

// SubjectVO ...
type SubjectVO struct {
	BaseVO

	Subject *SubjectDTO `json:"subject"`
}

// TokenVO ...
type TokenVO struct {
	BaseVO

	Tokens []*TokenDTO `json:"tokens"`
}
