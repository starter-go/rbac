package rbac

import "github.com/starter-go/rbac/api/dxo"

func ParseFullPhoneNumber(str string) (dxo.FullPhoneNumber, error) {
	return dxo.ParseFullPhoneNumber(str)
}
