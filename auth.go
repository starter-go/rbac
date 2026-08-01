package rbac

import (
	"context"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/lib/classes/authx"
	"github.com/starter-go/rbac/lib/dxo"
)

////////////////////////////////////////////////////////////////////////////////

type AuthAction = authx.Action

type AuthMechanism = authx.Mechanism

type AuthStep = authx.Step

type AuthRegistry = authx.Registry

////////////////////////////////////////////////////////////////////////////////

// 定义几种常用的授权动作
const (
	ActionLogin          AuthAction = "login"
	ActionSignUp         AuthAction = "sign-up"
	ActionResetPassword  AuthAction = "reset-password"
	ActionChangePassword AuthAction = "change-password"
	ActionSendCode       AuthAction = "send-code"
)

// 定义几种常用的验证机制
const (
	MechanismPassword AuthMechanism = "password"
	MechanismEmail    AuthMechanism = "email"
	MechanismPhone    AuthMechanism = "sms"
	MechanismSMS      AuthMechanism = "sms"
)

// 定义几个常用的验证步骤
const (
	StepInit     AuthStep = "init"     // 初始化
	StepPrepare  AuthStep = "prepare"  // 准备
	StepHelp     AuthStep = "help"     // 获取帮助信息
	StepSendCode AuthStep = "sendcode" // 发送验证码
	StepApply    AuthStep = "apply"    // 应用
	StepAuth     AuthStep = "auth"     // 验证与授权
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

// AuthResult 表示认证结果
type AuthResult struct {

	// 有效期:
	dxo.Term

	Mechanism  string `json:"mechanism"` // 采用的认证机制
	UserID     UserID `json:"uid"`       // 被认证的用户ID
	DomainName string `json:"domain"`    // 认证针对的域名
	OK         string `json:"ok"`        // 是否认证成功

}

// AuthService 是针对 AuthDTO 的服务
type AuthService interface {
	Handle(c context.Context, action string, a []*AuthDTO) ([]*AuthDTO, error)
}
