package contacts

// Options represents optional parameters.
type Options struct {
	index *Form
	show  *ShowContact
}

func WithIndex(indexDoc Form) func(opts *Options) {
	return func(opts *Options) {
		opts.index = &indexDoc
	}
}

func WithShowContact(contact ShowContact) func(opts *Options) {
	return func(opts *Options) {
		opts.show = &contact
	}
}

type LayoutOpts func(opts *Options)

func Layout(opts ...LayoutOpts) Doc {
	var options Options
	for _, opt := range opts {
		opt(&options)
	}

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
				},
			},
		},
	}

	if options.index != nil {
		doc.Screen.Body.View.Index = options.index
	}

	if options.show != nil {
		header := Header{
			Text: Text{Style: "header-button", Content: "Back"},
			Behavior: &Behavior{
				Trigger: "press",
				Action:  "back",
			},
		}

		doc.Screen.Body.Header = header
		doc.Screen.Body.View.Show = options.show
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
