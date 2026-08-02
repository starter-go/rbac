package regions

import "github.com/starter-go/rbac/lib/dxo"

type ID = dxo.RegionID

////////////////////////////////////////////////////////////////////////////////

type DTO = RegionDTO

type VO = RegionVO

type Entity = RegionEntity

type Pagination = dxo.Pagination

////////////////////////////////////////////////////////////////////////////////

type RegionDTO struct {
	ID ID

	dxo.BaseDTO
}

type RegionVO struct {
	dxo.BaseVO

	Items []*DTO `json:"regions"`
}

type RegionEntity struct {
	ID ID

	dxo.BaseEntity
}

////////////////////////////////////////////////////////////////////////////////
