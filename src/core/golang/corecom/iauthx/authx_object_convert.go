package iauthx

import (
	"fmt"
	"strings"

	"github.com/starter-go/application/parameters"
	"github.com/starter-go/rbac"
)

type innerAuthObjectConvertor struct {
}

func (inst *innerAuthObjectConvertor) convertAuthenticationD2A(src *rbac.AuthDTO, dst *rbac.Authentication) error {

	dst.Params = parameters.NewTable(nil)
	dst.Params.Import(src.Parameters)

	dst.Mechanism = src.Mechanism
	dst.Step = src.Step
	dst.CommonName = src.Account
	dst.Secret = src.Secret.Bytes()

	dst.Challenge = src.Challenge
	dst.Error = inst.convertErrorS2E(src.Error)
	dst.Message = src.Message
	dst.OK = src.OK

	return nil
}

func (inst *innerAuthObjectConvertor) convertAuthorizationD2A(src *rbac.AuthDTO, dst *rbac.Authorization) error {

	// params

	dst.Params = parameters.NewTable(nil)
	dst.Params.Import(src.Parameters)

	dst.Action = src.Action
	dst.Step = src.Step

	// result

	dst.Error = inst.convertErrorS2E(src.Error)
	dst.Message = src.Message
	dst.OK = src.OK

	return nil
}

func (inst *innerAuthObjectConvertor) convertAuthenticationA2D(src *rbac.Authentication, dst *rbac.AuthDTO) error {

	// params

	params := src.Params
	if params == nil {
		dst.Parameters = make(map[string]string)
	} else {
		dst.Parameters = params.Export(nil)
	}

	dst.Mechanism = src.Mechanism
	dst.Step = src.Step
	dst.Account = src.CommonName

	// result

	dst.Challenge = src.Challenge
	dst.Error = inst.convertErrorE2S(src.Error)
	dst.Message = src.Message
	dst.OK = src.OK

	return nil
}

func (inst *innerAuthObjectConvertor) convertAuthorizationA2D(src *rbac.Authorization, dst *rbac.AuthDTO) error {

	// params

	params := src.Params
	if params == nil {
		dst.Parameters = make(map[string]string)
	} else {
		dst.Parameters = params.Export(nil)
	}

	dst.Action = src.Action
	dst.Step = src.Step

	// result

	dst.Error = inst.convertErrorE2S(src.Error)
	dst.Message = src.Message
	dst.OK = src.OK

	return nil
}

func (inst *innerAuthObjectConvertor) isAuthorization(dto *rbac.AuthDTO) bool {

	if dto == nil {
		return false
	}

	if dto.Action == "" {
		return false
	}

	return true
}

func (inst *innerAuthObjectConvertor) isAuthentication(dto *rbac.AuthDTO) bool {

	if dto == nil {
		return false
	}

	if dto.Mechanism == "" {
		return false
	}

	return true
}

func (inst *innerAuthObjectConvertor) convertErrorS2E(err string) error {
	err = strings.TrimSpace(err)
	if err == "" {
		return nil
	}
	return fmt.Errorf(err)
}

func (inst *innerAuthObjectConvertor) convertErrorE2S(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
