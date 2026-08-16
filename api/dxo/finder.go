package dxo

import (
	"fmt"

	"gorm.io/gorm"
)

type Finder struct {
	db   *gorm.DB
	page *Pagination

	all bool

	list  any
	want  any
	model any
}

func (inst *Finder) Find() error {

	db := inst.db
	model := inst.model
	want := inst.want
	page := inst.page
	list := inst.list
	all := inst.all

	if db == nil {
		return inst.innerMakeErr("param 'db' is nil")
	}

	if model != nil {
		db = db.Model(model)
	}

	if all {
		return inst.innerFindAll(db)
	}

	if want != nil {
		db = db.Where(want)
	}

	if page == nil {
		page = new(Pagination)
		page.Limit = 5
		page.Offset = 0
		inst.page = page
	}

	// count
	var count int64
	res := db.Count(&count)
	if res.Error == nil {
		page.Total = count
	}

	// limit
	offset := page.Offset
	limit := page.Limit
	db = db.Limit(limit).Offset(int(offset))

	// find
	db = db.Find(list)
	return db.Error
}

func (inst *Finder) innerFindAll(db *gorm.DB) error {
	list := inst.list
	res := db.Find(list)
	return res.Error
}

func (inst *Finder) innerMakeErr(f string, args ...any) error {
	f = "gorm.DB.Finder: " + f
	return fmt.Errorf(f, args...)
}

func (inst *Finder) SetDB(db *gorm.DB) *Finder {
	inst.db = db
	return inst
}

func (inst *Finder) SetPagination(p *Pagination) *Finder {
	inst.page = p
	return inst
}

func (inst *Finder) SetList(list any) *Finder {
	inst.list = list
	return inst
}

func (inst *Finder) SetWant(w any) *Finder {
	inst.want = w
	return inst
}

func (inst *Finder) SetModel(m any) *Finder {
	inst.model = m
	return inst
}

func (inst *Finder) SetAll(all bool) *Finder {
	inst.all = all
	return inst
}

func (inst *Finder) Reset() *Finder {

	inst.all = false

	inst.db = nil
	inst.page = nil
	inst.list = nil
	inst.model = nil
	inst.want = nil

	return inst
}
