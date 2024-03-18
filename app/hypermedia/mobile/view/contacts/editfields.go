package contacts

import (
	"fmt"

	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xmlmodel"
)

func EditFields(contact UpdateContact, saved bool) xmlmodel.View {
	view := xmlmodel.View{
		Style: "edit-group",
		View: []xmlmodel.View{
			{
				Style: "edit-field",
				TextField: &xmlmodel.TextField{
					Name:        "first_name",
					Placeholder: "First name",
					Value:       contact.FirstName,
				},
				Text: []xmlmodel.Text{
					{
						Style:   "edit-field-error",
						Content: contact.FieldErrs.FirstName,
					},
				},
			},
			{
				Style: "edit-field",
				TextField: &xmlmodel.TextField{
					Name:        "last_name",
					Placeholder: "Last name",
					Value:       contact.LastName,
				},
				Text: []xmlmodel.Text{
					{
						Style:   "edit-field-error",
						Content: contact.FieldErrs.LastName,
					},
				},
			},
			{
				Style: "edit-field",
				TextField: &xmlmodel.TextField{
					Name:        "email",
					Placeholder: "Email",
					Value:       contact.Email,
					//Behavior: &xmlmodel.Behavior{
					//	Trigger: "change",
					//	Action:  "replace-inner",
					//	Target:  "edit-field-error",
					//	Href:    fmt.Sprintf("/contacts/%d/email", contact.ID),
					//	Verb:    "get",
					//},
				},
				Text: []xmlmodel.Text{
					{
						Style:   "edit-field-error",
						Content: contact.FieldErrs.LastName,
					},
				},
			},
			{
				Style: "edit-field",
				TextField: &xmlmodel.TextField{
					Name:        "phone",
					Placeholder: "Phone",
					Value:       contact.Phone,
				},
				Text: []xmlmodel.Text{
					{
						Style:   "edit-field-error",
						Content: contact.FieldErrs.LastName,
					},
				},
			},
		},
	}

	// Hyperview can't handle server-directed redirects.
	if saved {
		//view.Xmlns = "http://hyperview.org/hyperview"
		view.Behavior = &xmlmodel.Behavior{
			Trigger: "load",
			Action:  "reload",
			Href:    fmt.Sprintf("/mobile/contacts/%d", contact.ID),
		}
	}

	return view
}
