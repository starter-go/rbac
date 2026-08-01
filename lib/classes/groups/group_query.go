package groups

// Group_Query 是 Group 的查询参数
type Query struct {
	All bool // 查询全部条目

	Pagination Pagination

	Want *GroupEntity
}
