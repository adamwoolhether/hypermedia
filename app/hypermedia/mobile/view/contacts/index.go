package contacts

import "github.com/adamwoolhether/hypermedia/business/contacts"

func Index(contacts []contacts.Contact, page int) Doc {
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

	doc := Layout(WithIndex(form))

	return doc
}
