package groups

import "context"

// Group_Service 是针对 GroupDTO 的服务
type Service interface {

	//fetch

	Find(c context.Context, id GroupID) (*GroupDTO, error)

	Query(c context.Context, q *Query) ([]*GroupDTO, error)

	// edit

	Insert(c context.Context, o *GroupDTO) (*GroupDTO, error)

	Update(c context.Context, id GroupID, o *GroupDTO) (*GroupDTO, error)

	Delete(c context.Context, id GroupID) error
}
