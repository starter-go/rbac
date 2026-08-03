package emailaddresses

import "github.com/starter-go/rbac/lib/dxo"

func ConvertD2E(src *DTO, dst *EmailAddressEntity) error {

	dst.ID = src.ID

	dxo.CopyBaseFieldsD2E(src, dst)

	dst.Address = src.Address

	return nil
}

func ConvertE2D(src *EmailAddressEntity, dst *DTO) error {

	dst.ID = src.ID

	dxo.CopyBaseFieldsE2D(src, dst)

	dst.Address = src.Address

	return nil
}
