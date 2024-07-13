package rbac

import (
	"context"

	"github.com/starter-go/base/lang"
)

// SessionID ... 用 uuid 来作为会话的标识符
type SessionID lang.UUID

// SessionDTO 表示会话信息
type SessionDTO struct {
	ID SessionID `json:"id"`

	BaseDTO
	CurrentUser

	Authenticated bool `json:"authenticated"` // 是否已验证
}

// SessionService ...
type SessionService interface {
	// GetCurrent(c context.Context) (*SessionDTO, error)

	Find(c context.Context, id SessionID) (*SessionDTO, error)

	Insert(c context.Context, se *SessionDTO) (*SessionDTO, error)

	Update(c context.Context, id SessionID, se *SessionDTO) (*SessionDTO, error)
}
