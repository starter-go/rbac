package authx

import "context"

type Handler interface {
	Handle(c *Context) error
}

type Service interface {
	Handler

	HandleDTO(cc context.Context, items []*AuthDTO) ([]*AuthDTO, error)
}
