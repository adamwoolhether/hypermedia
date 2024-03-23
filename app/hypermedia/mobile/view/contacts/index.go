package contacts

import (
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/layout"
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xmlmodel"
)

func Index(contacts []ContactMobile, page int) xmlmodel.Doc {
	form := xmlmodel.Form{
		Behavior: []xmlmodel.Behavior{
			{
				Trigger:   "on-event",
				EventName: "contact-updated",
				Action:    "replace-inner",
				Target:    "contacts-list",
				Href:      "/mobile/contacts?rows_only=true",
				Verb:      "get",
			},
		},
		TextField: &xmlmodel.TextField{
			Name:        "q",
			Placeholder: "Search...",
			Style:       "search-field",
			Debounce:    "250",
			Behavior: &xmlmodel.Behavior{
				Trigger: "change",
				Action:  "replace-inner",
				Target:  "contacts-list",
				Href:    "/mobile/contacts?rows_only=true",
				Verb:    "get",
			},
		},
		List: &xmlmodel.List{
			ID: "contacts-list",
			Behavior: &xmlmodel.Behavior{
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
	doc.Screen.Body.View.Form = &form

	return doc
}
