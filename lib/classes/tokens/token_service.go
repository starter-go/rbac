package tokens

import "context"

// TokenService ...
type TokenService interface {
	GetCurrent(c context.Context) (*TokenDTO, error)

	PutCurrent(c context.Context, token *TokenDTO) (*TokenDTO, error)
}
