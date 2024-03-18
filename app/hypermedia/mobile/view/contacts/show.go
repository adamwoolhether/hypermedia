package contacts

import (
	"fmt"

	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/layout"
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xmlmodel"
	"github.com/adamwoolhether/hypermedia/business/contacts"
)

func Show(contact contacts.Contact) xmlmodel.Doc {
	contactView := xmlmodel.ShowContact{
		Style: "details",
		Text: xmlmodel.Text{
			Style:   "contact-name",
			Content: fmt.Sprintf("%s %s", contact.FirstName, contact.LastName),
		},
		Sub: []xmlmodel.SubShowContact{
			{
				Style: "contact-section",
				Text: []xmlmodel.Text{
					{Style: "contact-section-label", Content: "Phone"},
					{Style: "contact-section-info", Content: contact.Phone},
				},
			},
			{
				Style: "contact-section",
				Text: []xmlmodel.Text{
					{Style: "contact-section-label", Content: "Email"},
					{Style: "contact-section-info", Content: contact.Email},
				},
			},
		},
	}

	doc := layout.Layout(layout.WithShowContact(contactView))

	return doc
}
