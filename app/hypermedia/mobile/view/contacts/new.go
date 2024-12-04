package contacts

import (
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/layout"
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xml"
)

func New(contact UpdateContact) xml.Doc { // Use UpdateContact for ease of calling FormFields.
	header := xml.Header{
		Style: "buttons-row",
		Text: []xml.Text{
			{
				Style:   "header-button",
				Content: "Close",
				Behavior: &xml.Behavior{
					Trigger: "press",
					Action:  "close",
				},
			},
		},
	}

	form := xml.Form{
		View: []xml.View{
			{
				Style: "edit-fields",
				ID:    "form-fields",
				View:  []xml.View{FormFields(contact, false)},
			},
			{
				Style: "buttons-row",
				Text: []xml.Text{
					{
						Style:   "bottom-button-label",
						Content: "Add Contact",
						Behavior: &xml.Behavior{
							Trigger: "press",
							Action:  "replace-inner",
							Target:  "form-fields",
							Href:    "/mobile/contacts/new",
							Verb:    "post",
						},
					},
				},
			},
		},
	}

	doc := layout.Layout()
	doc.Screen.Body.Header = header
	doc.Screen.Body.View.Form = &form

	return doc
}
