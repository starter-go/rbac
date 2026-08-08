package dxo

import "strings"

// DomainName 以字符串形式表示一个域名
type DomainName string

func (dn DomainName) String() string {
	return string(dn)
}

func (dn DomainName) Normalize() DomainName {
	str := dn.String()
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	return DomainName(str)
}
