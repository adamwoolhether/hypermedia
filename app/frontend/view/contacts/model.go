package contacts

import "github.com/adamwoolhether/hypermedia/business/contacts"

type NewContact struct {
	//ID     int    `json:"id"`
	FirstName      string           `json:"first_name" validate:"required"`
	LastName       string           `json:"last_name" validate:"required"`
	Phone          string           `json:"phone" validate:"required"`
	Email          string           `json:"email" validate:"required,email"`
	FieldErrs      NewContactErrors `json:"filed_errors"`
	InternalErrors string           `json:"internal_errors"`
}

func (nc NewContact) ToDB() contacts.Contact {
	c := contacts.Contact{
		//ID:    0, // TODO: get from db
		First: nc.FirstName,
		Last:  nc.LastName,
		Phone: nc.Phone,
		Email: nc.Email,
	}

	return c
}

type NewContactErrors struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}
