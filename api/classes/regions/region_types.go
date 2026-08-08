package regions

import (
	"github.com/starter-go/base/lang"
	"github.com/starter-go/rbac/api/dxo"
)

type RegionID = dxo.RegionID

////////////////////////////////////////////////////////////////////////////////
// short type names

type ID = RegionID

type DTO = RegionDTO

type VO = RegionVO

type Entity = RegionEntity

type Pagination = dxo.Pagination

type RegionCode2 = dxo.RegionCode2

type RegionCode3 = dxo.RegionCode3

type RegionPhoneCode = dxo.RegionPhoneCode

type RegionService = Service

type RegionQuery = Query

////////////////////////////////////////////////////////////////////////////////

type RegionDTO struct {
	ID ID `json:"id"`

	dxo.BaseDTO

	FlagURL     lang.URL        `json:"flag_url"`    // 国旗（或区旗）图标的URL
	DisplayName string          `json:"label"`       // 显示名称，通常是本地化的名称
	SimpleName  string          `json:"simple_name"` // 区域简称，例如：chn(中国), fra(France), usa(United States)
	FullName    string          `json:"full_name"`   // 完整的名称，例如：中华人民共和国(PRC)
	Code2       RegionCode2     `json:"code_xx"`     // 二字符区域码
	Code3       RegionCode3     `json:"code_xxx"`    // 三字符区域码
	PhoneCode   RegionPhoneCode `json:"phone_code"`  // 电话区域码

}

type RegionVO struct {
	dxo.BaseVO

	Items []*DTO `json:"regions"`
}

type RegionEntity struct {
	ID ID

	dxo.BaseEntity

	FlagURL     lang.URL        // 国旗（或区旗）图标的URL
	DisplayName string          // 显示名称，通常是本地化的名称
	SimpleName  string          // 区域简称，例如：chn(中国), fra(France), usa(United States)
	FullName    string          // 完整的名称，例如：中华人民共和国(PRC)
	PhoneCode   RegionPhoneCode // 电话区域码

	Code2 RegionCode2 `gorm:"unique"` // 二字符区域码
	Code3 RegionCode3 `gorm:"unique"` // 三字符区域码

}

////////////////////////////////////////////////////////////////////////////////
