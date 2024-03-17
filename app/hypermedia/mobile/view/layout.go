package view

import "github.com/adamwoolhether/hypermedia/business/contacts"

func Layout(contacts []contacts.Contact, page int) Doc {
	doc := Doc{
		Xmlns: "https://hyperview.org/hyperview",
		Screen: Screen{
			Styles: Styles{
				Style: styles(),
			},

			Body: Body{
				Style:    "body",
				SafeArea: true,
				Header: Header{
					Text: Text{Style: "header-title", Content: "Contact.app"},
				},
				View: View{
					Style: "main",
					Form:  Index(contacts, page),
				},
			},
		},
	}

	return doc
}

func styles() []Style {
	styles := []Style{
		{
			ID:                "header-title",
			AlignItems:        "center",
			BorderBottomWidth: "1",
			BorderBottomColor: "#ccc",
			Display:           "flex",
			FontSize:          "24",
			JustifyContent:    "space-between",
			PaddingTop:        "10",
			PaddingBottom:     "10",
		},
		{
			ID:                "search-field",
			BackgroundColor:   "#E0E0E0",
			BorderBottomWidth: "1",
			BorderBottomColor: "#ccc",
			BorderTopColor:    "#ccc",
			BorderTopWidth:    "1",
			FontSize:          "16",
			Padding:           "10",
			Width:             "100%",
		},
		{
			ID:                "contact-item",
			BorderBottomWidth: "1",
			BorderBottomColor: "#ccc",
			Padding:           "0",
			Margin:            "0",
		},
		{
			ID:       "contact-item-label",
			FontSize: "18",
			Padding:  "10",
		},
	}

	return styles
}
