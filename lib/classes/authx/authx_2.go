package authx

type Authorizer interface {
	Accept(a2 *Authorization) bool

	Authorize(a2 *Authorization) error
}
