package rbac

// EmailAddressID ...
type EmailAddressID int64

// GroupID 是通用的资源组标识符
type GroupID int64

// PermissionID 是 Permission 的实体 ID
type PermissionID int64

// PhoneNumberID ...
type PhoneNumberID int64

// RoleID 是 Role 的实体 ID
type RoleID int64

// SessionIID : int-id of session
// SessionID  用 int64 来作为会话的标识符
type SessionIID int64
type SessionID = SessionIID

// TableID 是 数据表-meta-info 的标识符
type TableID int64

// UserID 是通用的用户标识符
type UserID int64
type UserGroupID int64
type UserAtGroupID int64
