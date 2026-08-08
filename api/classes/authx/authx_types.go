package authx

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/api/dxo"
)

////////////////////////////////////////////////////////////////////////////////

type Action = dxo.AuthAction

type Mechanism = dxo.AuthMechanism

type Step = dxo.AuthStep

////////////////////////////////////////////////////////////////////////////////

// AuthVO ...
type AuthVO struct {
	dxo.BaseVO

	Items []*AuthDTO `json:"auth"` // 用于验证的信息
}

// AuthDTO 用于身份认证
type AuthDTO struct {
	dxo.BaseDTO

	// 采用的验证机制
	Mechanism Mechanism `json:"mechanism"`

	// 最终要执行的动作,
	// 例如: login(登录), sign_up(注册), reset_password(重置密码), 等
	Action Action `json:"action"`

	// 表示验证的步骤
	Step Step `json:"step"`

	// 将要验证的账号,
	// 它可能是 UserName | UserID | UserEmail | UserPhone | ...
	Account string `json:"account"`

	// 将要验证的机密内容, 例如：密码, 等...
	Secret lang.Base64 `json:"secret"`

	// 其它扩展属性
	Parameters map[string]string `json:"parameters"`

	////////////////
	// results

	OK        bool
	Error     string
	Message   string
	Challenge string
}

////////////////////////////////////////////////////////////////////////////////
// EOF
