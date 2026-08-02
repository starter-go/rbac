package mem

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/lib/classes/sessions"
	"gorm.io/gorm"
)

type MemorySessionDao struct {

	//starter:component

	_as func(ISessionDAO) //starter:as("#")

	Engine IMemoryEngine //starter:inject("#")

}

func (inst *MemorySessionDao) innerMakeItem() *sessions.Entity {
	return new(sessions.Entity)
}

func (inst *MemorySessionDao) innerGenUUID() lang.UUID {
	return inst.Engine.NextUUID()
}

func (inst *MemorySessionDao) innerIsWant(want, have *sessions.Entity) bool {

	if want == nil || have == nil {
		return false
	}

	return true
}

// Delete implements [ISessionDAO].
func (inst *MemorySessionDao) Delete(db *gorm.DB, id sessions.ID) error {

	item := inst.innerMakeItem()
	ls := inst.Engine.NewLS()

	ls.SetIntID(int64(id))
	ls.SetItem(item)

	return ls.DoDelete()

}

// Find implements [ISessionDAO].
func (inst *MemorySessionDao) Find(db *gorm.DB, id sessions.ID) (*sessions.Entity, error) {

	item := inst.innerMakeItem()
	ls := inst.Engine.NewLS()

	ls.SetIntID(int64(id))
	ls.SetItem(item)

	err := ls.DoFind()
	return item, err

}

// GetDB implements [ISessionDAO].
func (inst *MemorySessionDao) GetDB(old *gorm.DB) *gorm.DB {

	return inst.Engine.GetDB(old)

}

// Insert implements [ISessionDAO].
func (inst *MemorySessionDao) Insert(db *gorm.DB, item *sessions.Entity) (*sessions.Entity, error) {

	uuid := inst.innerGenUUID()
	item.UUID = uuid

	ls := inst.Engine.NewLS()
	ls.OnSetIntID(func(id int64) {
		item.ID = sessions.ID(id)
	})
	ls.SetItem(item)

	err := ls.DoInsert()
	return item, err

}

// Query implements [ISessionDAO].
func (inst *MemorySessionDao) Query(db *gorm.DB, q *sessions.Query) ([]*sessions.Entity, error) {

	want := q.Want
	model := inst.innerMakeItem()
	ls := inst.Engine.NewLS()
	page := &q.Pagination
	results := make([]*sessions.Entity, 0)

	mq, err := DoQuery(ls, model)
	if err != nil {
		return nil, err
	}

	mq.NewItem(func() *sessions.Entity {
		return inst.innerMakeItem()

	}).Where(func(item *sessions.Entity) bool {
		return inst.innerIsWant(want, item)

	}).Then(func(item *sessions.Entity) {
		results = append(results, item)
	})

	err = mq.Query(page)
	return results, err

}

// Update implements [ISessionDAO].
func (inst *MemorySessionDao) Update(db *gorm.DB, id sessions.ID, callback func(old *sessions.Entity) error) (*sessions.Entity, error) {

	item := new(sessions.Entity)
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

func (inst *MemorySessionDao) _impl() ISessionDAO {
	return inst
}
