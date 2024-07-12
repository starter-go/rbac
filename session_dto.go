package rbac

// SessionDTO 表示会话信息
type SessionDTO struct {
	BaseDTO
	CurrentUser

	Authenticated bool `json:"authenticated"` // 是否已验证
}
