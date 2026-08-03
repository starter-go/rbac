package iservices

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/lib/classes/checkers"
	"github.com/starter-go/rbac/lib/classes/users"
)

type UserServiceImpl struct {

	//starter:component

	_as func(rbac.UserService) //starter:as("#")

	Dao        rbac.UserDAO        //starter:inject("#")
	CheckerSer rbac.CheckerService //starter:inject("#")

}

func (inst *UserServiceImpl) innerCheckEntity(ctx context.Context, action checkers.Action, elist ...*users.Entity) error {

	ch := new(checkers.Checking)
	ser := inst.CheckerSer

	ch.Context = ctx
	ch.Action = action

	for _, it := range elist {
		ch.Entities = append(ch.Entities, it)
	}

	return ser.Check(ch)
}

// Delete implements [users.UserService].
func (inst *UserServiceImpl) Delete(c context.Context, id users.UserID) error {

	db := inst.Dao.GetDB(nil)
	it1, err := inst.Dao.Find(db, id)
	if err != nil {
		return err
	}

	err = inst.innerCheckEntity(c, checkers.ActionDelete, it1)
	if err != nil {
		return err
	}

	return inst.Dao.Delete(db, id)
}

// Find implements [users.UserService].
func (inst *UserServiceImpl) Find(c context.Context, id users.UserID) (*users.UserDTO, error) {

	db := inst.Dao.GetDB(nil)
	it4 := new(rbac.UserDTO)

	it3, err := inst.Dao.Find(db, id)
	if err != nil {
		return nil, err
	}

	err = inst.innerCheckEntity(c, checkers.ActionFetch, it3)
	if err != nil {
		return nil, err
	}

	err = users.ConvertE2D(it3, it4)
	return it4, err
}

// Insert implements [users.UserService].
func (inst *UserServiceImpl) Insert(c context.Context, it1 *users.UserDTO) (*users.UserDTO, error) {

	db := inst.Dao.GetDB(nil)
	it2 := new(rbac.UserEntity)
	it4 := new(rbac.UserDTO)

	err := users.ConvertD2E(it1, it2)
	if err != nil {
		return nil, err
	}

	err = inst.innerCheckEntity(c, checkers.ActionInsert, it2)
	if err != nil {
		return nil, err
	}

	it3, err := inst.Dao.Insert(db, it2)
	if err != nil {
		return nil, err
	}

	err = users.ConvertE2D(it3, it4)
	return it4, err
}

// List implements [users.UserService].
func (inst *UserServiceImpl) List(c context.Context, q *users.UserQuery) ([]*users.UserDTO, error) {

	db := inst.Dao.GetDB(nil)

	list1, err := inst.Dao.Query(db, q)
	if err != nil {
		return nil, err
	}

	err = inst.innerCheckEntity(c, checkers.ActionFetch, list1...)
	if err != nil {
		return nil, err
	}

	return users.ConvertListE2D(list1, nil)
}

// Update implements [users.UserService].
func (inst *UserServiceImpl) Update(c context.Context, id users.UserID, it1 *users.UserDTO) (*users.UserDTO, error) {

	db := inst.Dao.GetDB(nil)
	it2 := new(users.Entity)

	err := users.ConvertD2E(it1, it2)
	if err != nil {
		return nil, err
	}

	it3, err := inst.Dao.Update(db, id, func(old *users.Entity) error {

		err = inst.innerCheckEntity(c, checkers.ActionUpdate, old)
		if err != nil {
			return err
		}

		return inst.innerUpdateItem(it2, old)
	})
	if err != nil {
		return nil, err
	}

	it4 := new(users.DTO)
	err = users.ConvertE2D(it3, it4)
	return it4, err
}

func (inst *UserServiceImpl) innerUpdateItem(src, dst *users.Entity) error {

	up := new(rbac.EntityUpdater)

	up.UpdateURL(&src.Avatar, &dst.Avatar)
	up.UpdateString(&src.NickName, &dst.NickName)
	up.UpdateLocale(&src.Language, &dst.Language)

	up.UpdateUserName(&src.Name, &dst.Name)
	up.UpdateEmail(&src.Email, &dst.Email)
	up.UpdatePhone(&src.Phone, &dst.Phone)
	up.UpdateRoles(&src.Roles, &dst.Roles)

	up.UpdateBool(&src.Enabled, &dst.Enabled)
	up.UpdateBool(&src.Locked, &dst.Locked)
	up.UpdateBool(&src.Use2FA, &dst.Use2FA)

	return nil
}

func (inst *UserServiceImpl) _impl() rbac.UserService {
	return inst
}
