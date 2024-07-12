package rbac

import "github.com/starter-go/base/lang"

// CurrentUser 表示当前用户
type CurrentUser struct {
	User      UserID       `json:"user"`
	Nickname  string       `json:"nickname"`
	Avatar    string       `json:"avatar"`
	Roles     RoleNameList `json:"roles"`
	StartedAt lang.Time    `json:"started_at"` // 生效时间
	ExpiredAt lang.Time    `json:"expired_at"` // 过期时间

	Properties map[string]string `json:"properties"`
}
