package sessions

type Query struct {
	All bool

	Pagination Pagination

	Want *Entity
}
