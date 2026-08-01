package authorizations

import "github.com/starter-go/rbac/lib/dxo"

// 授权 api

type Want struct {
	Action dxo.AuthAction
}

type Have struct {
	OK    bool
	Error error
}

type Authorization struct {
	Want *Want
	Have *Have
}
