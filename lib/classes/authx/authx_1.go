package authx

import (
	"github.com/starter-go/application/parameters"
	"github.com/starter-go/rbac/lib/dxo"
)

// 验证 api

// Authentication 表示一个验证请求
type Authentication struct {

	// want

	Context *Context

	CommonName string // [user_id, phone_num, email_addr, ...]

	Mechanism Mechanism
	Step      Step
	Params    parameters.Table
	Secret    []byte

	// have

	Challenge string
	Message   string
	Error     error
	OK        bool

	UserInfo *dxo.UserInfo
}

// Authenticator 表示验证器接口
type Authenticator interface {
	Accept(a2 *Authentication) bool

	Authenticate(a2 *Authentication) error
}

// AuthenticatorRegistration 结构包含验证器注册信息
type AuthenticatorRegistration struct {
	Authenticator Authenticator

	Enabled bool

	Label string

	Mechanism Mechanism

	Priority int
}
