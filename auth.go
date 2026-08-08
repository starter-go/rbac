package rbac

import (
	"github.com/starter-go/rbac/api/classes/authx"
	"github.com/starter-go/rbac/api/dxo"
)

////////////////////////////////////////////////////////////////////////////////

type AuthAction = authx.Action

type AuthMechanism = authx.Mechanism

type AuthStep = authx.Step

type AuthRegistry = authx.Registry

type AuthDTO = authx.AuthDTO

type AuthVO = authx.AuthVO

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
	authx.Service
}
