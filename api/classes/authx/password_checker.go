package authx

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/starter-go/base/lang"
)

type PasswordChecker struct {

	// common

	salt lang.Hex

	password lang.Hex

	account lang.Hex

	//////////////////////////////////////
	// raw (from UI)

	raw lang.SumSHA256 // the raw sum

	//////////////////////////////////////
	// target (from db)

	target lang.SumSHA256 // the target sum
}

func (inst *PasswordChecker) SetSalt(salt []byte) *PasswordChecker {
	inst.salt = lang.HexFromBytes(salt)
	return inst
}

func (inst *PasswordChecker) SetPassword(password []byte) *PasswordChecker {
	inst.password = lang.HexFromBytes(password)
	return inst
}

func (inst *PasswordChecker) SetAccount(account []byte) *PasswordChecker {
	inst.account = lang.HexFromBytes(account)
	return inst
}

func (inst *PasswordChecker) SetTarget(sum []byte) *PasswordChecker {
	src := sum
	dst := inst.target[:]
	copy(dst, src)
	return inst
}

func (inst *PasswordChecker) GetTargetSum() lang.Sum {
	return inst.target
}

func (inst *PasswordChecker) GetRawSum() lang.Sum {
	return inst.raw
}

func (inst *PasswordChecker) Compute() {

	buffer := bytes.NewBuffer(inst.account.Bytes())
	buffer.WriteByte(0)
	buffer.Write(inst.password.Bytes())
	buffer.WriteByte(0)
	buffer.Write(inst.salt.Bytes())

	bin := buffer.Bytes()
	sum := sha256.Sum256(bin)
	inst.raw = sum
}

func (inst *PasswordChecker) Check() error {

	s1 := inst.raw[:]
	s2 := inst.target[:]
	eq := bytes.Equal(s1, s2)

	if !eq {
		return fmt.Errorf("bad credential")
	}

	return nil
}
