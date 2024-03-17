package view

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/adamwoolhether/hypermedia/business/contacts"
)

// LAYOUT /////////////////////////////////////////
type Doc struct {
	XMLName xml.Name `xml:"doc"`
	Xmlns   string   `xml:"xmlns,attr"`
	Screen  Screen   `xml:"screen"`
}

type Screen struct {
	Styles Styles `xml:"styles"`
	Body   Body   `xml:"body"`
}

// Styles struct is left empty since the styles are omitted for brevity in the template.
type Styles struct {
	XMLName xml.Name `xml:"styles"`
	Style   []Style  `xml:"style"`
}

type Style struct {
	ID                string `xml:"id,attr"`
	AlignItems        string `xml:"alignItems,attr,omitempty"`
	BackgroundColor   string `xml:"backgroundColor,attr,omitempty"`
	BorderBottom      string `xml:"borderBottom,attr,omitempty"`
	BorderBottomWidth string `xml:"borderBottomWidth,attr,omitempty"`
	BorderTopColor    string `xml:"borderTopColor,attr,omitempty"`
	BorderTopWidth    string `xml:"borderTopWidth,attr,omitempty"`
	BorderBottomColor string `xml:"borderBottomColor,attr,omitempty"`
	Color             string `xml:"color,attr,omitempty"`
	Display           string `xml:"display,attr,omitempty"`
	Flex              string `xml:"flex,attr,omitempty"`
	FlexDirection     string `xml:"flexDirection,attr,omitempty"`
	FontSize          string `xml:"fontSize,attr,omitempty"`
	FontWeight        string `xml:"fontWeight,attr,omitempty"`
	JustifyContent    string `xml:"justifyContent,attr,omitempty"`
	Margin            string `xml:"margin,attr,omitempty"`
	Padding           string `xml:"padding,attr,omitempty"`
	PaddingTop        string `xml:"paddingTop,attr,omitempty"`
	PaddingBottom     string `xml:"paddingBottom,attr,omitempty"`
	Width             string `xml:"width,attr,omitempty"`
}

type Body struct {
	Style    string `xml:"style,attr"`
	SafeArea bool   `xml:"safe-area,attr"`
	Header   Header `xml:"header"`
	View     View   `xml:"view"`
}

type Header struct {
	Text Text `xml:"text"`
}

type View struct {
	Style string `xml:"style,attr"`
	Form  Form   `xml:"form"`
}

// INDEX /////////////////////////////////////////
type Form struct {
	TextField TextField `xml:"text-field"`
	List      List      `xml:"list"`
}

type TextField struct {
	Name        string `xml:"name,attr"`
	Value       string `xml:"value,attr"`
	Placeholder string `xml:"placeholder,attr"`
	Style       string `xml:"style,attr"`
}

type List struct {
	ID    string `xml:"id,attr"`
	Items Items  `xml:"items"`
}

// ROWS /////////////////////////////////////////
type Items struct {
	XMLName xml.Name `xml:"items"`
	Xmlns   string   `xml:"xmlns,attr"`
	Item    []Item   `xml:"item"`
}

type Item struct {
	Key   string `xml:"key,attr"`
	Style string `xml:"style,attr"`
	Text  Text   `xml:"text"`
}

type Text struct {
	Style   string `xml:"style,attr"`
	Content string `xml:",chardata"`
}

func Layout(contacts []contacts.Contact) Doc {
	doc := Doc{
		Xmlns: "https://hyperview.org/hyperview",
		Screen: Screen{
			Styles: Styles{
				Style: stylesTemplate(),
			},

			Body: Body{
				Style:    "body",
				SafeArea: true,
				Header: Header{
					Text: Text{Style: "header-title", Content: "Contact.app"},
				},
				View: View{
					Style: "main",
					Form: Form{
						TextField: TextField{
							Name:        "q",
							Value:       "",
							Placeholder: "Search...",
							Style:       "search-field",
						},
						List: List{
							ID:    "contacts-list",
							Items: RowsTemplate(contacts),
						},
					},
				},
			},
		},
	}

	return doc
}

func stylesTemplate() []Style {
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

func RowsTemplate(contacts []contacts.Contact) Items {
	contactItems := make([]Item, len(contacts))

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
			Text: Text{
				Style:   "contact-item-label",
				Content: itemTextContent,
			},
		}
	}

	items := Items{
		Xmlns: "https://hyperview.org/hyperview",
		Item:  contactItems,
	}

	return items
}
