package contacts

import (
	"fmt"

	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/layout"
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xmlmodel"
)

func Show(contact ContactMobile) xmlmodel.Doc {
	overrideHeader := xmlmodel.Header{
		Style: "buttons-row",
		Text: []xmlmodel.Text{
			{
				Style:   "header-button",
				Content: "Back",
				Behavior: &xmlmodel.Behavior{
					Trigger: "press",
					Action:  "back",
				},
			},
			{
				Style:   "header-button",
				Content: "Edit",
				Behavior: &xmlmodel.Behavior{
					Trigger: "press",
					Action:  "reload",
					Href:    fmt.Sprintf("/mobile/contacts/%d/edit", contact.ID),
				},
			},
		},
	}

	showContact := xmlmodel.View{
		Style: "details",
		Text: []xmlmodel.Text{{
			Style:   "contact-name",
			Content: fmt.Sprintf("%s %s", contact.FirstName, contact.LastName),
		},
		},
		View: []xmlmodel.View{
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

	doc := layout.Layout()
	doc.Screen.Body.Header = overrideHeader
	doc.Screen.Body.View.View = []xmlmodel.View{showContact}

	return doc
}
