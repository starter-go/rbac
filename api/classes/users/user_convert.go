package users

import (
	"fmt"

	"github.com/starter-go/rbac/api/dxo"
)

////////////////////////////////////////////////////////////////////////////////

func ConvertD2E(src *DTO, dst *Entity) error {

	dst.ID = src.ID

	dxo.CopyBaseFieldsD2E(src, dst)

	dst.Name = src.Name
	dst.Email = src.Email
	dst.Phone = src.Phone

	dst.Avatar = src.Avatar
	dst.Nickname = src.Nickname
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
	dst.Nickname = src.Nickname
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

////////////////////////////////////////////////////////////////////////////////

func ConvertE2I(src *Entity, dst *dxo.UserInfo) error {

	if src == nil || dst == nil {
		return fmt.Errorf("param(s) is nil")
	}

	dst.Avatar = src.Avatar
	dst.Nickname = src.Nickname
	dst.Email = src.Email
	dst.Language = src.Language
	dst.Roles = src.Roles

	dst.Enabled = src.Enabled
	dst.Locked = src.Locked
	dst.Use2FA = src.Use2FA

	dst.Username = src.Name
	dst.UserID = src.ID
	dst.UserUUID = src.UUID

	return nil
}

func ConvertI2E(src *dxo.UserInfo, dst *Entity) error {

	if src == nil || dst == nil {
		return fmt.Errorf("param(s) is nil")
	}

	dst.Avatar = src.Avatar
	dst.Nickname = src.Nickname
	dst.Email = src.Email
	dst.Language = src.Language
	dst.Roles = src.Roles

	dst.Enabled = src.Enabled
	dst.Locked = src.Locked
	dst.Use2FA = src.Use2FA

	dst.Name = src.Username
	dst.ID = src.UserID
	dst.UUID = src.UserUUID

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
