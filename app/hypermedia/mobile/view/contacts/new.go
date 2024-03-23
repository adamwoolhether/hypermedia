package contacts

import (
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/layout"
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xmlmodel"
)

func New(contact UpdateContact) xmlmodel.Doc { // Use UpdateContact for ease of calling FormFields.
	header := xmlmodel.Header{
		Style: "buttons-row",
		Text: []xmlmodel.Text{
			{
				Style:   "header-button",
				Content: "Close",
				Behavior: &xmlmodel.Behavior{
					Trigger: "press",
					Action:  "close",
				},
			},
		},
	}

	form := xmlmodel.Form{
		View: []xmlmodel.View{
			{
				Style: "edit-fields",
				ID:    "form-fields",
				View:  []xmlmodel.View{FormFields(contact, false)},
			},
			{
				Style: "buttons-row",
				Text: []xmlmodel.Text{
					{
						Style:   "bottom-button-label",
						Content: "Add Contact",
						Behavior: &xmlmodel.Behavior{
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
