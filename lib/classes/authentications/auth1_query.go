package authentications

type Query struct {
	All bool

	Pagination Pagination

	Want *Entity
}
