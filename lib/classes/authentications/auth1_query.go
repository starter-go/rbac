package authentications

// authentications.Query
type Query struct {
	All bool

	Pagination Pagination

	Want *Entity

	Q string // the query.text
	A []any  // the query.args

}
