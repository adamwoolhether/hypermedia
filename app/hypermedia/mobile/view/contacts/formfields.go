package contacts

import (
	"fmt"
	"slices"

	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/layout"
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xml"
)

func FormFields(contact UpdateContact, saved bool, toasts ...string) xml.View {
	view := xml.View{
		Xmlns: xml.Namespace,
		Style: "edit-group",
		View: []xml.View{ // Maybe map is better to add errors by name instead of index.
			{
				Style: "edit-field",
				TextField: &xml.TextField{
					Style:       "edit-field-text",
					Name:        "first_name",
					Placeholder: "First name",
					Value:       contact.FirstName,
				},
			},
			{
				Style: "edit-field",
				TextField: &xml.TextField{
					Style:       "edit-field-text",
					Name:        "last_name",
					Placeholder: "Last name",
					Value:       contact.LastName,
				},
			},
			{
				Xmlns: xml.Namespace,
				Style: "edit-field",
				TextField: &xml.TextField{
					Style:       "edit-field-text",
					Name:        "email",
					Placeholder: "Email",
					Value:       contact.Email,
					Debounce:    "200",
					Behavior: &xml.Behavior{
						Trigger: "change",
						Action:  "replace",
						Target:  "edit-email-error",
						Href:    fmt.Sprintf("/mobile/contacts/%d/email", contact.ID),
						Verb:    "get",
					},
				},
				Text: []xml.Text{
					CheckEmailErr(contact),
				},
			},
			{
				Style: "edit-field",
				TextField: &xml.TextField{
					Style:       "edit-field-text",
					Name:        "phone",
					Placeholder: "Phone",
					Value:       contact.Phone,
				},
			},
		},
	}

	if contact.FieldErrs.FirstName != "" {
		fNameErr := xml.Text{
			Style:   "edit-field-error",
			Content: contact.FieldErrs.FirstName,
		}
		view.View[0].Text = append(view.View[0].Text, fNameErr)
	}

	if contact.FieldErrs.LastName != "" {
		lNameErr := xml.Text{
			Style:   "edit-field-error",
			Content: contact.FieldErrs.LastName,
		}
		view.View[1].Text = append(view.View[1].Text, lNameErr)
	}

	if contact.FieldErrs.Phone != "" {
		phoneErr := xml.Text{
			ID:      "edit-email-error",
			Style:   "edit-field-error",
			Content: contact.FieldErrs.Phone,
		}
		view.View[3].Text = append(view.View[3].Text, phoneErr)
	}

	// Hyperview can't handle server-directed redirects.
	if saved {
		toasts := layout.ShowToasts(toasts...)
		savedEvents := []xml.Behavior{
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

		view.Behavior = slices.Concat(toasts, savedEvents)
	}

	return view
}

func CheckEmailErr(contact UpdateContact) xml.Text {
	style := "hide"
	if contact.FieldErrs.Email != "" {
		style = "edit-field-error"
	}

	emailErr := xml.Text{
		Xmlns:   xml.Namespace,
		ID:      "edit-email-error",
		Style:   style,
		Content: contact.FieldErrs.Email,
	}

	return emailErr
}
