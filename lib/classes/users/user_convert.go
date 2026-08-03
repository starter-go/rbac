package users

import "github.com/starter-go/rbac/lib/dxo"

func ConvertD2E(src *DTO, dst *Entity) error {

	dst.ID = src.ID

	dxo.CopyBaseFieldsD2E(src, dst)

	dst.Name = src.Name
	dst.Email = src.Email
	dst.Phone = src.Phone

	dst.Avatar = src.Avatar
	dst.NickName = src.NickName
	dst.Language = src.Language
	dst.Roles = src.Roles

	dst.Enabled = src.Enabled
	dst.Locked = src.Locked
	dst.Use2FA = src.Use2FA

	return nil
}

func ConvertE2D(src *Entity, dst *DTO) error {

	dst.ID = src.ID

	dxo.CopyBaseFieldsE2D(src, dst)

	dst.Name = src.Name
	dst.Email = src.Email
	dst.Phone = src.Phone

	dst.Avatar = src.Avatar
	dst.NickName = src.NickName
	dst.Language = src.Language
	dst.Roles = src.Roles

	dst.Enabled = src.Enabled
	dst.Locked = src.Locked
	dst.Use2FA = src.Use2FA

	return nil
}

////////////////////////////////////////////////////////////////////////////////

func ConvertListE2D(src []*Entity, dst []*DTO) ([]*DTO, error) {

	if dst == nil {
		dst = make([]*DTO, 0)
	}

	for _, it1 := range src {
		it2 := new(DTO)
		err := ConvertE2D(it1, it2)
		if err != nil {
			return nil, err
		}
		dst = append(dst, it2)
	}

	return dst, nil
}
