package dxo

// Pagination 是通用的分页参数
type Pagination struct {
	Limit  int   `json:"limit"`
	Offset int64 `json:"offset"`
	Total  int64 `json:"total"` // 所有页面的条目总数
}

// Limit 用于 SQL 查询的页面大小
func (inst *Pagination) GetSize() int64 {

	const min = 1

	size := inst.Limit
	if size < min {
		size = min
	}
	return int64(size)
}

// Offset 用于 SQL 查询的条目位置
func (inst *Pagination) GetPage() int64 {

	const min = 1

	offs := inst.Offset
	size := inst.GetSize()

	if offs < 0 {
		offs = 0
	}

	page := (offs / size) + 1
	if page < min {
		page = min
	}
	return page
}
