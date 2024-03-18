package contacts

import (
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/layout"
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xmlmodel"
	"github.com/adamwoolhether/hypermedia/business/contacts"
)

func Index(contacts []contacts.Contact, page int) xmlmodel.Doc {
	form := xmlmodel.Form{
		TextField: xmlmodel.TextField{
			Name:        "q",
			Placeholder: "Search...",
			Style:       "search-field",
			Behavior: &xmlmodel.Behavior{
				Trigger: "change",
				Action:  "replace-inner",
				Target:  "contacts-list",
				Href:    "/mobile/contacts?rows_only=true",
				Verb:    "get",
			},
		},
		List: xmlmodel.List{
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

	doc := layout.Layout(layout.WithIndex(form))

	return doc
}
