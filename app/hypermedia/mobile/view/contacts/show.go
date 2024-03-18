package contacts

import (
	"fmt"

	"github.com/adamwoolhether/hypermedia/business/contacts"
)

func Show(contact contacts.Contact) Doc {
	contactView := ShowContact{
		Style: "details",
		Text: Text{
			Style:   "contact-name",
			Content: fmt.Sprintf("%s %s", contact.FirstName, contact.LastName),
		},
		Sub: []SubShowContact{
			{
				Style: "contact-section",
				Text: []Text{
					{Style: "contact-section-label", Content: "Phone"},
					{Style: "contact-section-info", Content: contact.Phone},
				},
			},
			{
				Style: "contact-section",
				Text: []Text{
					{Style: "contact-section-label", Content: "Email"},
					{Style: "contact-section-info", Content: contact.Email},
				},
			},
		},
	}

	doc := Layout(WithShowContact(contactView))

	return doc
}
