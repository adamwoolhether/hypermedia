package contacts

import "github.com/adamwoolhether/hypermedia/business/contacts"

type ContactMobile struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type ContactErrors struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type UpdateContact struct {
	ID             int           `json:"id" validate:"required"`
	FirstName      string        `json:"first_name" validate:"required"`
	LastName       string        `json:"last_name" validate:"required"`
	Phone          string        `json:"phone" validate:"required"`
	Email          string        `json:"email" validate:"required,email"`
	FieldErrs      ContactErrors `json:"field_errors"`
	InternalErrors string        `json:"internal_errors"`
}

func (uu UpdateContact) ToDB() contacts.Contact {
	c := contacts.Contact{
		ID:        uu.ID,
		FirstName: uu.FirstName,
		LastName:  uu.LastName,
		Phone:     uu.Phone,
		Email:     uu.Email,
	}

	return c
}
