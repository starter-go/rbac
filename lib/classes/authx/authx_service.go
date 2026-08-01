package authx

type Service interface {
	Handle(c *Context) error
}
