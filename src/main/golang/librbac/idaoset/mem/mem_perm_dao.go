package mem

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/lib/classes/permissions"
	"gorm.io/gorm"
)

type MemoryPermissionDao struct {

	//starter:component

	_as func(IPermissionDAO) //starter:as("#")

	Engine IMemoryEngine //starter:inject("#")

}

func (inst *MemoryPermissionDao) innerMakeItem() *permissions.Entity {
	return new(permissions.Entity)
}

func (inst *MemoryPermissionDao) innerGenUUID() lang.UUID {
	return inst.Engine.NextUUID()
}

func (inst *MemoryPermissionDao) innerIsWant(want, have *permissions.Entity) bool {

	if want == nil || have == nil {
		return false
	}

	return true
}

// Delete implements [permissions.DAO].
func (inst *MemoryPermissionDao) Delete(db *gorm.DB, id permissions.ID) error {

	item := inst.innerMakeItem()
	ls := inst.Engine.NewLS()

	ls.SetIntID(int64(id))
	ls.SetItem(item)

	return ls.DoDelete()

}

// Find implements [permissions.DAO].
func (inst *MemoryPermissionDao) Find(db *gorm.DB, id permissions.ID) (*permissions.Entity, error) {

	item := inst.innerMakeItem()
	ls := inst.Engine.NewLS()

	ls.SetIntID(int64(id))
	ls.SetItem(item)

	err := ls.DoFind()
	return item, err

}

// GetDB implements [permissions.DAO].
func (inst *MemoryPermissionDao) GetDB(old *gorm.DB) *gorm.DB {

	return inst.Engine.GetDB(old)
}

// Insert implements [permissions.DAO].
func (inst *MemoryPermissionDao) Insert(db *gorm.DB, item *permissions.Entity) (*permissions.Entity, error) {

	uuid := inst.innerGenUUID()
	item.UUID = uuid

	ls := inst.Engine.NewLS()
	ls.OnSetIntID(func(id int64) {
		item.ID = permissions.ID(id)
	})
	ls.SetItem(item)

	err := ls.DoInsert()
	return item, err

}

// Query implements [permissions.DAO].
func (inst *MemoryPermissionDao) Query(db *gorm.DB, q *permissions.Query) ([]*permissions.Entity, error) {

	want := q.Want
	model := inst.innerMakeItem()
	ls := inst.Engine.NewLS()
	page := &q.Pagination
	results := make([]*permissions.Entity, 0)

	mq, err := DoQuery(ls, model)
	if err != nil {
		return nil, err
	}

	mq.NewItem(func() *permissions.Entity {
		return inst.innerMakeItem()

	}).Where(func(item *permissions.Entity) bool {
		return inst.innerIsWant(want, item)

	}).Then(func(item *permissions.Entity) {
		results = append(results, item)
	})

	err = mq.Query(page)
	return results, err

}

// Update implements [permissions.DAO].
func (inst *MemoryPermissionDao) Update(db *gorm.DB, id permissions.ID, callback func(old *permissions.Entity) error) (*permissions.Entity, error) {

	item := new(permissions.Entity)
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

func (inst *MemoryPermissionDao) _impl() rbac.PermissionDAO {
	return inst
}
