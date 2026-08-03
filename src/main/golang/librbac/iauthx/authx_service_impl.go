package iauthx

import (
	"context"

	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/lib/classes/authx"
)

type AuthxServiceImpl struct {

	//starter:component

	_as func(rbac.AuthService) //starter:as("#")

	RegistryList []authx.Registry //starter:as(".")

	cache *innerAuthxCache
}

// Handle implements [rbac.AuthService].
func (inst *AuthxServiceImpl) Handle(c *authx.Context) error {
	panic("unimplemented")
}

// HandleDTO implements [rbac.AuthService].
func (inst *AuthxServiceImpl) HandleDTO(cc context.Context, items []*authx.AuthDTO) ([]*authx.AuthDTO, error) {
	panic("unimplemented")
}

func (inst *AuthxServiceImpl) innerGetCache() (*innerAuthxCache, error) {
	c := inst.cache
	if c == nil {
		c2, err := inst.innerLoadCache()
		if err != nil {
			return nil, err
		}
		c = c2
		inst.cache = c2
	}
	return c, nil
}

func (inst *AuthxServiceImpl) innerLoadCache() (*innerAuthxCache, error) {
	ldr := new(innerAuthxCacheLoader)
	return ldr.load(inst.RegistryList)
}

func (inst *AuthxServiceImpl) _impl() rbac.AuthService {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// inner types

type innerAuthxCache struct {
	authenticators []*innerAuthenticatorHolder

	authorizers []*innerAuthorizerHolder
}

type innerAuthenticatorHolder struct {
	info *authx.Registration

	authenticator authx.Authenticator
	mechanism     authx.Mechanism
	priority      int
}

type innerAuthorizerHolder struct {
	info *authx.Registration

	authorizer authx.Authorizer
	action     authx.Action
	priority   int
}

type innerAuthxCacheLoader struct {
}

////////////////////////////////////////////////////////////////////////////////

func (inst *innerAuthxCacheLoader) load(src []authx.Registry) (*innerAuthxCache, error) {

	// todo ...
	panic("no impl")
}

////////////////////////////////////////////////////////////////////////////////
// EOF
