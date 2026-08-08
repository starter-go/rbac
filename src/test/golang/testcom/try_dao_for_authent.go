package testcom

import (
	"context"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/authentications"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

type TryDaoForAuthent struct {

	//starter:component

	_as func(units.Unit) //starter:as(".")

	Dao     rbac.AuthenticationDAO //starter:inject("#")
	AuthSer rbac.AuthService       //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *TryDaoForAuthent) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:     "try-authent-dao",
		Enabled:  true,
		Do:       inst.runTryAuthentDao,
		Priority: 10,
	}

	u2 := &units.Registration{
		Name:     "try-login",
		Enabled:  true,
		Do:       inst.runTryLogin,
		Priority: 0,
	}

	list = append(list, u1, u2)
	return list
}

func (inst *TryDaoForAuthent) runTryLogin(cc context.Context) error {

	ser := inst.AuthSer
	username := "foo"
	password := "123"
	passb64 := lang.Base64FromBytes([]byte(password))

	a1 := &rbac.AuthDTO{
		Mechanism: rbac.MechanismPassword,
		Account:   username,
		Secret:    passb64,
	}
	a2 := &rbac.AuthDTO{
		Action: rbac.ActionLogin,
	}
	a3 := &rbac.AuthDTO{
		Mechanism: rbac.MechanismSMS,
		Action:    rbac.ActionResetPassword,
		Account:   "1234567890",
	}

	list1 := []*rbac.AuthDTO{a1, a2, a3}
	for i, it := range list1 {
		inst.innerLogAuthDTO("want", i, it)
	}

	list2, err := ser.HandleDTO(cc, list1)
	if err != nil {
		return err
	}
	for i, it := range list2 {
		inst.innerLogAuthDTO("have", i, it)
	}

	return nil
}

func (inst *TryDaoForAuthent) innerLogAuthDTO(tag string, index int, it *rbac.AuthDTO) {

	act := it.Action
	mech := it.Mechanism
	msg := it.Message

	vlog.Info("%s [rbac.AuthDTO index:%d action:'%s' mechanism:'%s' msg:'%s' ]", tag, index, act, mech, msg)
}

func (inst *TryDaoForAuthent) runTryAuthentDao(cc context.Context) error {

	dao := inst.Dao
	it1 := new(rbac.AuthenticationEntity)
	db := dao.GetDB(nil)

	it2, err := dao.Insert(db, it1)
	if err != nil {
		return err
	}

	id := it2.ID

	it3, err := dao.Find(db, id)
	if err != nil {
		return err
	}

	it4, err := dao.Update(db, id, func(old *authentications.Entity) error {

		old.CommonName = "foo"

		return nil
	})
	if err != nil {
		return err
	}

	err = dao.Delete(db, id)
	if err != nil {
		return err
	}

	_, err = dao.Find(db, id)
	if err != nil {
		vlog.Warn(" (find after delete) error : %s", err.Error())
	}

	vlog.Debug("insert user[1].entity.id = %d", it1.ID)
	vlog.Debug("insert user[2].entity.id = %d", it2.ID)
	vlog.Debug("insert user[3].entity.id = %d", it3.ID)
	vlog.Debug("insert user[4].entity.id = %d", it4.ID)

	return nil
}

func (inst *TryDaoForAuthent) _impl() units.Unit {
	return inst
}
