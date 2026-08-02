package mem

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/lib/classes/tables"
	"github.com/starter-go/rbac/lib/dxo"
	"gorm.io/gorm"
)

type innerTableName tables.Name
type innerRowID string

////////////////////////////////////////////////////////////////////////////////

type innerTableNamer interface {
	TableName() string
}

////////////////////////////////////////////////////////////////////////////////

type MemoryQuery[T any] struct {
	ls    *MemoryLS
	table *innerTable

	handlerNewItem func() T
	handlerThen    func(item T)
	handlerWhere   func(item T) bool
}

func (inst *MemoryQuery[T]) NewItem(fn func() T) *MemoryQuery[T] {
	inst.handlerNewItem = fn
	return inst
}

func (inst *MemoryQuery[T]) Where(fn func(item T) bool) *MemoryQuery[T] {
	inst.handlerWhere = fn
	return inst
}

func (inst *MemoryQuery[T]) Then(fn func(item T)) *MemoryQuery[T] {
	inst.handlerThen = fn
	return inst
}

func (inst *MemoryQuery[T]) Query(page *dxo.Pagination) error {

	if page == nil {
		page = new(dxo.Pagination)
	}

	// todo ...

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type MemoryLS struct {
	core *innerEngineCore

	id innerRowID

	tableName innerTableName

	item any

	handlerOnSetIntID func(id int64)
}

func (inst *MemoryLS) DoInsert() error {
	return inst.core.insert(inst)
}

func (inst *MemoryLS) DoUpdate() error {
	return inst.core.update(inst)
}

func (inst *MemoryLS) DoDelete() error {
	return inst.core.remove(inst)
}

func (inst *MemoryLS) DoFind() error {
	return inst.core.find(inst)
}

func (inst *MemoryLS) SetItem(item any) *MemoryLS {
	inst.item = item
	return inst
}

func (inst *MemoryLS) SetIntID(id int64) *MemoryLS {
	const base = 10
	str := strconv.FormatInt(id, base)
	inst.id = innerRowID(str)

	callback := inst.handlerOnSetIntID
	if callback != nil {
		callback(id)
	}

	return inst
}

func (inst *MemoryLS) OnSetIntID(callback func(id int64)) *MemoryLS {
	inst.handlerOnSetIntID = callback
	return inst
}

////////////////////////////////////////////////////////////////////////////////

func DoQuery[T any](ls *MemoryLS, model T) (*MemoryQuery[T], error) {

	core := ls.core

	tab, err := core.innerGetTableForLS(ls)
	if err != nil {
		return nil, err
	}

	mq := new(MemoryQuery[T])
	mq.ls = ls
	mq.table = tab

	return mq, nil
}

////////////////////////////////////////////////////////////////////////////////

type innerRow struct {
	id innerRowID

	index int

	body []byte // json-data of entity
}

func (inst *innerRow) encode(item any) error {
	bin, err := json.Marshal(item)
	if err == nil {
		inst.body = bin
	}
	return err
}

func (inst *innerRow) decode(item any) error {
	bin := inst.body
	return json.Unmarshal(bin, item)
}

////////////////////////////////////////////////////////////////////////////////

type innerTable struct {
	name innerTableName

	idcounter int64

	model any

	rows map[innerRowID]*innerRow

	mu sync.Mutex
}

func (inst *innerTable) init() error {
	inst.rows = make(map[innerRowID]*innerRow)
	return nil
}

func (inst *innerTable) nextIntID() int64 {

	mu := &inst.mu
	mu.Lock()
	defer mu.Unlock()

	inst.idcounter++
	n := inst.idcounter
	return n
}

////////////////////////////////////////////////////////////////////////////////

type innerEngineCore struct {
	tables map[innerTableName]*innerTable

	mu sync.Mutex
}

func (inst *innerEngineCore) innerGetTableForLS(ls *MemoryLS) (*innerTable, error) {

	tname, err := inst.innerGetTableNameForLS(ls)
	if err != nil {
		return nil, err
	}

	mu := &inst.mu
	mu.Lock()
	defer mu.Unlock()

	tab := inst.tables[tname]
	if tab == nil {
		return nil, fmt.Errorf("no table with name '%s'", tname)
	}
	return tab, nil
}

func (inst *innerEngineCore) innerGetTableNameForLS(ls *MemoryLS) (innerTableName, error) {
	mo := ls.item
	return inst.innerGetTableNameForModel(mo)
}

func (inst *innerEngineCore) innerGetTableNameForModel(model any) (innerTableName, error) {

	const empty = ""

	namer, ok := model.(innerTableNamer)
	if !ok {
		return empty, fmt.Errorf("the model object is not a table-entity")
	}

	name := namer.TableName()
	if name == empty {
		return empty, fmt.Errorf("the table name is empty")
	}

	return innerTableName(name), nil
}

func (inst *innerEngineCore) insert(ls *MemoryLS) error {

	tab, err := inst.innerGetTableForLS(ls)
	if err != nil {
		return err
	}

	idnum := tab.nextIntID()

	// lock

	mu := &tab.mu
	mu.Lock()
	defer mu.Unlock()

	// set id

	ls.SetIntID(idnum)
	idstr := ls.id
	item := ls.item

	// set time

	inst.innerTryUpdateItemTime(ls, true, true)

	// make row

	row := new(innerRow)
	row.id = idstr
	row.index = int(idnum)

	err = row.encode(item)
	if err != nil {
		return err
	}

	tab.rows[idstr] = row
	return nil
}

func (inst *innerEngineCore) update(ls *MemoryLS) error {

	pkey := ls.id
	item := ls.item

	// get table

	tab, err := inst.innerGetTableForLS(ls)
	if err != nil {
		return err
	}

	// lock

	mu := &tab.mu
	mu.Lock()
	defer mu.Unlock()

	// find

	row := tab.rows[pkey]
	if row == nil {
		return fmt.Errorf("no record")
	}

	// update time
	inst.innerTryUpdateItemTime(ls, false, true)

	// write
	err = row.encode(item)
	return err
}

func (inst *innerEngineCore) remove(ls *MemoryLS) error {

	pkey := ls.id

	// get table

	tab, err := inst.innerGetTableForLS(ls)
	if err != nil {
		return err
	}

	// lock

	mu := &tab.mu
	mu.Lock()
	defer mu.Unlock()

	// find

	row := tab.rows[pkey]
	if row == nil {
		return fmt.Errorf("no record")
	}

	// delete
	tab.rows[pkey] = nil
	return nil
}

func (inst *innerEngineCore) find(ls *MemoryLS) error {

	pkey := ls.id
	item := ls.item

	// get table

	tab, err := inst.innerGetTableForLS(ls)
	if err != nil {
		return err
	}

	// lock

	mu := &tab.mu
	mu.Lock()
	defer mu.Unlock()

	// find

	row := tab.rows[pkey]
	if row == nil {
		return fmt.Errorf("no record")
	}

	// read
	err = row.decode(item)
	return err
}

func (inst *innerEngineCore) query(ls *MemoryLS) error {
	return fmt.Errorf("no impl")
}

func (inst *innerEngineCore) innerTryUpdateItemTime(ls *MemoryLS, set_created_at bool, set_updated_at bool) {
	item := ls.item
	ref, ok := item.(rbac.EntityRef)
	if ok && (ref != nil) {
		tar := ref.GetTarget()
		if tar == nil {
			return
		}
		now := time.Now()
		if set_created_at {
			tar.CreatedAt = now
		}
		if set_updated_at {
			tar.UpdatedAt = now
		}
	}
}

func (inst *innerEngineCore) init() error {

	inst.tables = make(map[innerTableName]*innerTable)
	return nil
}

func (inst *innerEngineCore) initTable(model any) error {

	name, err := inst.innerGetTableNameForModel(model)
	if err != nil {
		return err
	}

	mu := &inst.mu
	mu.Lock()
	defer mu.Unlock()
	all := inst.tables

	if err == nil {
		tab := new(innerTable)
		tab.name = innerTableName(name)
		tab.model = model
		err = tab.init()
		if err == nil {
			all[tab.name] = tab
		}
	}

	return err
}

////////////////////////////////////////////////////////////////////////////////

type MemoryEngineFacade struct {

	//starter:component

	_as func(IMemoryEngine) //starter:as("#")

	core *innerEngineCore
}

// GetDB implements [IMemoryEngine].
func (inst *MemoryEngineFacade) GetDB(old *gorm.DB) *gorm.DB {
	if old == nil {
		old = new(gorm.DB)
	}
	return old
}

// NextUUID implements [IMemoryEngine].
func (inst *MemoryEngineFacade) NextUUID() lang.UUID {
	var buffer [16]byte
	core, _ := inst.innerGetCore()
	if core != nil {
		mu := &core.mu
		mu.Lock()
		defer mu.Unlock()
		rand.Read(buffer[:])
	}
	return lang.NewUUID(buffer[:])
}

func (inst *MemoryEngineFacade) innerGetCore() (*innerEngineCore, error) {
	core := inst.core
	if core == nil {
		c2, err := inst.innerLoadCore()
		if err != nil {
			return nil, err
		}
		core = c2
		inst.core = c2
	}
	return core, nil
}

func (inst *MemoryEngineFacade) innerLoadCore() (*innerEngineCore, error) {
	core := new(innerEngineCore)
	err := core.init()
	return core, err
}

// InitTable implements [IMemoryEngine].
func (inst *MemoryEngineFacade) InitTable(model any) error {

	core, err := inst.innerGetCore()
	if err != nil {
		return err
	}
	err = core.initTable(model)
	return err
}

// NewLS implements [IMemoryEngine].
func (inst *MemoryEngineFacade) NewLS() *MemoryLS {
	ls := new(MemoryLS)
	ls.core = inst.core
	return ls
}

func (inst *MemoryEngineFacade) _impl() IMemoryEngine {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
