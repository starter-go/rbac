package phonenumbers

import "github.com/starter-go/rbac/lib/dxo"

////////////////////////////////////////////////////////////////////////////////

type PhoneNumberID = dxo.PhoneNumberID

type FullPhoneNumber = dxo.FullPhoneNumber

type SimplePhoneNumber = dxo.SimplePhoneNumber

type RegionPhoneCode = dxo.RegionPhoneCode

////////////////////////////////////////////////////////////////////////////////
// short aliases

type Pagination = dxo.Pagination

type DTO = PhoneNumberDTO

type VO = PhoneNumberVO

type Entity = PhoneNumberEntity

type ID = dxo.PhoneNumberID

type Query = PhoneNumberQuery

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

type PhoneNumberVO struct {
	dxo.VO

	Items []*PhoneNumberDTO `json:"phone_numbers"`
}

////////////////////////////////////////////////////////////////////////////////
// EOF
