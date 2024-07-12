package rbac

import "context"

// SubjectID 是通用的操作主体标识符
type SubjectID int64

// SubjectDTO 表示 操作主体 的 REST 网络对象
type SubjectDTO struct {
	ID SubjectID `json:"id"`

	BaseDTO

	// NickName      string       `json:"nickname"`
	// Avatar        string       `json:"avatar"`
	// Roles         RoleNameList `json:"roles"`

	CurrentUser

	Authenticated bool        `json:"authenticated"` // 是否已验证
	Token         *TokenDTO   `json:"token"`
	Session       *SessionDTO `json:"session"`
}

// SubjectService 是针对 SubjectDTO 的服务
type SubjectService interface {
	GetCurrent(c context.Context) (*SubjectDTO, error)
}
