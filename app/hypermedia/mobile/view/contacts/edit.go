package contacts

import (
	"fmt"

	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/layout"
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xmlmodel"
)

func Edit(contact UpdateContact) xmlmodel.Doc {
	overrideHeader := xmlmodel.Header{
		Style: "header-buttons",
		Text: []xmlmodel.Text{
			{
				Style:   "header-button",
				Content: "Back",
				Behavior: &xmlmodel.Behavior{
					Trigger: "press",
					Action:  "back",
					Href:    "#",
				},
			},
		},
	}

	form := xmlmodel.Form{
		View: []xmlmodel.View{
			{
				ID:   "form-fields",
				View: []xmlmodel.View{FormFields(contact, false)},
			},
			{
				Style: "button",
				Behavior: &xmlmodel.Behavior{
					Trigger: "press",
					Action:  "replace-inner",
					Target:  "form-fields",
					Href:    fmt.Sprintf("/mobile/contacts/%d/edit", contact.ID),
					Verb:    "post",
				},
				Text: []xmlmodel.Text{
					{
						Style:   "button-label",
						Content: "Save",
					},
				},
			},
			{
				Style: "button",
				Behavior: &xmlmodel.Behavior{
					Trigger: "press",
					Action:  "reload",
					Href:    fmt.Sprintf("/mobile/contacts/%d", contact.ID),
				},
				Text: []xmlmodel.Text{
					{Style: "button-label", Content: "Cancel"},
				},
			},
		},
	}

	doc := layout.Layout()
	doc.Screen.Body.Header = overrideHeader
	doc.Screen.Body.View.Form = &form

	return doc
}
