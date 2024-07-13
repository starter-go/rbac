package rbac

import (
	"context"
)

// TokenDTO 表示令牌信息
type TokenDTO struct {
	// BaseDTO

	CurrentUser

	Session SessionID `json:"sessionid"` // 会话的 UUID

}

////////////////////////////////////////////////////////////////////////////////

// TokenService ...
type TokenService interface {
	GetCurrent(c context.Context) (*TokenDTO, error)

	PutCurrent(c context.Context, token *TokenDTO) (*TokenDTO, error)
}
