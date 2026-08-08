package rbac

import "github.com/starter-go/rbac/api/dxo"

func CopyBaseFieldsE2D(src EntityRef, dst DTORef) error {
	return dxo.CopyBaseFieldsE2D(src, dst)
}

func CopyBaseFieldsD2E(src DTORef, dst EntityRef) error {
	return dxo.CopyBaseFieldsD2E(src, dst)
}
