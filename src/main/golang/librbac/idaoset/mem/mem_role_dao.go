package mem

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/lib/classes/roles"
	"gorm.io/gorm"
)

type MemoryRoleDao struct {

	//starter:component

	_as func(IRoleDAO) //starter:as("#")

	Engine IMemoryEngine //starter:inject("#")

}

func (inst *MemoryRoleDao) innerMakeItem() *roles.Entity {
	return new(roles.Entity)
}

func (inst *MemoryRoleDao) innerGenUUID() lang.UUID {
	return inst.Engine.NextUUID()
}

func (inst *MemoryRoleDao) innerIsWant(want, have *roles.Entity) bool {

	if want == nil || have == nil {
		return false
	}

	return true
}

// Delete implements [roles.DAO].
func (inst *MemoryRoleDao) Delete(db *gorm.DB, id roles.ID) error {

	item := inst.innerMakeItem()
	ls := inst.Engine.NewLS()

	ls.SetIntID(int64(id))
	ls.SetItem(item)

	return ls.DoDelete()

}

// Find implements [roles.DAO].
func (inst *MemoryRoleDao) Find(db *gorm.DB, id roles.ID) (*roles.Entity, error) {

	item := inst.innerMakeItem()
	ls := inst.Engine.NewLS()

	ls.SetIntID(int64(id))
	ls.SetItem(item)

	err := ls.DoFind()
	return item, err
}

// GetDB implements [roles.DAO].
func (inst *MemoryRoleDao) GetDB(old *gorm.DB) *gorm.DB {
	return inst.Engine.GetDB(old)
}

// Insert implements [roles.DAO].
func (inst *MemoryRoleDao) Insert(db *gorm.DB, item *roles.Entity) (*roles.Entity, error) {

	uuid := inst.innerGenUUID()
	item.UUID = uuid

	ls := inst.Engine.NewLS()
	ls.OnSetIntID(func(id int64) {
		item.ID = roles.ID(id)
	})
	ls.SetItem(item)

	err := ls.DoInsert()
	return item, err
}

// Query implements [roles.DAO].
func (inst *MemoryRoleDao) Query(db *gorm.DB, q *roles.Query) ([]*roles.Entity, error) {

	want := q.Want
	model := inst.innerMakeItem()
	ls := inst.Engine.NewLS()
	page := &q.Pagination
	results := make([]*roles.Entity, 0)

	mq, err := DoQuery(ls, model)
	if err != nil {
		return nil, err
	}

	mq.NewItem(func() *roles.Entity {
		return inst.innerMakeItem()

	}).Where(func(item *roles.Entity) bool {
		return inst.innerIsWant(want, item)

	}).Then(func(item *roles.Entity) {
		results = append(results, item)
	})

	err = mq.Query(page)
	return results, err

}

// Update implements [roles.DAO].
func (inst *MemoryRoleDao) Update(db *gorm.DB, id roles.ID, callback func(old *roles.Entity) error) (*roles.Entity, error) {

	item := new(roles.Entity)
	ls := inst.Engine.NewLS()

	ls.SetIntID(int64(id))
	ls.SetItem(item)

	err := ls.DoFind()
	if err != nil {
		return nil, err
	}

	err = callback(item)
	if err != nil {
		return nil, err
	}

	err = ls.DoUpdate()
	return item, err

}

func (inst *MemoryRoleDao) _impl() rbac.RoleDAO {
	return inst
}
