package dxo

import (
	"strconv"

	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/localization"
)

////////////////////////////////////////////////////////////////////////////////

// UserName 表示用户名
type UserName string

// UserID 是通用的用户标识符
type UserID int64
type UserGroupID int64
type UserAtGroupID int64

////////////////////////////////////////////////////////////////////////////////

// UserInfo 用于在各个对象之间交换用户信息
type UserInfo struct {

	// 有效期
	Term

	UserID   UserID              `json:"uid"`
	Username UserName            `json:"username"`
	Nickname string              `json:"nickname"`
	Language localization.Locale `json:"language"` // 该用户的本地化语言，取值示例：("zh_cn")
	Avatar   lang.URL            `json:"avatar"`
	Email    EmailAddress        `json:"email"`
	Roles    RoleNameList        `json:"roles"`

	Properties properties.Map `json:"properties"`

	// AuthResults []*AuthResult `json:"auth_results"` // 主要用于 2FA
}

////////////////////////////////////////////////////////////////////////////////

func (name UserName) String() string {
	return string(name)
}

////////////////////////////////////////////////////////////////////////////////

func (id UserID) String() string {
	n := int64(id)
	return strconv.FormatInt(n, 10)
}

// ParseUserID 把字符串解析为 UserID
func ParseUserID(s string) (UserID, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return UserID(n), nil
}

////////////////////////////////////////////////////////////////////////////////

// GetProperty ...
func (inst *UserInfo) GetProperty(name string) string {
	table := inst.Properties
	if table == nil {
		return ""
	}
	return table[name]
}

// SetProperty ...
func (inst *UserInfo) SetProperty(name, value string) {
	table := inst.Properties
	if table == nil {
		table = make(map[string]string)
		inst.Properties = table
	}
	table[name] = value
}

////////////////////////////////////////////////////////////////////////////////
