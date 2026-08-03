package rbac

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/lib/classes/authentications"
	"github.com/starter-go/rbac/lib/classes/checkers"
	"github.com/starter-go/rbac/lib/classes/groups"
	"github.com/starter-go/rbac/lib/classes/permissions"
	"github.com/starter-go/rbac/lib/classes/regions"
	"github.com/starter-go/rbac/lib/classes/roles"
	"github.com/starter-go/rbac/lib/classes/sessions"
	"github.com/starter-go/rbac/lib/classes/tables"
	"github.com/starter-go/rbac/lib/classes/users"
	"github.com/starter-go/rbac/lib/dxo"
)

////////////////////////////////////////////////////////////////////////////////
// ids

// type AuthenticationID = dxo.AuthenticationID

// EmailAddressID ...
// type EmailAddressID = dxo.EmailAddressID

// GroupID 是通用的资源组标识符
// type GroupID = dxo.GroupID

// PermissionID 是 Permission 的实体 ID
// type PermissionID = dxo.PermissionID

// PhoneNumberID ...
type PhoneNumberID = dxo.PhoneNumberID

// RoleID 是 Role 的实体 ID
// type RoleID = dxo.RoleID

// SessionIID : int-id of session
// SessionID  用 int64 来作为会话的标识符
type SessionIID = dxo.SessionIID
type SessionID = dxo.SessionIID

// TableID 是 数据表-meta-info 的标识符
// type TableID = dxo.TableID

// UserID 是通用的用户标识符
// type UserID = dxo.UserID
type UserGroupID = dxo.UserGroupID
type UserAtGroupID = dxo.UserAtGroupID

////////////////////////////////////////////////////////////////////////////////
// authentications

type AuthenticationID = authentications.ID

type AuthenticationEntity = authentications.Entity

type AuthenticationDAO = authentications.DAO

type AuthenticationService = authentications.Service

////////////////////////////////////////////////////////////////////////////////
// checkers

type CheckerService = checkers.Service

type CheckerChain = checkers.CheckerChain

type CheckerRegistry = checkers.CheckerRegistry

////////////////////////////////////////////////////////////////////////////////
// email

type EmailAddress = dxo.EmailAddress

type EmailAddressID = dxo.EmailAddressID

////////////////////////////////////////////////////////////////////////////////
// groups

type GroupID = groups.GroupID

type GroupVO = groups.GroupVO

type GroupDTO = groups.GroupDTO

type GroupEntity = groups.GroupEntity

type GroupDAO = groups.DAO

type GroupService = groups.Service

////////////////////////////////////////////////////////////////////////////////
// page

type Pagination = dxo.Pagination

////////////////////////////////////////////////////////////////////////////////
// locations

type URI = lang.URI

type URL = lang.URL

////////////////////////////////////////////////////////////////////////////////
// permissions

type PermissionID = permissions.ID

type PermissionEntity = permissions.Entity

type PermissionDTO = permissions.DTO

type PermissionVO = permissions.VO

type PermissionDAO = permissions.DAO

type PermissionService = permissions.Service

////////////////////////////////////////////////////////////////////////////////
// phone

type PhoneNumber = dxo.PhoneNumber

type FullPhoneNumber = dxo.FullPhoneNumber

type SimplePhoneNumber = dxo.SimplePhoneNumber

type PurePhoneNumber = dxo.PurePhoneNumber

////////////////////////////////////////////////////////////////////////////////
// refs

type EntityRef = dxo.EntityRef

type DTORef = dxo.DTORef

type VORef = dxo.VORef

////////////////////////////////////////////////////////////////////////////////
// regions

type RegionID = regions.ID

type RegionDTO = regions.DTO

type RegionVO = regions.VO

type RegionEntity = regions.Entity

type RegionDAO = regions.DAO

type RegionService = regions.Service

////////////////////////////////////////////////////////////////////////////////
// role

type RoleID = roles.ID

type RoleName = roles.Name

type RoleNameList = dxo.RoleNameList

type RoleEntity = roles.Entity

type RoleDTO = roles.DTO

type RoleVO = roles.VO

type RoleDAO = roles.DAO

type RoleService = roles.Service

////////////////////////////////////////////////////////////////////////////////
// session

type SessionDTO = sessions.DTO

type SessionVO = sessions.VO

type SessionEntity = sessions.Entity

type SessionDAO = sessions.DAO

type SessionService = sessions.SessionService

////////////////////////////////////////////////////////////////////////////////
// tables

type TableID = tables.ID

type TableName = tables.Name

type TableDTO = tables.DTO

type TableVO = tables.VO

type TableEntity = tables.Entity

type TableDAO = tables.DAO

type TableService = tables.Service

////////////////////////////////////////////////////////////////////////////////
// user

type UserID = users.UserID

type UserName = users.UserName

type UserVO = users.UserVO

type UserDTO = users.UserDTO

type UserEntity = users.UserEntity

type UserInfo = dxo.UserInfo

type UserDAO = users.DAO

type UserService = users.Service

////////////////////////////////////////////////////////////////////////////////
// EOF
