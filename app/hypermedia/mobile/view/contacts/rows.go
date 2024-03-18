package contacts

import (
	"fmt"
	"strconv"

	"github.com/adamwoolhether/hypermedia/business/contacts"
)

func Rows(contacts []contacts.Contact, page int) Items {
	if len(contacts) == 0 {
		return Items{}
	}

	contactItems := make([]Item, len(contacts)+1)

	for i, contact := range contacts {
		var itemTextContent string
		switch {
		case len(contact.FirstName) > 0:
			itemTextContent = fmt.Sprintf("%s %s", contact.FirstName, contact.LastName)
		case len(contact.Phone) > 0:
			itemTextContent = contact.Phone
		case len(contact.Email) > 0:
			itemTextContent = contact.Email
		}

		contactItems[i] = Item{
			Key:   strconv.Itoa(contact.ID),
			Style: "contact-item",
			Text: &Text{
				Style:   "contact-item-label",
				Content: itemTextContent,
			},
			Behavior: &Behavior{
				Trigger: "press",
				Action:  "push",
				Href:    fmt.Sprintf("/mobile/contacts/%d", contact.ID),
			},
		}
	}

	if len(contacts) == 10 {
		contactItems[len(contacts)] = Item{
			ID:    "load-more",
			Key:   "load-more",
			Style: "load-more-item",
			Behavior: &Behavior{
				Trigger: "visible",
				Action:  "replace",
				Target:  "load-more",
				Href:    fmt.Sprintf("/mobile/contacts?rows_only=true&page=%d", page+1),
				Verb:    "get",
			},
			Spinner: &Spinner{},
		}
	}

	items := Items{
		Xmlns: "https://hyperview.org/hyperview",
		Item:  contactItems,
	}

	return items
}
