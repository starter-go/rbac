package rbac

import (
	"time"

	"github.com/starter-go/base/lang"
)

// CurrentUser 表示当前用户
type CurrentUser struct {
	User      UserID       `json:"user"`
	Nickname  string       `json:"nickname"`
	Language  string       `json:"language"` // 该用户的本地化语言，取值示例：("zh_cn")
	Avatar    string       `json:"avatar"`
	Roles     RoleNameList `json:"roles"`
	StartedAt lang.Time    `json:"started_at"` // 生效时间
	ExpiredAt lang.Time    `json:"expired_at"` // 过期时间

	Properties map[string]string `json:"properties"`
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

// SetPeriod 以 (t0 + max_age) 的形式设置 StartedAt & ExpiredAt
func (inst *CurrentUser) SetPeriod(from time.Time, maxAge time.Duration) {
	const min = time.Millisecond
	if maxAge < min {
		maxAge = min
	}
	t0 := lang.NewTime(from)
	t1 := t0.Add(maxAge)
	inst.StartedAt = t0
	inst.ExpiredAt = t1
}
