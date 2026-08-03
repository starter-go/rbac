package phonenumbers

import "github.com/starter-go/rbac/lib/dxo"

////////////////////////////////////////////////////////////////////////////////

type PhoneNumberID = dxo.PhoneNumberID

type FullPhoneNumber = dxo.FullPhoneNumber

type SimplePhoneNumber = dxo.SimplePhoneNumber

type RegionPhoneCode = dxo.RegionPhoneCode

type Pagination = dxo.Pagination

////////////////////////////////////////////////////////////////////////////////

// PhoneNumberDTO ...
type PhoneNumberDTO struct {
	ID PhoneNumberID `json:"id"`

	dxo.DTO

	RegionCode2  RegionPhoneCode   `json:"region"`
	SimpleNumber SimplePhoneNumber `json:"simple_number"`
	FullNumber   FullPhoneNumber   `json:"full_number"`
}

// PhoneNumberEntity ...
type PhoneNumberEntity struct {
	ID PhoneNumberID `json:"id"`

	dxo.Entity

	RegionCode2  RegionPhoneCode
	SimpleNumber SimplePhoneNumber `gorm:"unique"`
	FullNumber   FullPhoneNumber
}

////////////////////////////////////////////////////////////////////////////////
// EOF
