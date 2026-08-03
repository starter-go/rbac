package authx

import (
	"github.com/starter-go/application/parameters"
)

// 授权 api

// 表示一个授权请求
type Authorization struct {

	// want

	Context *Context

	Action Action

	Params parameters.Table

	// have

	Message string
	OK      bool
	Error   error
}

// 表示授权组件接口
type Authorizer interface {
	Accept(a2 *Authorization) bool

	Authorize(a2 *Authorization) error
}

// // 包含授权组件注册信息
// type AuthorizerRegistration struct {
// 	Authorizer Authorizer
// 	Action Action
// 	Enabled bool
// 	Label string
// 	Priority int
// }
