package iservices

import (
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/checkers"
	"github.com/starter-go/vlog"
)

type CheckerServiceImpl struct {

	//starter:component

	_as func(rbac.CheckerService) //starter:as("#")

	RegList []rbac.CheckerRegistry //starter:inject(".")

	cache *innerCheckerChainCache
}

// Check implements [checkers.Service].
func (inst *CheckerServiceImpl) Check(c *checkers.Checking) error {
	cache, err := inst.innerGetCache()
	if err != nil {
		return err
	}
	return cache.chain.Check(c)
}

func (inst *CheckerServiceImpl) innerGetCache() (*innerCheckerChainCache, error) {
	cache := inst.cache
	if cache == nil {
		ldr := new(innerCheckerChainLoader)
		c2, err := ldr.load(inst.RegList)
		if err != nil {
			return nil, err
		}
		cache = c2
		inst.cache = c2
	}
	return cache, nil
}

func (inst *CheckerServiceImpl) _impl() rbac.CheckerService {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerCheckerChainLoader struct {
}

func (inst *innerCheckerChainLoader) load(src []rbac.CheckerRegistry) (*innerCheckerChainCache, error) {

	builder := new(checkers.CheckerChainBuilder)
	builder.AddR1(src...)
	chain := builder.Build()
	cache := new(innerCheckerChainCache)

	cache.chain = chain
	inst.logChain(builder)

	return cache, nil
}

func (inst *innerCheckerChainLoader) logChain(b *checkers.CheckerChainBuilder) {

	if !vlog.IsDebugEnabled() {
		return
	}

	count := b.Len()

	vlog.Debug("<rbac:checkers>")

	for i := 0; i < count; i++ {
		it := b.GetItem(i)
		label := it.Label
		order := it.Order
		en := it.Enabled
		vlog.Debug("[rbac.Checker index:%d enabled:%v order:%d label:'%s']", i, en, order, label)
	}

	vlog.Debug("</rbac:checkers>")
}

////////////////////////////////////////////////////////////////////////////////

type innerCheckerChainCache struct {
	chain rbac.CheckerChain
}

////////////////////////////////////////////////////////////////////////////////
// EOF
