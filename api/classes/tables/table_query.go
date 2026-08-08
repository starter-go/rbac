package tables

type TableQuery struct {
	All bool

	Pagination Pagination

	Want *TableEntity

	Q string // the query.text
	A []any  // the query.args

}
