package dxo

// EmailAddress 表示 'user@domain' 形式的邮件地址
type EmailAddress string

// EmailAddressID ...
type EmailAddressID int64

////////////////////////////////////////////////////////////////////////////////

func (addr EmailAddress) String() string {
	return string(addr)
}
