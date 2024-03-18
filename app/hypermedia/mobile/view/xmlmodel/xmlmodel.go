package xmlmodel

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
	Body   Body   `xml:"body,omitempty"`
}

type Styles struct {
	XMLName xml.Name `xml:"styles"`
	Style   []Style  `xml:"style,omitempty"`
}

type Style struct {
	ID                string `xml:"id,attr,omitempty"`
	AlignItems        string `xml:"alignItems,attr,omitempty"`
	BackgroundColor   string `xml:"backgroundColor,attr,omitempty"`
	BorderBottom      string `xml:"borderBottom,attr,omitempty"`
	BorderBottomWidth string `xml:"borderBottomWidth,attr,omitempty"`
	BorderRadius      string `xml:"borderRadius,attr,omitempty"`
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
	MarginTop         string `xml:"marginTop,attr,omitempty"`
	Padding           string `xml:"padding,attr,omitempty"`
	PaddingTop        string `xml:"paddingTop,attr,omitempty"`
	PaddingBottom     string `xml:"paddingBottom,attr,omitempty"`
	PaddingVertical   string `xml:"paddingVertical,attr,omitempty"`
	PaddingHorizontal string `xml:"paddingHorizontal,attr,omitempty"`
	MarginVertical    string `xml:"marginVertical,attr,omitempty"`
	TextAlign         string `xml:"textAlign,attr,omitempty"`
	Width             string `xml:"width,attr,omitempty"`
}

type Body struct {
	Style    string `xml:"style,attr,omitempty"`
	SafeArea bool   `xml:"safe-area,attr"`
	Header   Header `xml:"header,omitempty"`
	View     View   `xml:"view,omitempty"`
}

type Header struct {
	Text     []Text    `xml:"text,omitempty"`
	Behavior *Behavior `xml:"behavior,omitempty"`
}

type View struct {
	XMLName   xml.Name   `xml:"view"`
	Xmlns     string     `xml:"xmlns,attr,omitempty"` /////////
	ID        string     `xml:"id,attr,omitempty"`
	Style     string     `xml:"style,attr,omitempty"`
	Form      *Form      `xml:"form,omitempty"`
	Behavior  *Behavior  `xml:"behavior,omitempty"`
	TextField *TextField `xml:"text-field,omitempty"`
	Text      []Text     `xml:"text,omitempty"`
	View      []View     `xml:"view,omitempty"`
}

// INDEX ///////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Form struct {
	TextField *TextField `xml:"text-field,omitempty"`
	List      *List      `xml:"list,omitempty"`
	View      []View     `xml:"view,omitempty"`
}

type TextField struct {
	Name        string    `xml:"name,attr,omitempty"`
	Value       string    `xml:"value,attr,omitempty"`
	Placeholder string    `xml:"placeholder,attr,omitempty"`
	Style       string    `xml:"style,attr,omitempty"`
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

// SHOW ////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//type ShowContact struct {
//	Style string           `xml:"style,attr,omitempty"`
//	Text  Text             `xml:"text,omitempty"`
//	Sub   []SubShowContact `xml:"view,omitempty"`
//}

//type SubShowContact struct {
//	Style string `xml:"style,attr,omitempty"`
//	Text  []Text `xml:"text,omitempty"`
//}

// ROWS ////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Items struct {
	XMLName xml.Name `xml:"items"`
	Xmlns   string   `xml:"xmlns,attr,omitempty"`
	Item    []Item   `xml:"item,omitempty"`
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
	Style    string    `xml:"style,attr,omitempty"`
	Content  string    `xml:",chardata"`
	Behavior *Behavior `xml:"behavior,omitempty"`
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
