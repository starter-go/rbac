package rbac

import (
	"context"

	"time"

	"github.com/starter-go/base/lang"
)

// TokenDTO 表示令牌信息
type TokenDTO struct {

	//// BaseDTO

	NotBefore lang.Time `json:"not_before"`

	NotAfter lang.Time `json:"not_after"`

	SessionID SessionID `json:"session_id"` // 会话的 ID

	SessionUUID SessionUUID `json:"session_uuid"` // 会话的 UUID

}

func (inst *TokenDTO) BindWithSession(se *SessionDTO) *TokenDTO {
	if se != nil && inst != nil {
		inst.SessionID = se.ID
		inst.SessionUUID = SessionUUID(se.UUID)
	}
	return inst
}

func (inst *TokenDTO) SetTimeFromTo(from, to time.Time) *TokenDTO {
	t1 := lang.NewTime(from)
	t2 := lang.NewTime(to)
	inst.NotBefore = t1
	inst.NotAfter = t2
	return inst
}

func (inst *TokenDTO) SetTimeWithMaxAge(from time.Time, maxAge time.Duration) *TokenDTO {

	to := from.Add(maxAge)

	t1 := lang.NewTime(from)
	t2 := lang.NewTime(to)

	inst.NotBefore = t1
	inst.NotAfter = t2
	return inst
}

////////////////////////////////////////////////////////////////////////////////

// TokenService ...
type TokenService interface {
	GetCurrent(c context.Context) (*TokenDTO, error)

	PutCurrent(c context.Context, token *TokenDTO) (*TokenDTO, error)
}
