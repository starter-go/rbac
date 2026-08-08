package iauthx

import (
	"context"
	"fmt"
	"sort"

	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/authx"
)

type AuthxServiceImpl struct {

	//starter:component

	_as func(rbac.AuthService) //starter:as("#")

	RegistryList []authx.Registry //starter:inject(".")

	cache *innerAuthxCache
}

// Handle implements [rbac.AuthService].
func (inst *AuthxServiceImpl) Handle(c *authx.Context) error {

	cache, err := inst.innerLoadCache()
	if err != nil {
		return err
	}

	for _, a1 := range c.Authentications {
		a1.Context = c
		err := cache.handleAuth1(a1)
		if err != nil {
			return err
		}
	}

	if c.Cancelled {
		return nil
	}

	err = inst.innerCheckAuth1(c)
	if err != nil {
		return fmt.Errorf("challenge: %s", err.Error())
	}

	c.Authenticated = true

	for _, a2 := range c.Authorizations {
		a2.Context = c
		err := cache.handleAuth2(a2)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *AuthxServiceImpl) innerCheckAuth1(c *authx.Context) error {

	a1all := c.Authentications
	use2fa := c.Use2FA
	info := c.UserInfo
	a1count := len(a1all)

	for _, a1 := range a1all {
		info1 := info
		info2 := a1.UserInfo
		if info == nil {
			info = info2
			info1 = info2
		}
		if info1 != nil && info2 != nil {
			if info1.UserID != info2.UserID {
				return fmt.Errorf("bad auth")
			}
		}
		if !a1.OK {
			return fmt.Errorf("bad auth")
		}
	}

	if info == nil {
		return fmt.Errorf("bad auth")
	}

	if info.Use2FA {
		use2fa = true
	}
	if info.Locked {
		return fmt.Errorf("account is locked")
	}
	if !info.Enabled {
		return fmt.Errorf("account is disabled")
	}

	if use2fa {
		if a1count < 2 {
			return fmt.Errorf("2FA is required")
		}
	} else {
		if a1count < 1 {
			return fmt.Errorf("no auth")
		}
	}

	return nil
}

// HandleDTO implements [rbac.AuthService].
func (inst *AuthxServiceImpl) HandleDTO(cc context.Context, items []*authx.AuthDTO) ([]*authx.AuthDTO, error) {

	ctx2 := new(authx.Context)
	ctx2.CC = cc

	err := inst.innerConvertRequest(ctx2, items)
	if err != nil {
		return nil, err
	}

	err = inst.Handle(ctx2)
	if err != nil {
		return nil, err
	}

	return inst.innerConvertResponse(ctx2)
}

func (inst *AuthxServiceImpl) innerConvertRequest(ctx *authx.Context, items []*authx.AuthDTO) error {

	convertor := new(innerAuthObjectConvertor)

	for _, it := range items {
		if convertor.isAuthentication(it) {
			a1 := new(rbac.Authentication)
			convertor.convertAuthenticationD2A(it, a1)
			ctx.Authentications = append(ctx.Authentications, a1)
		}
		if convertor.isAuthorization(it) {
			a2 := new(rbac.Authorization)
			convertor.convertAuthorizationD2A(it, a2)
			ctx.Authorizations = append(ctx.Authorizations, a2)
		}
	}

	return nil
}

func (inst *AuthxServiceImpl) innerConvertResponse(ctx *authx.Context) ([]*authx.AuthDTO, error) {

	src1 := ctx.Authentications
	src2 := ctx.Authorizations
	dst := make([]*authx.AuthDTO, 0)
	convertor := new(innerAuthObjectConvertor)

	for _, it1 := range src1 {
		it2 := new(authx.AuthDTO)
		convertor.convertAuthenticationA2D(it1, it2)
		dst = append(dst, it2)
	}

	for _, it1 := range src2 {
		it2 := new(authx.AuthDTO)
		convertor.convertAuthorizationA2D(it1, it2)
		dst = append(dst, it2)
	}

	return dst, nil
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
	authorizers    []*innerAuthorizerHolder
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

func (inst *innerAuthxCache) handleAuth1(a1 *authx.Authentication) error {

	all := inst.authenticators

	for _, reg := range all {
		if a1.Mechanism != reg.mechanism {
			continue
		}
		handler := reg.authenticator
		if !handler.Accept(a1) {
			continue
		}
		err := handler.Authenticate(a1)
		if err == nil {
			return nil
		} else {
			return err
		}
	}

	return fmt.Errorf("no authenticator accept the auth-request")
}

func (inst *innerAuthxCache) handleAuth2(a2 *authx.Authorization) error {

	all := inst.authorizers

	for _, reg := range all {
		if a2.Action != reg.action {
			continue
		}
		handler := reg.authorizer
		if !handler.Accept(a2) {
			continue
		}
		err := handler.Authorize(a2)
		if err == nil {
			return nil
		} else {
			return err
		}
	}

	return fmt.Errorf("no authorizer accept the auth-request")
}

////////////////////////////////////////////////////////////////////////////////

func (inst *innerAuthxCacheLoader) load(src []authx.Registry) (*innerAuthxCache, error) {

	cache := new(innerAuthxCache)

	for _, r := range src {
		if r == nil {
			continue
		}
		for _, reg := range r.ListRegistrations() {
			inst.accept(cache, reg)
		}
	}

	sort.Slice(cache.authenticators, func(i1, i2 int) bool {
		a1 := cache.authenticators[i1]
		a2 := cache.authenticators[i2]
		return a1.priority < a2.priority
	})

	sort.Slice(cache.authorizers, func(i1, i2 int) bool {
		a1 := cache.authorizers[i1]
		a2 := cache.authorizers[i2]
		return a1.priority < a2.priority
	})

	return cache, nil
}

func (inst *innerAuthxCacheLoader) accept(cache *innerAuthxCache, reg *authx.Registration) {

	if reg == nil {
		return
	}
	if !reg.Enabled {
		return
	}

	if reg.Authenticator != nil {
		h := new(innerAuthenticatorHolder)
		h.info = reg
		h.authenticator = reg.Authenticator
		h.mechanism = reg.Mechanism
		h.priority = reg.Priority
		cache.authenticators = append(cache.authenticators, h)
	}

	if reg.Authorizer != nil {
		h := new(innerAuthorizerHolder)
		h.info = reg
		h.authorizer = reg.Authorizer
		h.action = reg.Action
		h.priority = reg.Priority
		cache.authorizers = append(cache.authorizers, h)
	}
}

////////////////////////////////////////////////////////////////////////////////
// EOF
