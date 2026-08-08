package rbac

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/api/classes/authentications"
	"github.com/starter-go/rbac/api/classes/authx"
	"github.com/starter-go/rbac/api/classes/checkers"
	"github.com/starter-go/rbac/api/classes/emailaddresses"
	"github.com/starter-go/rbac/api/classes/groups"
	"github.com/starter-go/rbac/api/classes/permissions"
	"github.com/starter-go/rbac/api/classes/phonenumbers"
	"github.com/starter-go/rbac/api/classes/regions"
	"github.com/starter-go/rbac/api/classes/roles"
	"github.com/starter-go/rbac/api/classes/sessions"
	"github.com/starter-go/rbac/api/classes/tables"
	"github.com/starter-go/rbac/api/classes/tokens"
	"github.com/starter-go/rbac/api/classes/users"
	"github.com/starter-go/rbac/api/dxo"
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

// PhoneNumberID 是电话号码记录(PhoneNumberEntity)的实体 ID
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

// UserGroupID 是用户组(UserGroup)的实体 ID
type UserGroupID = dxo.UserGroupID

// UserAtGroupID 是用户组成员关系(UserAtGroup)的实体 ID
type UserAtGroupID = dxo.UserAtGroupID

////////////////////////////////////////////////////////////////////////////////
// authentications

// AuthenticationID 是认证方案(AuthenticationEntity)的实体 ID
type AuthenticationID = authentications.ID

// AuthenticationEntity 表示一个具体的认证方案实例(数据实体)
type AuthenticationEntity = authentications.Entity

// AuthenticationDAO 是针对 AuthenticationEntity 的数据访问接口
type AuthenticationDAO = authentications.DAO

// AuthenticationService 是针对认证方案的服务接口
type AuthenticationService = authentications.Service

// authx

// Authentication 表示一次身份验证请求
type Authentication = authx.Authentication

// Authenticator 是身份验证器的接口, 用于实现具体的验证机制
type Authenticator = authx.Authenticator

// Authorization 表示一次授权请求
type Authorization = authx.Authorization

// Authorizer 是授权器的接口, 用于实现具体的授权检查
type Authorizer = authx.Authorizer

// [注意] 应该使用 authx.Registry 作为注入 api
// AuthxRegistry 是验证器与授权器的注册表接口(用于依赖注入)
type AuthxRegistry = authx.Registry

// AuthxRegistration 是验证器与授权器共用的注册信息结构
type AuthxRegistration = authx.Registration

////////////////////////////////////////////////////////////////////////////////
// checkers

// CheckerService 是针对实体的业务检查服务接口
type CheckerService = checkers.Service

// CheckerChain 是检查器责任链接口, 检查请求会依次经过链上的各个检查器
type CheckerChain = checkers.CheckerChain

// CheckerRegistry 是检查器的注册接口, 用于向检查链注册检查器
type CheckerRegistry = checkers.CheckerRegistry

// Checking 表示一次业务检查的请求上下文, 包含要检查的实体列表与 DTO 列表
type Checking = checkers.Checking

////////////////////////////////////////////////////////////////////////////////
// domain (name)

type DomainName = dxo.DomainName

////////////////////////////////////////////////////////////////////////////////
// email

// EmailAddress 表示 'user@domain' 形式的邮件地址
type EmailAddress = dxo.EmailAddress

// EmailAddressID 是邮件地址记录(EmailAddressEntity)的实体 ID
type EmailAddressID = dxo.EmailAddressID

// EmailAddressDTO 表示邮件地址记录的 REST 网络对象
type EmailAddressDTO = emailaddresses.EmailAddressDTO

// EmailAddressEntity 是邮件地址记录(EmailAddress)的数据实体
type EmailAddressEntity = emailaddresses.EmailAddressEntity

////////////////////////////////////////////////////////////////////////////////
// groups

// GroupID 是资源组(Group)的实体 ID
type GroupID = groups.GroupID

// GroupVO 是资源组列表的视图对象
type GroupVO = groups.GroupVO

// GroupDTO 表示 Group 的 REST 网络对象
type GroupDTO = groups.GroupDTO

// GroupEntity 是资源组(Group)的数据实体
type GroupEntity = groups.GroupEntity

// GroupDAO 是针对 GroupEntity 的数据访问接口
type GroupDAO = groups.DAO

// GroupService 是针对 GroupDTO 的服务接口
type GroupService = groups.Service

////////////////////////////////////////////////////////////////////////////////
// page

// Pagination 是通用的分页参数
type Pagination = dxo.Pagination

////////////////////////////////////////////////////////////////////////////////
// locations

// URI 是统一资源标识符(Uniform Resource Identifier)
type URI = lang.URI

// URL 是统一资源定位符(Uniform Resource Locator)
type URL = lang.URL

////////////////////////////////////////////////////////////////////////////////
// permissions

// PermissionID 是权限(Permission)的实体 ID
type PermissionID = permissions.ID

// PermissionEntity 是权限(Permission)的数据实体, 用于描述对 Method+Path 的访问许可
type PermissionEntity = permissions.Entity

// PermissionDTO 表示 Permission 的 REST 网络对象
type PermissionDTO = permissions.DTO

// PermissionVO 是权限列表的视图对象
type PermissionVO = permissions.VO

// PermissionDAO 是针对 PermissionEntity 的数据访问接口
type PermissionDAO = permissions.DAO

// PermissionService 是针对 PermissionDTO 的服务接口
type PermissionService = permissions.Service

////////////////////////////////////////////////////////////////////////////////
// phone

// PhoneNumber 表示电话号码
type PhoneNumber = dxo.PhoneNumber

// FullPhoneNumber 表示完整的电话号码, 例如: "+86-123-4567-8901"
type FullPhoneNumber = dxo.FullPhoneNumber

// SimplePhoneNumber 表示简短的电话号码, 标准化后为纯数字形式, 例如: "12345678901"
type SimplePhoneNumber = dxo.SimplePhoneNumber

// PurePhoneNumber 表示完整且纯粹的电话号码, 例如: "8612345678901"
type PurePhoneNumber = dxo.PurePhoneNumber

// phone-num:

// PhoneNumberDTO 表示电话号码记录的 REST 网络对象
type PhoneNumberDTO = phonenumbers.PhoneNumberDTO

// PhoneNumberVO 是电话号码列表的视图对象
type PhoneNumberVO = phonenumbers.PhoneNumberVO

// PhoneNumberEntity 是电话号码记录(PhoneNumber)的数据实体
type PhoneNumberEntity = phonenumbers.PhoneNumberEntity

// PhoneNumberDAO 是针对 PhoneNumberEntity 的数据访问接口
type PhoneNumberDAO = phonenumbers.PhoneNumberDAO

// PhoneNumberService 是针对 PhoneNumberDTO 的服务接口
type PhoneNumberService = phonenumbers.PhoneNumberService

// PhoneNumberQuery 是电话号码的查询参数
type PhoneNumberQuery = phonenumbers.PhoneNumberQuery

////////////////////////////////////////////////////////////////////////////////
// refs

// EntityRef 是数据实体(Entity)的引用接口, 用于获取 *Entity
type EntityRef = dxo.EntityRef

// DTORef 是数据传输对象(DTO)的引用接口, 用于获取 *DTO
type DTORef = dxo.DTORef

// VORef 是视图对象(VO)的引用接口, 用于获取 *VO
type VORef = dxo.VORef

////////////////////////////////////////////////////////////////////////////////
// regions

// RegionID 是地区(Region)的实体 ID
type RegionID = regions.ID

// RegionDTO 表示地区的 REST 网络对象
type RegionDTO = regions.DTO

// RegionVO 是地区列表的视图对象
type RegionVO = regions.VO

// RegionEntity 是地区(Region)的数据实体
type RegionEntity = regions.Entity

// RegionDAO 是针对 RegionEntity 的数据访问接口
type RegionDAO = regions.DAO

// RegionService 是针对地区 DTO 的服务接口
type RegionService = regions.Service

////////////////////////////////////////////////////////////////////////////////
// role

// RoleID 是角色(Role)的实体 ID
type RoleID = roles.ID

// RoleName 是角色的正式名称
type RoleName = roles.Name

// RoleNameList 是一组以逗号分隔的角色名称列表
type RoleNameList = dxo.RoleNameList

// RoleEntity 是角色(Role)的数据实体
type RoleEntity = roles.Entity

// RoleDTO 表示 Role 的 REST 网络对象
type RoleDTO = roles.DTO

// RoleVO 是角色列表的视图对象
type RoleVO = roles.VO

// RoleDAO 是针对 RoleEntity 的数据访问接口
type RoleDAO = roles.DAO

// RoleService 是针对 RoleDTO 的服务接口
type RoleService = roles.Service

////////////////////////////////////////////////////////////////////////////////
// session

// SessionDTO 表示会话的 REST 网络对象
type SessionDTO = sessions.DTO

// SessionVO 是会话列表的视图对象
type SessionVO = sessions.VO

// SessionEntity 是会话(Session)的数据实体
type SessionEntity = sessions.Entity

// SessionDAO 是针对 SessionEntity 的数据访问接口
type SessionDAO = sessions.DAO

// SessionService 是针对会话的服务接口
type SessionService = sessions.SessionService

////////////////////////////////////////////////////////////////////////////////
// tables

// TableID 是数据表元信息(Table)的实体 ID
type TableID = tables.ID

// TableName 是 entity.TableName() 方法的返回值
type TableName = tables.Name

// TableDTO 表示数据表元信息的 REST 网络对象
type TableDTO = tables.DTO

// TableVO 是数据表元信息列表的视图对象
type TableVO = tables.VO

// TableEntity 是数据表元信息(Table)的数据实体
type TableEntity = tables.Entity

// TableDAO 是针对 TableEntity 的数据访问接口
type TableDAO = tables.DAO

// TableService 是针对数据表元信息的服务接口
type TableService = tables.Service

////////////////////////////////////////////////////////////////////////////////
// tokens

// TokenDTO 表示令牌的 REST 网络对象
type TokenDTO = tokens.DTO

// TokenVO 是令牌的视图对象
type TokenVO = tokens.VO

////////////////////////////////////////////////////////////////////////////////
// user

// UserID 是用户(User)的实体 ID
type UserID = users.UserID

// UserName 表示用户名
type UserName = users.UserName

// UserVO 是用户列表的视图对象
type UserVO = users.UserVO

// UserDTO 表示 User 的 REST 网络对象
type UserDTO = users.UserDTO

// UserEntity 是用户(User)的数据实体
type UserEntity = users.UserEntity

// UserInfo 用于在各个对象之间交换用户信息
type UserInfo = dxo.UserInfo

// UserDAO 是针对 UserEntity 的数据访问接口
type UserDAO = users.DAO

// UserService 是针对 UserDTO 的服务接口
type UserService = users.Service

////////////////////////////////////////////////////////////////////////////////
// EOF
