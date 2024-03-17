package view

import (
	"encoding/xml"
)

// LAYOUT //////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Doc struct {
	XMLName xml.Name `xml:"doc"`
	Xmlns   string   `xml:"xmlns,attr"`
	Screen  Screen   `xml:"screen"`
}

type Screen struct {
	Styles Styles `xml:"styles"`
	Body   Body   `xml:"body"`
}

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

// INDEX ///////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Form struct {
	TextField TextField `xml:"text-field"`
	List      List      `xml:"list"`
}

type TextField struct {
	Name        string    `xml:"name,attr"`
	Value       string    `xml:"value,attr"`
	Placeholder string    `xml:"placeholder,attr"`
	Style       string    `xml:"style,attr"`
	Behavior    *Behavior `xml:"behavior,omitempty"`
}

type List struct {
	ID    string `xml:"id,attr"`
	Items Items  `xml:"items"`
	// Behavior is embedded here, adding the behavior attributes directly
	// to the list element, rather than adding the element to the list.
	// This is a convenient shorthand. We do this when  because of the use
	// of `replace-inner`. This replaces all child elements of the target
	// with the new content.
	*Behavior
}

// ROWS ////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Items struct {
	XMLName xml.Name `xml:"items"`
	Xmlns   string   `xml:"xmlns,attr"`
	Item    []Item   `xml:"item"`
}

type Item struct {
	ID       string    `xml:"id,attr,omitempty"`
	Key      string    `xml:"key,attr,omitempty"`
	Style    string    `xml:"style,attr,omitempty"`
	Text     *Text     `xml:"text,omitempty"`
	Behavior *Behavior `xml:"behavior,omitempty"`
	Spinner  *Spinner  `xml:"spinner,omitempty"`
}

// SHARED //////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Text struct {
	Style   string `xml:"style,attr"`
	Content string `xml:",chardata"`
}

type Behavior struct {
	Trigger string `xml:"trigger,attr,omitempty"`
	Action  string `xml:"action,attr,omitempty"`
	Target  string `xml:"target,attr,omitempty"`
	Href    string `xml:"href,attr,omitempty"`
	Verb    string `xml:"verb,attr,omitempty"`
}

type Spinner struct {
	XMLName xml.Name `xml:"spinner"`
}
