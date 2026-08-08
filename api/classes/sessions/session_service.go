package sessions

import "context"

// SessionService ...
type SessionService interface {
	// GetCurrent(c context.Context) (*SessionDTO, error)

	Find(c context.Context, id SessionID) (*SessionDTO, error)

	Insert(c context.Context, se *SessionDTO) (*SessionDTO, error)

	Update(c context.Context, id SessionID, se *SessionDTO) (*SessionDTO, error)
}
