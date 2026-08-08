package authx

import "testing"

func TestPasswordChecker(t *testing.T) {

	user := "foo"
	pass := "bar"
	salt := []byte{1, 2, 3, 4}
	target := []byte{5, 6, 7, 8}

	chr := new(PasswordChecker)
	chr.SetAccount([]byte(user))
	chr.SetPassword([]byte(pass))
	chr.SetSalt(salt)
	chr.SetTarget(target)

	chr.Compute()
	err := chr.Check()

	t.Log("sum(raw) = ", chr.GetRawSum().String())
	t.Log("sum(tar) = ", chr.GetTargetSum().String())
	t.Log("result = ", err)

}
