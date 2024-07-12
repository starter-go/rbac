package rbac

import "github.com/starter-go/base/lang"

// TokenDTO 表示令牌信息
type TokenDTO struct {
	// BaseDTO

	CurrentUser

	Session lang.UUID `json:"session"` // 会话的 UUID

}
