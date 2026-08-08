package roles

import "github.com/starter-go/rbac/api/dxo"

// Role_Query 查询参数
type Query struct {
	// Conditions Conditions
	Pagination dxo.Pagination
	All        bool // 查询全部条目
	Want       *Entity
}
