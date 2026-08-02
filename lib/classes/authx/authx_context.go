package authx

import (
	"context"

	"github.com/starter-go/rbac/lib/dxo"
)

type Context struct {
	CC context.Context

	Authentications []*Authentication

	Authorizations []*Authorization

	UserInfo *dxo.UserInfo

	Authenticated bool

	Cancelled bool

	Use2FA bool
}

func (inst *Context) Cancel() {
	inst.Cancelled = true
}
