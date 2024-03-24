package contacts

import (
	"fmt"

	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/layout"
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xml"
)

func Show(contact ContactMobile) xml.Doc {
	overrideHeader := xml.Header{
		Style: "buttons-row",
		Text: []xml.Text{
			{
				Style:   "header-button",
				Content: "Back",
				Behavior: &xml.Behavior{
					Trigger: "press",
					Action:  "back",
				},
			},
			{
				Style:   "header-button",
				Content: "Edit",
				Behavior: &xml.Behavior{
					Trigger: "press",
					Action:  "reload",
					Href:    fmt.Sprintf("/mobile/contacts/%d/edit", contact.ID),
				},
			},
		},
	}

	showContact := xml.View{
		Style: "details",
		Text: []xml.Text{{
			Style:   "contact-name",
			Content: fmt.Sprintf("%s %s", contact.FirstName, contact.LastName),
		},
		},
		View: []xml.View{
			{
				Style: "contact-section",
				Text: []xml.Text{
					{Style: "contact-section-label", Content: "Phone"},
					{Style: "contact-section-info", Content: contact.Phone},
				},
			},
			{
				Style: "contact-section",
				Text: []xml.Text{
					{Style: "contact-section-label", Content: "Email"},
					{Style: "contact-section-info", Content: contact.Email},
				},
			},
		},
	}

	doc := layout.Layout()
	doc.Screen.Body.Header = overrideHeader
	doc.Screen.Body.View.View = []xml.View{showContact}

	return doc
}
