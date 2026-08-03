package core

import (
	"errors"
	"fmt"
	"sort"

	"github.com/starter-go/rbac"
)

type RbacDaoSetLoader struct {

	//starter:component

	_as func(rbac.DaoSetLoader) //starter:as("#")

	DSRegList []rbac.DaoSetRegistry //starter:inject(".")
}

// Load implements [rbac.DaoSetLoader].
func (inst *RbacDaoSetLoader) Load() (*rbac.DaoSet, error) {
	loading := new(innerDSLoading)
	return loading.load(inst.DSRegList)
}

func (inst *RbacDaoSetLoader) _impl() rbac.DaoSetLoader {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerDSLoading struct {
	items []*rbac.DaoSetRegistration
}

// Len implements [sort.Interface].
func (inst *innerDSLoading) Len() int {
	return len(inst.items)
}

// Less implements [sort.Interface].
func (inst *innerDSLoading) Less(i1, i2 int) bool {
	n1 := inst.items[i1].Priority
	n2 := inst.items[i2].Priority
	return (n1 < n2)
}

// Swap implements [sort.Interface].
func (inst *innerDSLoading) Swap(i1, i2 int) {
	l := inst.items
	l[i1], l[i2] = l[i2], l[i1]
}

func (inst *innerDSLoading) load(src []rbac.DaoSetRegistry) (*rbac.DaoSet, error) {

	inst.items = nil

	for _, reg1 := range src {
		reg2 := reg1.Registration()
		inst.add(reg2)
	}

	inst.sort()
	all := inst.items
	tmp := new(rbac.DaoSet)
	dst := new(rbac.DaoSet)

	for _, it := range all {
		tmp2 := it.Provider.Provide(tmp)
		inst.copyWithoutItemsNil(tmp2, dst)
	}

	err := inst.checkDS(dst)
	return dst, err
}

func (inst *innerDSLoading) copyWithoutItemsNil(src, dst *rbac.DaoSet) {

	if src.Authentications != nil {
		dst.Authentications = src.Authentications
	}

	if src.Permissions != nil {
		dst.Permissions = src.Permissions
	}

	if src.Roles != nil {
		dst.Roles = src.Roles
	}

	if src.Sessions != nil {
		dst.Sessions = src.Sessions
	}

	if src.Tables != nil {
		dst.Tables = src.Tables
	}

	if src.Users != nil {
		dst.Users = src.Users
	}

}

func (inst *innerDSLoading) checkDS(ds *rbac.DaoSet) error {

	all := make(map[string]any)
	errlist := make([]error, 0)

	all["authent_dao"] = ds.Authentications
	all["perm_dao"] = ds.Permissions
	all["role_dao"] = ds.Roles
	all["session_dao"] = ds.Sessions
	all["table_dao"] = ds.Tables
	all["user_dao"] = ds.Users

	for name, ptr := range all {
		if ptr == nil {
			e2 := fmt.Errorf("rbac.DaoSetLoader: DAO '%s' is nil", name)
			errlist = append(errlist, e2)
		}
	}

	return errors.Join(errlist...)
}

func (inst *innerDSLoading) accept(item *rbac.DaoSetRegistration) bool {

	if item == nil {
		return false
	}

	if item.Provider == nil {
		return false
	}

	if !item.Enabled {
		return false
	}

	return true
}

func (inst *innerDSLoading) add(item *rbac.DaoSetRegistration) {
	if !inst.accept(item) {
		return
	}
	inst.items = append(inst.items, item)
}

func (inst *innerDSLoading) sort() {
	sort.Sort(inst)
}
