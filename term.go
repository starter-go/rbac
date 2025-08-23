package rbac

import (
	"time"

	"github.com/starter-go/base/lang"
)

// Term 表示有效期
type Term struct {
	StartedAt lang.Time `json:"started_at"` // 生效时间
	ExpiredAt lang.Time `json:"expired_at"` // 过期时间
}

func (inst *Term) SetTermFromTo(from, to time.Time) *Term {
	if inst != nil {
		t1 := lang.NewTime(from)
		t2 := lang.NewTime(to)
		inst.StartedAt = t1
		inst.ExpiredAt = t2
	}
	return inst
}

func (inst *Term) SetTermWithAge(from time.Time, maxAge time.Duration) *Term {
	to := from.Add(maxAge)
	return inst.SetTermFromTo(from, to)
}

func (inst *Term) IsInTerm(t time.Time) bool {
	t1 := inst.StartedAt
	t2 := inst.ExpiredAt
	tt := lang.NewTime(t)
	return (t1 <= tt) && (tt <= t2)
}
