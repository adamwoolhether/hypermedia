package contactsgrp

import (
	fe "github.com/adamwoolhether/hypermedia/app/hypermedia/frontend/view/contacts"
	"github.com/adamwoolhether/hypermedia/business/contacts"
)

func contactsToWeb(contacts []contacts.Contact) []fe.ContactWeb {
	views := make([]fe.ContactWeb, len(contacts))
	for i := range contacts {
		views[i] = contactToWeb(contacts[i])
	}

	return views
}

func contactToWeb(contact contacts.Contact) fe.ContactWeb {
	contactView := fe.ContactWeb{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
		Email:     contact.Email,
	}

	return contactView
}
