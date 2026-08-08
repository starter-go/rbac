package regions

import "context"

// Service ...
type Service interface {

	// edit

	Insert(c context.Context, o *DTO) (*DTO, error)
	Update(c context.Context, id ID, o *DTO) (*DTO, error)
	Delete(c context.Context, id ID) error

	// fetch

	Find(c context.Context, id ID) (*DTO, error)
	List(c context.Context, q *Query) ([]*DTO, error)
}
