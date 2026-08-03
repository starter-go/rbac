package mem

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/lib/classes/tables"
	"gorm.io/gorm"
)

type MemoryTableDao struct {

	//starter:component

	_as func(ITableDAO) //starter:as("#")

	Engine IMemoryEngine //starter:inject("#")

}

func (inst *MemoryTableDao) innerMakeItem() *tables.Entity {
	return new(tables.Entity)
}

func (inst *MemoryTableDao) innerGenUUID() lang.UUID {
	return inst.Engine.NextUUID()
}

func (inst *MemoryTableDao) innerIsWant(want, have *tables.Entity) bool {

	if want == nil || have == nil {
		return false
	}

	return true
}

// Delete implements [ITableDAO].
func (inst *MemoryTableDao) Delete(db *gorm.DB, id tables.ID) error {

	item := inst.innerMakeItem()
	ls := inst.Engine.NewLS()

	ls.SetIntID(int64(id))
	ls.SetItem(item)

	return ls.DoDelete()

}

// Find implements [ITableDAO].
func (inst *MemoryTableDao) Find(db *gorm.DB, id tables.ID) (*tables.Entity, error) {

	item := inst.innerMakeItem()
	ls := inst.Engine.NewLS()

	ls.SetIntID(int64(id))
	ls.SetItem(item)

	err := ls.DoFind()
	return item, err

}

// GetDB implements [ITableDAO].
func (inst *MemoryTableDao) GetDB(old *gorm.DB) *gorm.DB {

	return inst.Engine.GetDB(old)

}

// Insert implements [ITableDAO].
func (inst *MemoryTableDao) Insert(db *gorm.DB, item *tables.Entity) (*tables.Entity, error) {

	uuid := inst.innerGenUUID()
	item.UUID = uuid

	ls := inst.Engine.NewLS()
	ls.OnSetIntID(func(id int64) {
		item.ID = tables.ID(id)
	})
	ls.SetItem(item)

	err := ls.DoInsert()
	return item, err

}

// Query implements [ITableDAO].
func (inst *MemoryTableDao) Query(db *gorm.DB, q *tables.Query) ([]*tables.Entity, error) {

	want := q.Want
	model := inst.innerMakeItem()
	ls := inst.Engine.NewLS()
	page := &q.Pagination
	results := make([]*tables.Entity, 0)

	mq, err := DoQuery(ls, model)
	if err != nil {
		return nil, err
	}

	mq.NewItem(func() *tables.Entity {
		return inst.innerMakeItem()

	}).Where(func(item *tables.Entity) bool {
		return inst.innerIsWant(want, item)

	}).Then(func(item *tables.Entity) {
		results = append(results, item)
	})

	err = mq.Query(page)
	return results, err

}

// Update implements [ITableDAO].
func (inst *MemoryTableDao) Update(db *gorm.DB, id tables.ID, callback func(old *tables.Entity) error) (*tables.Entity, error) {

	item := new(tables.Entity)
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

func (inst *MemoryTableDao) _impl() ITableDAO {
	return inst
}
