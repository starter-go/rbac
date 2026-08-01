package authentications

import "github.com/starter-go/rbac/lib/dxo"

// 验证 api

type Want struct {
	CommonName string // [user_id, phone_num, email_addr, ...]

	Mechanism dxo.AuthMechanism

	Step dxo.AuthStep

	Secret []byte
}

type Have struct {
	Challenge string
	Message   string
	Error     error
	OK        bool

	User *dxo.UserInfo
}

type Authentication struct {
	Want *Want
	Have *Have
}
