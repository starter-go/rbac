package dxo

import (
	"time"

	"github.com/starter-go/base/lang"
)

// Term 表示有效期
type Term struct {
	NotBefore lang.Time `json:"not_before"` // 生效时间, '不早于'
	NotAfter  lang.Time `json:"not_after"`  // 过期时间, '不晚于'
}

func (inst *Term) SetTermFromTo(from, to time.Time) *Term {

	if inst == nil {
		inst = new(Term)
	}

	t1 := lang.NewTime(from)
	t2 := lang.NewTime(to)
	inst.NotBefore = t1
	inst.NotAfter = t2

	return inst
}

func (inst *Term) SetTermWithAge(from time.Time, maxAge time.Duration) *Term {
	to := from.Add(maxAge)
	return inst.SetTermFromTo(from, to)
}

func (inst *Term) IsInTerm(t time.Time) bool {
	t1 := inst.NotBefore
	t2 := inst.NotAfter
	tt := lang.NewTime(t)
	return (t1 <= tt) && (tt <= t2)
}
