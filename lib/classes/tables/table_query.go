package tables

type TableQuery struct {
	All bool

	Pagination Pagination

	Want *TableEntity
}
