package contacts

import (
	"fmt"

	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/layout"
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xmlmodel"
)

func FormFields(contact UpdateContact, saved bool) xmlmodel.View {
	view := xmlmodel.View{
		Xmlns: layout.Namespace,
		Style: "edit-group",
		View: []xmlmodel.View{ // Maybe map is better to add errors by name instead of index.
			{
				Style: "edit-field",
				TextField: &xmlmodel.TextField{
					Style:       "edit-field-text",
					Name:        "first_name",
					Placeholder: "First name",
					Value:       contact.FirstName,
				},
			},
			{
				Style: "edit-field",
				TextField: &xmlmodel.TextField{
					Style:       "edit-field-text",
					Name:        "last_name",
					Placeholder: "Last name",
					Value:       contact.LastName,
				},
			},
			// We need to put email in its own view, to allow dynamic
			// updating of the error-field and replace the entire view.
			EmailView(contact),
			{
				Style: "edit-field",
				TextField: &xmlmodel.TextField{
					Style:       "edit-field-text",
					Name:        "phone",
					Placeholder: "Phone",
					Value:       contact.Phone,
				},
			},
		},
	}

	if contact.FieldErrs.FirstName != "" {
		fNameErr := xmlmodel.Text{
			Style:   "edit-field-error",
			Content: contact.FieldErrs.FirstName,
		}
		view.View[0].Text = append(view.View[0].Text, fNameErr)
	}

	if contact.FieldErrs.LastName != "" {
		lNameErr := xmlmodel.Text{
			Style:   "edit-field-error",
			Content: contact.FieldErrs.LastName,
		}
		view.View[1].Text = append(view.View[1].Text, lNameErr)
	}

	if contact.FieldErrs.Email != "" {
		emailErr := xmlmodel.Text{
			ID:      "edit-email-error",
			Style:   "edit-field-error",
			Content: contact.FieldErrs.Email,
		}
		view.View[2].Text = append(view.View[2].Text, emailErr)
	}

	if contact.FieldErrs.Phone != "" {
		phoneErr := xmlmodel.Text{
			ID:      "edit-email-error",
			Style:   "edit-field-error",
			Content: contact.FieldErrs.Phone,
		}
		view.View[3].Text = append(view.View[3].Text, phoneErr)
	}

	// Hyperview can't handle server-directed redirects.
	if saved {
		//view.Xmlns = "http://hyperview.org/hyperview"
		view.Behavior = []xmlmodel.Behavior{
			{
				Trigger:   "load",
				Action:    "dispatch-event",
				EventName: "contact-updated",
			},
			{
				Trigger: "load",
				Action:  "reload",
				Href:    fmt.Sprintf("/mobile/contacts/%d", contact.ID),
			},
		}
	}

	return view
}

func EmailView(contact UpdateContact) xmlmodel.View {
	emailField := xmlmodel.View{
		Xmlns: layout.Namespace,
		ID:    "email",
		Style: "edit-field",
		TextField: &xmlmodel.TextField{
			Style:       "edit-field-text",
			Name:        "email",
			Placeholder: "Email",
			Value:       contact.Email,
			Debounce:    "500",
			Behavior: &xmlmodel.Behavior{
				Trigger: "change",
				Action:  "replace",
				Target:  "email",
				Href:    fmt.Sprintf("/mobile/contacts/%d/email", contact.ID),
				Verb:    "get",
			},
		},
	}

	if contact.FieldErrs.Email != "" {
		emailErr := xmlmodel.Text{
			ID:      "edit-email-error",
			Style:   "edit-field-error",
			Content: contact.FieldErrs.Email,
		}
		emailField.Text = append(emailField.Text, emailErr)
	}

	return emailField
}
