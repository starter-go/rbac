package tokens

import (
	"time"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/lang"

	"github.com/starter-go/rbac/api/classes/sessions"
	"github.com/starter-go/rbac/api/dxo"
)

////////////////////////////////////////////////////////////////////////////////

type SessionID = dxo.SessionID

type SessionUUID = dxo.SessionUUID

type DTO = TokenDTO

type VO = TokenVO

////////////////////////////////////////////////////////////////////////////////

// TokenDTO 表示令牌信息
type TokenDTO struct {

	//// BaseDTO

	// NotBefore lang.Time `json:"not_before"`
	// NotAfter  lang.Time `json:"not_after"`

	dxo.Term

	SessionID SessionID `json:"session_id"` // 会话的 ID

	SessionUUID SessionUUID `json:"session_uuid"` // 会话的 UUID

	Properties properties.Map `json:"properties"` // 其它扩展的属性

}

// TokenVO 表示令牌信息
type TokenVO struct {

	// vo === dto
	TokenDTO
}

////////////////////////////////////////////////////////////////////////////////

func (inst *TokenDTO) BindWithSession(se *sessions.SessionDTO) *TokenDTO {
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
