package rbac

import "context"

// AvatarID 是 Avatar 的实体 ID
type AvatarID int64

// AvatarDTO 表示一个头像资源
type AvatarDTO struct {
	ID AvatarID `json:"id"`

	BaseDTO

	Label string `json:"label"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

// AvatarQuery ...
type AvatarQuery struct {
	Conditions Conditions
	Pagination Pagination
	All        bool // 查询全部条目
}

// AvatarService ...
type AvatarService interface {
	Find(c context.Context, id AvatarID) (*AvatarDTO, error)
	List(c context.Context, q *AvatarQuery) ([]*AvatarDTO, error)
	ListAll(c context.Context) ([]*AvatarDTO, error)
}
