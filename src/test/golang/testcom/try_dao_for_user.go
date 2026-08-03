package testcom

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/lib/classes/users"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

type TryDaoForUser struct {

	//starter:component

	_as func(units.Unit) //starter:as(".")

	UserDao rbac.UserDAO     //starter:inject("#")
	UserSer rbac.UserService //starter:inject("#")

}

// ListRegistrations implements units.Unit.
func (inst *TryDaoForUser) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:    "try-user-dao",
		Enabled: true,
		Do:      inst.runTryUserDao,
	}

	u2 := &units.Registration{
		Name:    "try-user-service",
		Enabled: true,
		Do:      inst.runTryUserService,
	}

	list = append(list, u1, u2)
	return list
}

func (inst *TryDaoForUser) runTryUserService(cc context.Context) error {

	ser := inst.UserSer
	it1 := new(rbac.UserDTO)

	it1.Avatar = "https://x.y.z/abc"
	it1.Email = "foo@bar"
	it1.Phone = "1234567890"
	it1.Name = "foo"
	it1.NickName = "Fu"
	it1.Language = "zh_cn"
	it1.Roles = "user,sudo"

	it1.Enabled = true
	it1.Locked = false
	it1.Use2FA = false

	// insert

	it2, err := ser.Insert(cc, it1)
	if err != nil {
		return err
	}

	it1.Enabled = false
	it1.Locked = true
	it1.Use2FA = true

	it3, err := ser.Insert(cc, it1)
	if err != nil {
		return err
	}

	id := it2.ID

	// find

	it4, err := ser.Find(cc, id)
	if err != nil {
		return err
	}

	// query

	q := new(users.Query)
	q.Want = new(users.Entity)
	q.Pagination.Limit = 10

	q.Want.Email = "foo@bar"

	list, err := ser.List(cc, q)
	if err != nil {
		return err
	}

	for index, it := range list {
		inst.logDTO(index, it)
	}

	// update

	it4.Language = "en_us"
	it4.Roles = "root,admin"
	it4.Avatar = "https://example.com/image/avatar.png"
	it4.Phone = ""
	it4.Email = ""
	it4.Name = ""

	it5, err := ser.Update(cc, it4.ID, it4)
	if err != nil {
		return err
	}

	// delete

	err = ser.Delete(cc, id)
	if err != nil {
		return err
	}

	// find (2)

	it6, err := ser.Find(cc, id)
	if err != nil {
		vlog.Warn("find_after_delete: %s", err.Error())
		// return err
	}

	all := []*users.DTO{it1, it2, it3, it4, it5, it6}
	for index, it := range all {
		inst.logDTO(index, it)
	}

	return nil
}

func (inst *TryDaoForUser) logDTO(index int, it *users.DTO) {
	if it == nil {
		it = new(users.DTO)
	}
	id := it.ID
	uuid := it.UUID
	name := it.Name
	phone := it.Phone
	email := it.Email
	const f = "[users.DTO index:%d id:%d uuid:'%s' name:'%s' phone:'%s' email:'%s' ]"
	vlog.Info(f, index, id, uuid, name, phone, email)
}

func (inst *TryDaoForUser) runTryUserDao(cc context.Context) error {

	dao := inst.UserDao
	it1 := new(rbac.UserEntity)
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

	it4, err := dao.Update(db, id, func(old *users.Entity) error {

		old.Avatar = "http://"
		old.NickName = "foo"

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

func (inst *TryDaoForUser) _impl() units.Unit {
	return inst
}
