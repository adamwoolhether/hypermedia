package contacts

import (
	"fmt"
	"strconv"

	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xml"
)

func Rows(contacts []ContactMobile, page int) xml.Items {
	if len(contacts) == 0 {
		return xml.Items{}
	}

	contactItems := make([]xml.Item, len(contacts)+1)

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

		contactItems[i] = xml.Item{
			Key: strconv.Itoa(contact.ID),
			SwipeRow: &xml.SwipeRowParams{
				SwipeRow: xml.SwipeRow{
					XmlnsSwipe: xml.NamespaceSwipe,
				},
				SwipeMain: xml.SwipeMainParams{
					View: &xml.View{
						Style: "contact-item",
						Text: []xml.Text{
							{Style: "contact-item-label", Content: itemTextContent},
						},
						Behavior: []xml.Behavior{
							{
								Trigger: "press",
								Action:  "push",
								Href:    fmt.Sprintf("/mobile/contacts/%d", contact.ID),
							},
						},
					},
				},
				SwipeButtons: []xml.SwipeButton{
					{
						View: &xml.View{
							Style: "swipe-button",
							Behavior: []xml.Behavior{
								{
									Trigger: "press",
									Action:  "push",
									Href:    fmt.Sprintf("/mobile/contacts/%d/edit", contact.ID),
								},
							},
							Text: []xml.Text{
								{Style: "bottom-button-label", Content: "Edit"},
							},
						},
					},
					{
						View: &xml.View{
							Style: "swipe-button",
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
											Target:  fmt.Sprintf("item-%d", contact.ID),
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
								{Style: "button-delete", Content: "Delete"},
							},
						},
					},
				},
			},
		}
	}

	if len(contacts) == 10 {
		contactItems[len(contacts)] = xml.Item{
			ID:    "load-more",
			Key:   "load-more",
			Style: "load-more-item",
			Behavior: &xml.Behavior{
				Trigger: "visible",
				Action:  "replace",
				Target:  "load-more",
				Href:    fmt.Sprintf("/mobile/contacts?rows_only=true&page=%d", page+1),
				Verb:    "get",
			},
			Spinner: &xml.Spinner{},
		}
	}

	items := xml.Items{
		Xmlns: xml.Namespace,
		Item:  contactItems,
	}

	return items
}
