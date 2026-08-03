package users

// User_Query 是 User 的查询参数
type UserQuery struct {
	All bool // 查询全部条目

	Pagination Pagination

	Want *UserEntity

	Q string // the query.text
	A []any  // the query.args
}
