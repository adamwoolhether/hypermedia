package contactsgrp

import (
	fe "github.com/adamwoolhether/hypermedia/app/hypermedia/frontend/view/contacts"
	"github.com/adamwoolhether/hypermedia/business/contacts"
)

func contactsToView(contacts []contacts.Contact) []fe.ContactView {
	views := make([]fe.ContactView, len(contacts))
	for i := range contacts {
		views[i] = contactToView(contacts[i])
	}

	return views
}

func contactToView(contact contacts.Contact) fe.ContactView {
	contactView := fe.ContactView{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
		Email:     contact.Email,
	}

	return contactView
}
