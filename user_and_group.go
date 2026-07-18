package rbac

////////////////////////////////////////////////////////////////////////////////
// UserGroup 表示一个用户组

type UserGroupName string

type UserGroupDTO struct {
	ID UserGroupID

	DTO

	Name        UserGroupName `json:"name"`
	Roles       RoleNameList  `json:"roles"`
	Enabled     bool          `json:"enabled"`
	Label       string        `json:"label"`
	Description string        `json:"description"`
}

type UserGroupEntity struct {
	ID UserGroupID

	Entity

	Name        UserGroupName `gorm:"unique"`
	Roles       RoleNameList
	Enabled     bool
	Label       string
	Description string
}

////////////////////////////////////////////////////////////////////////////////
// UserAtGroup 表示用户组中的一个用户, 也就是 (User <---> UserGroup) 之间的(多-多)关系

// 表示 (user<--->group) relation-ship
type UserAtGroupDTO struct {
	ID UserAtGroupID

	DTO

	UID UserID
	GID UserGroupID
}

// 表示 (user<--->group) relation-ship
type UserAtGroupEntity struct {
	ID UserAtGroupID

	Entity

	URI URI `gorm:"unique"`

	UID UserID
	GID UserGroupID
}

////////////////////////////////////////////////////////////////////////////////
// EOF
