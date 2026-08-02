package authx

// Registration 是授权器和验证器和用的注册信息结构
type Registration struct {
	A1 *AuthenticatorRegistration // optional
	A2 *AuthorizerRegistration    // optional
}

type Registry interface {
	ListRegistrations() []*Registration
}
