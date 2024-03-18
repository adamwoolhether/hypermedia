package contactsgrp

import (
	fe "github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/contacts"
	"github.com/adamwoolhether/hypermedia/business/contacts"
)

func contactsToMobile(contacts []contacts.Contact) []fe.ContactMobile {
	views := make([]fe.ContactMobile, len(contacts))
	for i := range contacts {
		views[i] = contactToMobile(contacts[i])
	}

	return views
}

func contactToMobile(contact contacts.Contact) fe.ContactMobile {
	contactView := fe.ContactMobile{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
		Email:     contact.Email,
	}

	return contactView
}
