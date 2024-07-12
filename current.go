package rbac

import "github.com/starter-go/base/lang"

// CurrentUser 表示当前用户
type CurrentUser struct {
	User      UserID       `json:"user"`
	Nickname  string       `json:"nickname"`
	Avatar    string       `json:"avatar"`
	Roles     RoleNameList `json:"roles"`
	CreatedAt lang.Time    `json:"created_at"`
	ExpiredAt lang.Time    `json:"expired_at"`

	Properties map[string]string `json:"properties"`
}
