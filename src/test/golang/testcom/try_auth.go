package testcom

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

type TryAuth struct {

	//starter:component

	_as func(units.Unit) //starter:as(".")

	AuthService rbac.AuthService //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *TryAuth) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:     "try-auth",
		Enabled:  true,
		Do:       inst.runTryAuth,
		Priority: -10,
	}

	list = append(list, u1)
	return list
}

func (inst *TryAuth) runTryAuth(cc context.Context) error {

	a1 := new(rbac.AuthDTO)
	alist := []*rbac.AuthDTO{a1}
	ser := inst.AuthService

	alist2, err := ser.HandleDTO(cc, alist)
	if err != nil {
		return err
	}

	for _, it := range alist2 {
		vlog.Debug("%s", it.Message)
	}

	return nil
}

func (inst *TryAuth) _impl() units.Unit {
	return inst
}
