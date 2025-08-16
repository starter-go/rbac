package rbac

import (
	"context"
)

// TokenDTO 表示令牌信息
type TokenDTO struct {
	// BaseDTO

	SessionID   SessionID   `json:"session_id"`   // 会话的 ID
	SessionUUID SessionUUID `json:"session_uuid"` // 会话的 UUID

}

////////////////////////////////////////////////////////////////////////////////

// TokenService ...
type TokenService interface {
	GetCurrent(c context.Context) (*TokenDTO, error)

	PutCurrent(c context.Context, token *TokenDTO) (*TokenDTO, error)
}
