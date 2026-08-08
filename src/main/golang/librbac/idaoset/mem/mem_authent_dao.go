package mem

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/api/classes/authentications"
	"gorm.io/gorm"
)

type MemoryAuthenticationDAO struct {

	//starter:component

	_as func(IAuthenticationDAO) //starter:as("#")

	Engine IMemoryEngine //starter:inject("#")

}

func (inst *MemoryAuthenticationDAO) innerMakeItem() *authentications.Entity {
	return new(authentications.Entity)
}

func (inst *MemoryAuthenticationDAO) innerGenUUID() lang.UUID {
	return inst.Engine.NextUUID()
}

func (inst *MemoryAuthenticationDAO) innerIsWant(want, have *authentications.Entity) bool {

	if want == nil || have == nil {
		return false
	}

	return true
}

// Delete implements [IAuthenticationDAO].
func (inst *MemoryAuthenticationDAO) Delete(db *gorm.DB, id authentications.ID) error {

	item := inst.innerMakeItem()
	ls := inst.Engine.NewLS()

	ls.SetIntID(int64(id))
	ls.SetItem(item)

	return ls.DoDelete()

}

// Find implements [IAuthenticationDAO].
func (inst *MemoryAuthenticationDAO) Find(db *gorm.DB, id authentications.ID) (*authentications.Entity, error) {

	item := inst.innerMakeItem()
	ls := inst.Engine.NewLS()

	ls.SetIntID(int64(id))
	ls.SetItem(item)

	err := ls.DoFind()
	return item, err

}

// GetDB implements [IAuthenticationDAO].
func (inst *MemoryAuthenticationDAO) GetDB(old *gorm.DB) *gorm.DB {

	return inst.Engine.GetDB(old)

}

// Insert implements [IAuthenticationDAO].
func (inst *MemoryAuthenticationDAO) Insert(db *gorm.DB, item *authentications.Entity) (*authentications.Entity, error) {

	uuid := inst.innerGenUUID()
	item.UUID = uuid

	ls := inst.Engine.NewLS()
	ls.OnSetIntID(func(id int64) {
		item.ID = authentications.ID(id)
	})
	ls.SetItem(item)

	err := ls.DoInsert()
	return item, err

}

// Query implements [IAuthenticationDAO].
func (inst *MemoryAuthenticationDAO) Query(db *gorm.DB, q *authentications.Query) ([]*authentications.Entity, error) {

	want := q.Want
	model := inst.innerMakeItem()
	ls := inst.Engine.NewLS()
	page := &q.Pagination
	results := make([]*authentications.Entity, 0)

	mq, err := DoQuery(ls, model)
	if err != nil {
		return nil, err
	}

	mq.NewItem(func() *authentications.Entity {
		return inst.innerMakeItem()

	}).Where(func(item *authentications.Entity) bool {
		return inst.innerIsWant(want, item)

	}).Then(func(item *authentications.Entity) {
		results = append(results, item)
	})

	err = mq.Query(page)
	return results, err

}

// Update implements [IAuthenticationDAO].
func (inst *MemoryAuthenticationDAO) Update(db *gorm.DB, id authentications.ID, callback func(old *authentications.Entity) error) (*authentications.Entity, error) {

	item := new(authentications.Entity)
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

func (inst *MemoryAuthenticationDAO) _impl() IAuthenticationDAO {
	return inst
}
