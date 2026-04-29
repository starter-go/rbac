package rbac

import (
	"github.com/starter-go/application/properties"
	"github.com/starter-go/rbac/localization"
)

// CurrentUser 表示当前用户
type CurrentUser struct {

	// 有效期
	Term

	UserID   UserID              `json:"user"`
	Username UserName            `json:"username"`
	Nickname string              `json:"nickname"`
	Language localization.Locale `json:"language"` // 该用户的本地化语言，取值示例：("zh_cn")
	Avatar   string              `json:"avatar"`
	Email    EmailAddress        `json:"email"`
	Roles    RoleNameList        `json:"roles"`

	Properties properties.Map `json:"properties"`

	AuthResults []*AuthResult `json:"auth_results"` // 主要用于 2FA
}

// GetProperty ...
func (inst *CurrentUser) GetProperty(name string) string {
	table := inst.Properties
	if table == nil {
		return ""
	}
	return table[name]
}

// SetProperty ...
func (inst *CurrentUser) SetProperty(name, value string) {
	table := inst.Properties
	if table == nil {
		table = make(map[string]string)
		inst.Properties = table
	}
	table[name] = value
}

// // SetPeriod 以 (t0 + max_age) 的形式设置 StartedAt & ExpiredAt
// func (inst *CurrentUser) SetPeriod(from time.Time, maxAge time.Duration) {
// 	const min = time.Millisecond
// 	if maxAge < min {
// 		maxAge = min
// 	}
// 	t0 := lang.NewTime(from)
// 	t1 := t0.Add(maxAge)
// 	inst.StartedAt = t0
// 	inst.ExpiredAt = t1
// }
