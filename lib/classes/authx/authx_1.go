package authx

type Authenticator interface {
	Accept(a2 *Authentication) bool

	Authenticate(a2 *Authentication) error
}
