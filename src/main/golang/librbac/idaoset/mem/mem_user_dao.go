package mem

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/lib/classes/users"
	"gorm.io/gorm"
)

type MemoryUserDao struct {

	//starter:component

	_as func(IUserDao) //starter:as("#")

	Engine IMemoryEngine //starter:inject("#")

}

func (inst *MemoryUserDao) innerMakeItem() *users.Entity {
	return new(users.Entity)
}

func (inst *MemoryUserDao) innerGenUUID() lang.UUID {
	return inst.Engine.NextUUID()
}

// Delete implements [users.UserDAO].
func (inst *MemoryUserDao) Delete(db *gorm.DB, id users.ID) error {

	item := inst.innerMakeItem()
	ls := inst.Engine.NewLS()

	ls.SetIntID(int64(id))
	ls.SetItem(item)

	return ls.DoDelete()
}

// Find implements [users.UserDAO].
func (inst *MemoryUserDao) Find(db *gorm.DB, id users.ID) (*users.Entity, error) {

	item := inst.innerMakeItem()
	ls := inst.Engine.NewLS()

	ls.SetIntID(int64(id))
	ls.SetItem(item)

	err := ls.DoFind()
	return item, err
}

// GetDB implements [users.UserDAO].
func (inst *MemoryUserDao) GetDB(old *gorm.DB) *gorm.DB {
	return inst.Engine.GetDB(old)
}

// Insert implements [users.UserDAO].
func (inst *MemoryUserDao) Insert(db *gorm.DB, item *users.Entity) (*users.Entity, error) {

	uuid := inst.innerGenUUID()
	item.UUID = uuid

	ls := inst.Engine.NewLS()
	ls.OnSetIntID(func(id int64) {
		item.ID = users.UserID(id)
	})
	ls.SetItem(item)

	err := ls.DoInsert()
	return item, err
}

// Query implements [users.UserDAO].
func (inst *MemoryUserDao) Query(db *gorm.DB, q *users.Query) ([]*users.Entity, error) {

	want := q.Want
	model := inst.innerMakeItem()
	ls := inst.Engine.NewLS()
	page := &q.Pagination
	results := make([]*users.Entity, 0)

	mq, err := DoQuery(ls, model)
	if err != nil {
		return nil, err
	}

	mq.NewItem(func() *users.Entity {
		return inst.innerMakeItem()

	}).Where(func(item *users.Entity) bool {
		return inst.innerIsWant(want, item)

	}).Then(func(item *users.Entity) {
		results = append(results, item)
	})

	err = mq.Query(page)
	return results, err
}

func (inst *MemoryUserDao) innerIsWant(want, have *users.Entity) bool {

	if want == nil || have == nil {
		return false
	}

	return true
}

// Update implements [users.UserDAO].
func (inst *MemoryUserDao) Update(db *gorm.DB, id users.ID, callback func(old *users.Entity) error) (*users.Entity, error) {

	item := new(users.Entity)
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

func (inst *MemoryUserDao) _impl() rbac.UserDAO {
	return inst
}
