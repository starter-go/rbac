package dxo

import (
	"strings"
)

// RegionID ...
type RegionID int

// RegionPhoneCode 是数字形式的国际电话区号，
// 例如：中国(86)， 法国(33)， 俄国(7)， 美国(1)， 英国(44)
type RegionPhoneCode string

// RegionCode2 是 ISO 3166-1 标准的二字节地区码
// 例如：中国(CN)， 法国(FR)， 俄国(RU)， 美国(US)， 英国(GB)
type RegionCode2 string

// RegionCode3 是 ISO 3166-1 标准的三字节地区码
// 例如：中国(CHN)， 法国(FRA)， 俄国(RUS)， 美国(USA)， 英国(GBR)
type RegionCode3 string

////////////////////////////////////////////////////////////////////////////////

func (code RegionPhoneCode) String() string {
	return string(code)
}

// Normalize 标准化代码
func (code RegionPhoneCode) Normalize() RegionPhoneCode {
	str := code.String()
	str = normalizePureNumber(str)
	return RegionPhoneCode(str)
}

////////////////////////////////////////////////////////////////////////////////

func (code RegionCode2) String() string {
	return string(code)
}

// Normalize 标准化代码
func (code RegionCode2) Normalize() RegionCode2 {
	str := code.String()
	str = strings.TrimSpace(str)
	str = strings.ToUpper(str)
	return RegionCode2(str)
}

////////////////////////////////////////////////////////////////////////////////

func (code RegionCode3) String() string {
	return string(code)
}

// Normalize 标准化代码
func (code RegionCode3) Normalize() RegionCode3 {
	str := code.String()
	str = strings.TrimSpace(str)
	str = strings.ToUpper(str)
	return RegionCode3(str)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
