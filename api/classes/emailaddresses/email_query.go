package emailaddresses

// EmailAddress_Query 查询参数
type Query struct {
	All bool // 查询全部条目

	Pagination Pagination

	Want *Entity

	Q string // the query.text
	A []any  // the query.args

}
