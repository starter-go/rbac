package roles

import "context"

// Role_Service 是针对 RoleDTO 的服务
type Service interface {
	Insert(c context.Context, o *DTO) (*DTO, error)
	Update(c context.Context, id ID, o *DTO) (*DTO, error)
	Delete(c context.Context, id ID) error

	Find(c context.Context, id ID) (*DTO, error)
	List(c context.Context, q *Query) ([]*DTO, error)
}
