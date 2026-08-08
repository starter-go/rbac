package iauth

import (
	"fmt"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac"
	"github.com/starter-go/rbac/api/classes/authx"
	"github.com/starter-go/rbac/api/classes/users"
)

type PasswordAuthenticator struct {

	//starter:component

	_as func(authx.Registry) //starter:as(".")

	UserDao users.DAO //starter:as("#")
}

// Accept implements [authx.Authenticator].
func (inst *PasswordAuthenticator) Accept(a1 *authx.Authentication) bool {
	return (a1.Mechanism == rbac.MechanismPassword)
}

// Authenticate implements [authx.Authenticator].
func (inst *PasswordAuthenticator) Authenticate(a1 *authx.Authentication) error {

	cname := a1.CommonName

	user, err := inst.innerFindUserByCommonName(cname)
	if err != nil {
		return err
	}

	if user.Locked {
		return fmt.Errorf("account is locked")
	}
	if !user.Enabled {
		return fmt.Errorf("account is disabled")
	}

	err = inst.innerCheckPassword(a1, user)
	if err != nil {
		return err
	}

	info := inst.innerGetUserInfo(user)

	a1.Message = "OK"
	a1.UserInfo = info
	a1.OK = true
	return nil
}

func (inst *PasswordAuthenticator) innerFindUserByCommonName(commonName string) (*users.Entity, error) {

	// by [username|user_id|phone|email|...]

	finder := new(innerUserEntityFinder)
	finder.init(inst.UserDao, commonName)
	return finder.findUser()
}

func (inst *PasswordAuthenticator) innerCheckPassword(a1 *authx.Authentication, user *users.Entity) error {

	uuid := user.UUID.Bytes()
	target := user.Password.Bytes()
	password := a1.Secret
	salt := user.Salt.Bytes()
	checker := new(authx.PasswordChecker)

	checker.SetAccount(uuid)
	checker.SetTarget(target)
	checker.SetPassword(password)
	checker.SetSalt(salt)

	checker.Compute()

	return checker.Check()
}

func (inst *PasswordAuthenticator) innerGetUserInfo(user *users.Entity) *rbac.UserInfo {

	info := new(rbac.UserInfo)

	return info
}

// ListRegistrations implements [authx.Registry].
func (inst *PasswordAuthenticator) ListRegistrations() []*authx.Registration {
	cl := lang.ClassOf(inst)
	r1 := &authx.Registration{
		Label:         cl.SimpleName(),
		Mechanism:     rbac.MechanismPassword,
		Priority:      0,
		Enabled:       true,
		Authenticator: inst,
	}
	return []*authx.Registration{r1}
}

func (inst *PasswordAuthenticator) _impl() (authx.Registry, authx.Authenticator) {
	return inst, inst
}

////////////////////////////////////////////////////////////////////////////////

type innerUserEntityFinder struct {
	cname string
	dao   users.DAO

	countMarkAt int
	countDigit  int
	countABC    int
}

func (inst *innerUserEntityFinder) init(dao users.DAO, cname string) {
	inst.cname = cname
	inst.dao = dao
	inst.checkCommonName()
}

func (inst *innerUserEntityFinder) checkCommonName() {
	array := []rune(inst.cname)
	for _, ch := range array {
		if ('0' <= ch) && (ch <= '9') {
			inst.countDigit++
		} else if ('a' <= ch) && (ch <= 'z') {
			inst.countABC++
		} else if ('A' <= ch) && (ch <= 'Z') {
			inst.countABC++
		} else if ch == '@' {
			inst.countMarkAt++
		}
	}
}

func (inst *innerUserEntityFinder) isPhoneNumber() bool {
	mark := inst.countMarkAt
	numb := inst.countDigit
	abcd := inst.countABC
	return ((mark == 0) && (abcd == 0) && (numb > 0))
}

func (inst *innerUserEntityFinder) isEmailAddr() bool {
	mark := inst.countMarkAt
	numb := inst.countDigit
	abcd := inst.countABC
	return ((mark == 1) && ((abcd + numb) > 0))
}

func (inst *innerUserEntityFinder) isUserName() bool {
	mark := inst.countMarkAt
	numb := inst.countDigit
	abcd := inst.countABC
	return ((mark == 0) && ((abcd + numb) > 0))
}

func (inst *innerUserEntityFinder) findUser() (*rbac.UserEntity, error) {

	dao := inst.dao
	name := inst.cname

	if inst.isEmailAddr() {

		// find by email
		return dao.FindByEmail(nil, users.EmailAddress(name))

	} else if inst.isPhoneNumber() {

		// find by phone
		return dao.FindByPhone(nil, users.PhoneNumber(name))

	} else if inst.isUserName() {

		// find by username
		return dao.FindByName(nil, users.UserName(name))
	}

	return nil, fmt.Errorf("no account")
}

////////////////////////////////////////////////////////////////////////////////
