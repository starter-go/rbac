package rbac

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/api/classes/tables"
	"github.com/starter-go/rbac/api/classes/users"
	"github.com/starter-go/rbac/api/localization"
)

type EntityUpdater struct {
}

func (inst *EntityUpdater) UpdateInt(src, dst *int) {
	if src == nil || dst == nil {
		return
	}

	const empty = 0
	value := *src

	if value != empty {
		*dst = value
	}

}

func (inst *EntityUpdater) UpdateString(src, dst *string) {
	if src == nil || dst == nil {
		return
	}

	const empty = ""
	value := *src

	if value != empty {
		*dst = value
	}

}

func (inst *EntityUpdater) UpdateURL(src, dst *lang.URL) {
	if src == nil || dst == nil {
		return
	}

	const empty = ""
	value := *src

	if value != empty {
		*dst = value
	}

}

func (inst *EntityUpdater) UpdateURI(src, dst *lang.URI) {

	if src == nil || dst == nil {
		return
	}
	const empty = ""
	value := *src
	if value != empty {
		*dst = value
	}

}

func (inst *EntityUpdater) UpdateUserName(src, dst *users.Name) {
	if src == nil || dst == nil {
		return
	}

	const empty = ""
	value := *src

	if value != empty {
		*dst = value
	}

}

func (inst *EntityUpdater) UpdateTableName(src, dst *tables.Name) {
	if src == nil || dst == nil {
		return
	}

	const empty = ""
	value := *src

	if value != empty {
		*dst = value
	}

}

func (inst *EntityUpdater) UpdateBool(src, dst *bool) {
	if src == nil || dst == nil {
		return
	}
	value := *src
	*dst = value
}

func (inst *EntityUpdater) UpdateBase64(src, dst *lang.Base64) {
	if src == nil || dst == nil {
		return
	}

	const empty = ""
	value := *src

	if value != empty {
		*dst = value
	}

}

func (inst *EntityUpdater) UpdateHex(src, dst *lang.Hex) {
	if src == nil || dst == nil {
		return
	}

	const empty = ""
	value := *src

	if value != empty {
		*dst = value
	}

}

func (inst *EntityUpdater) UpdatePhone(src, dst *PhoneNumber) {
	if src == nil || dst == nil {
		return
	}

	const empty = ""
	value := *src

	if value != empty {
		*dst = value
	}

}

func (inst *EntityUpdater) UpdateEmail(src, dst *EmailAddress) {
	if src == nil || dst == nil {
		return
	}

	const empty = ""
	value := *src

	if value != empty {
		*dst = value
	}

}

func (inst *EntityUpdater) UpdateLocale(src, dst *localization.Locale) {
	if src == nil || dst == nil {
		return
	}

	const empty = ""
	value := *src

	if value != empty {
		*dst = value
	}

}

func (inst *EntityUpdater) UpdateRoles(src, dst *RoleNameList) {

	if src == nil || dst == nil {
		return
	}

	const empty = ""
	value := *src

	if value != empty {
		*dst = value
	}

}
