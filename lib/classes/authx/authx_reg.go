package authx

type Registration struct {
	Enabled bool

	Priority int

	Action Action // for Authorizer

	Mechanism Mechanism // for Authenticator

	Authenticator Authenticator // optional

	Authorizer Authorizer // optional
}

type Registry interface {
	Registration() *Registration
}
