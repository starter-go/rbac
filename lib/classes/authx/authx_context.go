package authx

import (
	"context"

	"github.com/starter-go/rbac/lib/classes/authentications"
	"github.com/starter-go/rbac/lib/classes/authorizations"
	"github.com/starter-go/rbac/lib/dxo"
)

type Authentication = authentications.Authentication

type Authorization = authorizations.Authorization

type Context struct {
	CC context.Context

	Authentications []*Authentication
	Authorizations  []*Authorization

	User *dxo.UserInfo

	Authenticated bool

	Use2FA bool
}
