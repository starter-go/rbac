package tables

import "github.com/starter-go/rbac/api/dxo"

func ConvertD2E(src *DTO, dst *Entity) error {

	dst.ID = src.ID

	dxo.CopyBaseFieldsD2E(src, dst)

	dst.Name = src.Name
	dst.Label = src.Label
	dst.Description = src.Description

	dst.GroupName = src.GroupName
	dst.GroupURI = src.GroupURI
	dst.TableURI = src.TableURI

	return nil
}

func ConvertE2D(src *Entity, dst *DTO) error {

	dst.ID = src.ID

	dxo.CopyBaseFieldsE2D(src, dst)

	dst.Name = src.Name
	dst.Label = src.Label
	dst.Description = src.Description

	dst.GroupName = src.GroupName
	dst.GroupURI = src.GroupURI
	dst.TableURI = src.TableURI

	return nil
}
