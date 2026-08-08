package phonenumbers

import "context"

// PhoneNumberService ...
type PhoneNumberService interface {

	// edit

	Insert(c context.Context, o *PhoneNumberDTO) (*PhoneNumberDTO, error)
	Update(c context.Context, id PhoneNumberID, o *PhoneNumberDTO) (*PhoneNumberDTO, error)
	Delete(c context.Context, id PhoneNumberID) error

	// query

	Find(c context.Context, id PhoneNumberID) (*PhoneNumberDTO, error)
	List(c context.Context, q *PhoneNumberQuery) ([]*PhoneNumberDTO, error)
}
