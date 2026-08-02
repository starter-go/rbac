package iauthx

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/lib/classes/authx"
)

type AuthxServiceImpl struct {

	//starter:component

	_as func(rbac.AuthService) //starter:as("#")

}

// Handle implements [rbac.AuthService].
func (inst *AuthxServiceImpl) Handle(c *authx.Context) error {
	panic("unimplemented")
}

// HandleDTO implements [rbac.AuthService].
func (inst *AuthxServiceImpl) HandleDTO(cc context.Context, items []*authx.AuthDTO) ([]*authx.AuthDTO, error) {
	panic("unimplemented")
}

func (inst *AuthxServiceImpl) _impl() rbac.AuthService {
	return inst
}
