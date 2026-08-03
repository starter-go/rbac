package checkers

import (
	"context"

	"github.com/starter-go/rbac/lib/dxo"
)

type Checking struct {
	Context context.Context

	Action Action

	TableName dxo.TableName

	Entities []dxo.EntityRef

	Dtos []dxo.DTORef
}

type Checker interface {
	Check(c *Checking, next CheckerChain) error
}

type CheckerChain interface {
	Check(c *Checking) error
}
