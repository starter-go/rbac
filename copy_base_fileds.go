package rbac

import (
	"fmt"
	"time"

	"github.com/starter-go/base/lang"
)

func CopyBaseFieldsE2D(src EntityRef, dst DTORef) error {

	const errMsg = "rbac.CopyBaseFieldsD2E() : param is nil"
	if src == nil || dst == nil {
		return fmt.Errorf(errMsg)
	}
	s2 := src.GetTarget()
	d2 := dst.GetTarget()
	if s2 == nil || d2 == nil {
		return fmt.Errorf(errMsg)
	}

	d2.UUID = s2.UUID
	d2.Creator = s2.Creator
	d2.Updater = s2.Updater
	d2.Owner = s2.Owner
	d2.Group = s2.Group
	d2.CreatedAt = innerConvertTimeT(s2.CreatedAt)
	d2.UpdatedAt = innerConvertTimeT(s2.UpdatedAt)

	return nil
}

func CopyBaseFieldsD2E(src DTORef, dst EntityRef) error {

	const errMsg = "rbac.CopyBaseFieldsD2E() : param is nil"
	if src == nil || dst == nil {
		return fmt.Errorf(errMsg)
	}
	s2 := src.GetTarget()
	d2 := dst.GetTarget()
	if s2 == nil || d2 == nil {
		return fmt.Errorf(errMsg)
	}

	d2.UUID = s2.UUID
	d2.Creator = s2.Creator
	d2.Updater = s2.Updater
	d2.Owner = s2.Owner
	d2.Group = s2.Group
	d2.CreatedAt = innerConvertTimeL(s2.CreatedAt)
	d2.UpdatedAt = innerConvertTimeL(s2.UpdatedAt)

	return nil
}

func innerConvertTimeL(n lang.Time) time.Time {
	return n.Time()
}

func innerConvertTimeT(t time.Time) lang.Time {
	return lang.NewTime(t)
}
