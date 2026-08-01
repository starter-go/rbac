package permissions

// Permission_Query 查询参数
type Query struct {
	All bool // 查询全部条目

	Pagination Pagination

	Want *Entity
}
