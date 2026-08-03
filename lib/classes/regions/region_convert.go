package regions

import "github.com/starter-go/rbac/lib/dxo"

func ConvertD2E(src *DTO, dst *Entity) error {

	dst.ID = src.ID

	dxo.CopyBaseFieldsD2E(src, dst)

	dst.Code2 = src.Code2
	dst.Code3 = src.Code3
	dst.PhoneCode = src.PhoneCode
	dst.DisplayName = src.DisplayName
	dst.FullName = src.FullName
	dst.SimpleName = src.SimpleName
	dst.FlagURL = src.FlagURL

	return nil

}

func ConvertE2D(src *Entity, dst *DTO) error {

	dst.ID = src.ID

	dxo.CopyBaseFieldsE2D(src, dst)

	dst.Code2 = src.Code2
	dst.Code3 = src.Code3
	dst.PhoneCode = src.PhoneCode
	dst.DisplayName = src.DisplayName
	dst.FullName = src.FullName
	dst.SimpleName = src.SimpleName
	dst.FlagURL = src.FlagURL

	return nil
}
