package checkers

import "sort"

////////////////////////////////////////////////////////////////////////////////

type CheckerChainBuilder struct {
	items []*CheckerRegistration
}

func (inst *CheckerChainBuilder) AddR1(list ...CheckerRegistry) {

	for _, it := range list {
		if it == nil {
			continue
		}
		r2 := it.Registration()
		inst.AddR2(r2)
	}

}

func (inst *CheckerChainBuilder) AddR2(list ...*CheckerRegistration) {

	for _, it := range list {
		if inst.accept(it) {
			inst.items = append(inst.items, it)
		}
	}

}

// Len implements [sort.Interface].
func (inst *CheckerChainBuilder) Len() int {
	return len(inst.items)
}

// Less implements [sort.Interface].
func (inst *CheckerChainBuilder) Less(i1, i2 int) bool {
	n1 := inst.items[i1].Order
	n2 := inst.items[i2].Order
	return (n1 < n2)
}

// Swap implements [sort.Interface].
func (inst *CheckerChainBuilder) Swap(i1, i2 int) {
	l := inst.items
	l[i1], l[i2] = l[i2], l[i1]
}

func (inst *CheckerChainBuilder) GetItem(index int) *CheckerRegistration {
	return inst.items[index]
}

func (inst *CheckerChainBuilder) Build() CheckerChain {

	inst.sort()

	var chain CheckerChain = nil
	all := inst.items
	end := new(innerCheckerChainEnd)

	chain = end

	for _, it := range all {
		if inst.accept(it) {
			node := new(innerCheckerChainNode)
			node.checker = it.Checker
			node.next = chain
			chain = node
		}
	}

	return chain
}

func (inst *CheckerChainBuilder) sort() {
	sort.Sort(inst)
}

func (inst *CheckerChainBuilder) accept(item *CheckerRegistration) bool {

	if item == nil {
		return false
	}

	if item.Checker == nil {
		return false
	}

	if !item.Enabled {
		return false
	}

	return true
}

////////////////////////////////////////////////////////////////////////////////

type innerCheckerChainNode struct {
	next    CheckerChain
	checker Checker
}

// Check implements [CheckerChain].
func (inst *innerCheckerChainNode) Check(ctx *Checking) error {
	n := inst.next
	c := inst.checker
	return c.Check(ctx, n)
}

func (inst *innerCheckerChainNode) _impl() CheckerChain {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerCheckerChainEnd struct{}

// Check implements [CheckerChain].
func (inst *innerCheckerChainEnd) Check(c *Checking) error {
	return nil
}

func (inst *innerCheckerChainEnd) _impl() CheckerChain {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
