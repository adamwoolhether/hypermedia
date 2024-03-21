package contacts

import (
	"fmt"

	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xmlmodel"
)

func EditFields(contact UpdateContact, saved bool) xmlmodel.View {
	view := xmlmodel.View{
		Xmlns: "https://hyperview.org/hyperview",
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
					Debounce:    "250",
					Behavior: &xmlmodel.Behavior{
						Trigger: "change",
						Action:  "replace-inner",
						Target:  "edit-email-error",
						Href:    fmt.Sprintf("/mobile/contacts/%d/email", contact.ID),
						Verb:    "get",
					},
				},
				Text: []xmlmodel.Text{
					{
						ID:      "edit-email-error",
						Style:   "edit-field-error",
						Content: contact.FieldErrs.Email,
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
						Content: contact.FieldErrs.Phone,
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

func EmailValidationError(err string) xmlmodel.Text {
	return xmlmodel.Text{
		Xmlns:   "https://hyperview.org/hyperview",
		ID:      "edit-email-error",
		Style:   "edit-field-error",
		Content: err,
	}
}
