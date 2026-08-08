package tables

import "context"

type Service interface {

	// fetch

	Find(c context.Context, id TableID) (*TableDTO, error)

	Query(c context.Context, q *TableQuery) ([]*TableDTO, error)

	// modify

	Insert(c context.Context, item *TableDTO) (*TableDTO, error)

	Update(c context.Context, id TableID, item *TableDTO) (*TableDTO, error)

	Delete(c context.Context, id TableID) error
}
