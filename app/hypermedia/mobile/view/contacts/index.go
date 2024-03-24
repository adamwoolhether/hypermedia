package contacts

import (
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/layout"
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xml"
)

func Index(contacts []ContactMobile, page int) xml.Doc {
	header := xml.Header{
		Style: "buttons-row",
		Text: []xml.Text{
			{
				Style:   "header-title",
				Content: "Contacts.app",
			},
			{
				Style:   "header-button",
				Content: "Add",
				Behavior: &xml.Behavior{
					Trigger: "press",
					Action:  "new",
					Href:    "/mobile/contacts/new",
				},
			},
		},
	}

	form := xml.Form{
		Behavior: []xml.Behavior{
			{
				Trigger:   "on-event",
				EventName: "contact-updated",
				Action:    "replace-inner",
				Target:    "contacts-list",
				Href:      "/mobile/contacts?rows_only=true",
				Verb:      "get",
			},
		},
		TextField: &xml.TextField{
			Name:        "q",
			Placeholder: "Search...",
			Style:       "search-field",
			Debounce:    "500",
			Behavior: &xml.Behavior{
				Trigger: "change",
				Action:  "replace-inner",
				Target:  "contacts-list",
				Href:    "/mobile/contacts?rows_only=true",
				Verb:    "get",
			},
		},
		List: &xml.List{
			ID: "contacts-list",
			Behavior: &xml.Behavior{
				Trigger: "refresh",
				Action:  "replace-inner",
				Target:  "contacts-list",
				Href:    "/mobile/contacts?rows_only=true",
				Verb:    "get",
			},
			Items: Rows(contacts, page),
		},
	}

	doc := layout.Layout()
	doc.Screen.Body.Header = header
	doc.Screen.Body.View.Form = &form

	return doc
}
