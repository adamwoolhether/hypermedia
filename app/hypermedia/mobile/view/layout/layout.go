package layout

import (
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xml"
)

func Layout() xml.Doc {
	doc := xml.Doc{
		Xmlns: xml.Namespace,
		Screen: xml.Screen{
			Styles: xml.Styles{
				Style: styles(),
			},

			Body: xml.Body{
				Style:    "body",
				SafeArea: true,
				View: xml.View{
					Style: "main",
				},
			},
		},
	}

	return doc
}

func styles() []xml.Style {
	styles := []xml.Style{
		// Main
		{
			ID:   "body",
			Flex: "1",
		},
		{
			ID:              "main",
			Flex:            "1",
			BackgroundColor: "#eee",
		},
		{
			ID:         "header-title",
			FontSize:   "16",
			Color:      "black",
			FontWeight: "500",
			//PaddingHorizontal: "22", // To align with `contact-section`: MarginHorizontal+Padding
		},
		{
			ID:                "buttons-row",
			BackgroundColor:   "white",
			BorderBottomColor: "#ccc",
			BorderBottomWidth: "1",
			FlexDirection:     "row",
			Height:            "50",
			AlignItems:        "center",
			JustifyContent:    "space-between",
			PaddingHorizontal: "22", // To align with `contact-section`: MarginHorizontal+Padding
		},
		{
			ID:                "buttons-row-bottom",
			Bottom:            "0",
			BackgroundColor:   "white",
			BorderBottomColor: "#ccc",
			BorderBottomWidth: "1",
			FlexDirection:     "row",
			Height:            "50",
			AlignItems:        "center",
			JustifyContent:    "space-between",
			PaddingHorizontal: "22", // To align with `contact-section`: MarginHorizontal+Padding
		},

		// Index
		{
			ID:                "header",
			FlexDirection:     "row",
			JustifyContent:    "space-between",
			AlignItems:        "center",
			BorderBottomColor: "#ccc",
			BorderBottomWidth: "1",
			PaddingLeft:       "24",
			PaddingRight:      "24",
			PaddingVertical:   "16",
			BackgroundColor:   "white",
		},
		{
			ID:                "search-field",
			PaddingHorizontal: "24",
			PaddingVertical:   "8",
			BorderBottomWidth: "1",
			BorderBottomColor: "#ddd",
			BackgroundColor:   "#eee",
		},
		{
			ID:                "contact-item",
			BorderBottomColor: "#ddd",
			BorderBottomWidth: "1",
			PaddingLeft:       "24",
			PaddingRight:      "24",
			PaddingVertical:   "16",
			BackgroundColor:   "white",
		},
		{
			ID:         "contact-item-label",
			FontWeight: "500",
		},
		{
			ID:              "load-more-item",
			PaddingVertical: "16",
		},

		// Show
		{
			ID:       "header-button",
			FontSize: "20",
			Color:    "blue",
		},
		{
			ID:             "contact-name",
			FontSize:       "24",
			TextAlign:      "center",
			MarginVertical: "32",
			FontWeight:     "500",
		},
		{
			ID:               "contact-section",
			Margin:           "8",
			BackgroundColor:  "white",
			BorderRadius:     "8",
			Padding:          "8",
			MarginHorizontal: "14",
		},
		{
			ID:           "contact-section-label",
			FontSize:     "12",
			Color:        "#aaa",
			MarginBottom: "4",
		},
		{
			ID:           "contact-section-info",
			FontSize:     "18",
			Color:        "blue",
			MarginBottom: "4",
		},

		// Edit
		{
			ID:             "edit-group",
			MarginVertical: "8",
		},
		{
			ID:                "edit-field",
			TextAlign:         "center",
			BorderBottomWidth: "1",
			BorderColor:       "#ddd",
			PaddingHorizontal: "24",
			PaddingVertical:   "16",
			BackgroundColor:   "white",
		},
		{
			ID:        "edit-field-error",
			Color:     "red",
			FontSize:  "12",
			MarginTop: "4",
		},
		{
			ID:         "bottom-button-label",
			Color:      "blue",
			FontWeight: "500",
		},
		{
			ID:         "button-delete",
			Color:      "red",
			FontWeight: "500",
		},
		{
			ID:             "swipe-button",
			Height:         "100%",
			JustifyContent: "center",
			TextAlign:      "center",
			PaddingLeft:    "24",
		},
	}

	return styles
}
