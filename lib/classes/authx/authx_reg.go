package authx

// Registration 是授权器和验证器和用的注册信息结构
type Registration struct {
	Label     string
	Mechanism Mechanism // for Authenticator
	Action    Action    // for Authorizer
	Priority  int
	Enabled   bool

	Authenticator Authenticator
	Authorizer    Authorizer
}

// Registry for inject
type Registry interface {
	ListRegistrations() []*Registration
}
