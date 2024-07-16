package rbac

import (
	"context"

	"github.com/starter-go/base/lang"
)

// 定义几种常用的授权动作
const (
	ActionLogin          = "login"
	ActionSignUp         = "sign-up"
	ActionResetPassword  = "reset-password"
	ActionChangePassword = "change-password"
	ActionSendCode       = "send-code"
)

// 定义几种常用的验证机制
const (
	MechanismPassword = "password"
	MechanismEmail    = "email"
	MechanismPhone    = "sms"
	MechanismSMS      = "sms"
)

// AuthDTO 用于身份认证
type AuthDTO struct {
	BaseDTO

	// 采用的验证机制
	Mechanism string `json:"mechanism"`

	// 最终要执行的动作,
	// 例如: login(登录), sign_up(注册), reset_password(重置密码), 等
	Action string `json:"action"`

	// 表示验证的步骤
	Step string `json:"step"`

	// 将要验证的账号,
	// 它可能是 UserName | UserID | UserEmail | UserPhone | ...
	Account string `json:"account"`

	// 将要验证的机密内容, 例如：密码, 等...
	Secret lang.Base64 `json:"secret"`

	// 其它扩展属性
	Parameters map[string]string `json:"parameters"`
}

// AuthService 是针对 AuthDTO 的服务
type AuthService interface {
	Handle(c context.Context, action string, a []*AuthDTO) ([]*AuthDTO, error)
}
