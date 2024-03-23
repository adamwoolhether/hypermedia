package contacts

import (
	"fmt"

	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/layout"
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xmlmodel"
)

func Edit(contact UpdateContact) xmlmodel.Doc {
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
		},
	}

	form := xmlmodel.Form{
		View: []xmlmodel.View{
			{
				ID:   "form-fields",
				View: []xmlmodel.View{FormFields(contact, false)},
			},
			{
				// HOW TO MAKE THIS SHOW AT THE BOTTOM OF THE SCREEN?
				Style: "buttons-row-bottom",
				View: []xmlmodel.View{
					{
						Behavior: &xmlmodel.Behavior{
							Trigger: "press",
							Action:  "replace-inner",
							Target:  "form-fields",
							Href:    fmt.Sprintf("/mobile/contacts/%d/edit", contact.ID),
							Verb:    "post",
						},
						Text: []xmlmodel.Text{
							{Style: "bottom-button-label", Content: "Save"},
						},
					},
					{
						Behavior: &xmlmodel.Behavior{
							Trigger: "press",
							Action:  "reload",
							Href:    fmt.Sprintf("/mobile/contacts/%d", contact.ID),
						},
						Text: []xmlmodel.Text{
							{Style: "bottom-button-label", Content: "Cancel"},
						},
					},
				},
			},
		},
	}

	doc := layout.Layout()
	doc.Screen.Body.Header = overrideHeader
	doc.Screen.Body.View.Form = &form

	return doc
}
