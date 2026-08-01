package dxo

import "github.com/starter-go/base/lang"

// SessionIID : int-id of session
// SessionID  用 int64 来作为会话的标识符
type SessionIID int64

type SessionID = SessionIID

type SessionUUID = lang.UUID
