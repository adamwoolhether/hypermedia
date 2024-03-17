package view

import "github.com/adamwoolhether/hypermedia/business/contacts"

func Index(contacts []contacts.Contact, page int) Form {
	form := Form{
		TextField: TextField{
			Name:        "q",
			Placeholder: "Search...",
			Style:       "search-field",
			Behavior: &Behavior{
				Trigger: "change",
				Action:  "replace-inner",
				Target:  "contacts-list",
				Href:    "/mobile/contacts?rows_only=true",
				Verb:    "get",
			},
		},
		List: List{
			ID: "contacts-list",
			Behavior: &Behavior{
				Trigger: "refresh",
				Action:  "replace-inner",
				Target:  "contacts-list",
				Href:    "/mobile/contacts?rows_only=true",
				Verb:    "get",
			},
			Items: Rows(contacts, page),
		},
	}

	return form
}
