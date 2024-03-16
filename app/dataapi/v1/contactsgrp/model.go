package contactsgrp

import (
	"github.com/adamwoolhether/hypermedia/business/contacts"
	"github.com/adamwoolhether/hypermedia/foundation/validate"
)

type NewContact struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

func (nc NewContact) ToDB() contacts.Contact {
	c := contacts.Contact{
		FirstName: nc.FirstName,
		LastName:  nc.LastName,
		Phone:     nc.Phone,
		Email:     nc.Email,
	}

	return c
}

func (nc NewContact) Validate() error {
	if err := validate.Check(nc); err != nil {
		return err
	}

	return nil
}

type ContactAPI struct {
	ID        int    `json:"id"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type Response struct {
	Contacts []ContactAPI `json:"contacts"`
	Page     int          `json:"page"`
	Pages    int          `json:"pages"`
	Total    int          `json:"total"`
}

func newResponse(contacts []contacts.Contact, total, page, rows int) Response {
	pages := total / rows
	if total%rows != 0 {
		pages += 1
	}

	response := Response{
		Contacts: contactsToAPI(contacts),
		Page:     page,
		Pages:    pages,
		Total:    total,
	}

	return response
}

func contactsToAPI(contacts []contacts.Contact) []ContactAPI {
	views := make([]ContactAPI, len(contacts))
	for i := range contacts {
		views[i] = contactToAPI(contacts[i])
	}

	return views
}

func contactToAPI(contact contacts.Contact) ContactAPI {
	contactView := ContactAPI{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
		Email:     contact.Email,
	}

	return contactView
}
