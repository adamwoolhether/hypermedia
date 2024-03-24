package contacts

import (
	"fmt"

	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/layout"
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xml"
)

func Edit(contact UpdateContact) xml.Doc {
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
		},
	}

	form := xml.Form{
		View: []xml.View{
			{
				ID:   "form-fields",
				View: []xml.View{FormFields(contact, false)},
			},
			{
				// HOW TO MAKE THIS SHOW AT THE BOTTOM OF THE SCREEN?
				Style: "buttons-row-bottom",
				View: []xml.View{
					{
						Behavior: []xml.Behavior{
							{
								Trigger: "press",
								Action:  "replace-inner",
								Target:  "form-fields",
								Href:    fmt.Sprintf("/mobile/contacts/%d/edit", contact.ID),
								Verb:    "post",
							},
						},
						Text: []xml.Text{
							{Style: "bottom-button-label", Content: "Save"},
						},
					},
					{
						Behavior: []xml.Behavior{
							{
								Trigger: "press",
								Action:  "reload",
								Href:    fmt.Sprintf("/mobile/contacts/%d", contact.ID),
							},
						},
						Text: []xml.Text{
							{Style: "bottom-button-label", Content: "Cancel"},
						},
					},
					{
						BehaviorWithAlertOpts: &xml.BehaviorAlertOpts{
							Behavior: xml.Behavior{
								XmlnsAlert:   xml.NamespaceAlert,
								Trigger:      "press",
								Action:       "alert",
								AlertTitle:   "Confirm delete",
								AlertMessage: fmt.Sprintf("Are you sure you want to delete %s?", contact.FirstName),
							},
							AlertOptions: []xml.AlertOption{
								{
									Label: "Confirm",
									Behavior: &xml.Behavior{
										Trigger: "press",
										Action:  "append",
										Target:  "form-fields",
										Href:    fmt.Sprintf("/mobile/contacts/%d/delete", contact.ID),
										Verb:    "post",
									},
								},
								{
									Label: "Cancel",
								},
							},
						},
						Text: []xml.Text{
							{Style: "bottom-button-label button-delete", Content: "Delete"},
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
