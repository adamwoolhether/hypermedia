package layout

import (
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xmlmodel"
)

func Layout() xmlmodel.Doc {
	doc := xmlmodel.Doc{
		Xmlns: "https://hyperview.org/hyperview",
		Screen: xmlmodel.Screen{
			Styles: xmlmodel.Styles{
				Style: styles(),
			},

			Body: xmlmodel.Body{
				Style:    "body",
				SafeArea: true,
				Header: xmlmodel.Header{
					Text: []xmlmodel.Text{
						{Style: "header-title", Content: "Contact.app"},
					},
				},
				View: xmlmodel.View{
					Style: "main",
				},
			},
		},
	}

	return doc
}

func styles() []xmlmodel.Style {
	styles := []xmlmodel.Style{
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
			ID:                "header-button",
			AlignItems:        "center",
			BackgroundColor:   "#FFFFFF",
			BorderBottomColor: "#ccc",
			BorderBottomWidth: "1",
			Color:             "blue",
			Display:           "flex",
			FontSize:          "24",
			JustifyContent:    "space-between",
			Padding:           "20",
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
			BorderBottomColor: "#ccc",
			BorderBottomWidth: "1",
			Padding:           "0",
			Margin:            "0",
		},
		{
			ID:       "contact-item-label",
			FontSize: "18",
			Padding:  "10",
		},
		{
			ID:         "contact-name",
			FontSize:   "32",
			FontWeight: "bold",
			Margin:     "20",
			TextAlign:  "center",
		},
		{
			ID:                "contact-section",
			BackgroundColor:   "white",
			BorderBottomColor: "#ccc",
			BorderRadius:      "10",
			Padding:           "8",
			Margin:            "16",
		},
		{
			ID:                "contact-section-label",
			Color:             "#D3D3D3",
			FontSize:          "16",
			FontWeight:        "bold",
			PaddingHorizontal: "12",
			MarginTop:         "6",
		},
		{
			ID:                "contact-section-info",
			BackgroundColor:   "#fff",
			BorderRadius:      "8",
			Color:             "blue",
			FontSize:          "18",
			PaddingHorizontal: "12",
			PaddingVertical:   "8",
		},
	}

	return styles
}
