package xml

import (
	"encoding/xml"
)

type NSSpace string

const (
	Namespace      = "https://hyperview.org/hyperview"
	NamespaceAlert = "https://hyperview.org/hyperview-alert"
	NamespaceComms = "https://hypermedia.systems/hyperview/communications"
	NamespaceToast = "https://hypermedia.systems/hyperview/toast"
	NamespaceSwipe = "https://hypermedia.systems/hyperview/swipeable"
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
	Bottom            string `xml:"bottom,attr,omitempty"`
	BorderBottom      string `xml:"borderBottom,attr,omitempty"`
	BorderBottomWidth string `xml:"borderBottomWidth,attr,omitempty"`
	BorderRadius      string `xml:"borderRadius,attr,omitempty"`
	BorderColor       string `xml:"borderColor,attr,omitempty"`
	BorderTopColor    string `xml:"borderTopColor,attr,omitempty"`
	BorderTopWidth    string `xml:"borderTopWidth,attr,omitempty"`
	BorderBottomColor string `xml:"borderBottomColor,attr,omitempty"`
	Color             string `xml:"color,attr,omitempty"`
	Display           string `xml:"display,attr,omitempty"`
	Flex              string `xml:"flex,attr,omitempty"`
	FlexDirection     string `xml:"flexDirection,attr,omitempty"`
	FontSize          string `xml:"fontSize,attr,omitempty"`
	FontWeight        string `xml:"fontWeight,attr,omitempty"`
	Height            string `xml:"height,attr,omitempty"`
	JustifyContent    string `xml:"justifyContent,attr,omitempty"`
	Left              string `xml:"left,attr,omitempty"`
	Margin            string `xml:"margin,attr,omitempty"`
	MarginTop         string `xml:"marginTop,attr,omitempty"`
	MarginBottom      string `xml:"marginBottom,attr,omitempty"`
	MarginHorizontal  string `xml:"marginHorizontal,attr,omitempty"`
	MarginVertical    string `xml:"marginVertical,attr,omitempty"`
	Padding           string `xml:"padding,attr,omitempty"`
	PaddingTop        string `xml:"paddingTop,attr,omitempty"`
	PaddingBottom     string `xml:"paddingBottom,attr,omitempty"`
	PaddingLeft       string `xml:"paddingLeft,attr,omitempty"`
	PaddingRight      string `xml:"paddingRight,attr,omitempty"`
	PaddingHorizontal string `xml:"paddingHorizontal,attr,omitempty"`
	PaddingVertical   string `xml:"paddingVertical,attr,omitempty"`
	Position          string `xml:"position,attr,omitempty"`
	Right             string `xml:"right,attr,omitempty"`
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
	Style    string    `xml:"style,attr,omitempty"`
	Text     []Text    `xml:"text,omitempty"`
	Behavior *Behavior `xml:",omitempty"`
}

type View struct {
	XMLName               xml.Name           `xml:"view"`
	Xmlns                 string             `xml:"xmlns,attr,omitempty"` /////////
	ID                    string             `xml:"id,attr,omitempty"`
	Style                 string             `xml:"style,attr,omitempty"`
	Form                  *Form              `xml:"form,omitempty"`
	Behavior              []Behavior         `xml:",omitempty"`
	BehaviorWithAlertOpts *BehaviorAlertOpts `xml:",omitempty"`
	TextField             *TextField         `xml:"text-field,omitempty"`
	Text                  []Text             `xml:"text,omitempty"`
	View                  []View             `xml:"view,omitempty"`
}

// INDEX ///////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Form struct {
	TextField *TextField `xml:"text-field,omitempty"`
	List      *List      `xml:"list,omitempty"`
	View      []View     `xml:"view,omitempty"`
	Behavior  []Behavior `xml:",omitempty"`
}

type TextField struct {
	Name        string    `xml:"name,attr,omitempty"`
	Value       string    `xml:"value,attr,omitempty"`
	Placeholder string    `xml:"placeholder,attr,omitempty"`
	Style       string    `xml:"style,attr,omitempty"`
	Behavior    *Behavior `xml:",omitempty"`
	Debounce    string    `xml:"debounce,attr,omitempty"`
}

type List struct {
	XMLName xml.Name `xml:"list"`
	ID      string   `xml:"id,attr"`
	Items   Items    `xml:"items"`
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
	ID       string          `xml:"id,attr,omitempty"`
	Key      string          `xml:"key,attr,omitempty"`
	Style    string          `xml:"style,attr,omitempty"`
	Text     *Text           `xml:"text,omitempty"`
	Behavior *Behavior       `xml:",omitempty"`
	Spinner  *Spinner        `xml:"spinner,omitempty"`
	SwipeRow *SwipeRowParams `xml:",omitempty"`
}

// SHARED //////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Text struct {
	XMLName  xml.Name  `xml:"text"`
	Xmlns    string    `xml:"xmlns,attr,omitempty"`
	ID       string    `xml:"id,attr,omitempty"`
	Style    string    `xml:"style,attr,omitempty"`
	Content  string    `xml:",chardata"`
	Behavior *Behavior `xml:",omitempty"`
	Debounce string    `xml:"debounce,attr,omitempty"`
}

type Behavior struct {
	XMLName   xml.Name `xml:"behavior"`
	Trigger   string   `xml:"trigger,attr,omitempty"`
	Action    string   `xml:"action,attr,omitempty"`
	Target    string   `xml:"target,attr,omitempty"`
	Href      string   `xml:"href,attr,omitempty"`
	Verb      string   `xml:"verb,attr,omitempty"`
	EventName string   `xml:"event-name,attr,omitempty"`
	// Alert
	XmlnsAlert   string `xml:"xmlns:alert,attr,omitempty"`
	AlertTitle   string `xml:"alert:title,attr,omitempty"`
	AlertMessage string `xml:"alert:message,attr,omitempty"`
	// Comms
	XmlnsComms       string `xml:"xmlns:comms,attr,omitempty"`
	CommsPhoneNumber string `xml:"comms:phone-number,attr,omitempty"`
	CommsEmailAddr   string `xml:"comms:email-address,attr,omitempty"`
	// Toast
	XmlnsToast string `xml:"xmlns:toast,attr,omitempty"`
	ToastText  string `xml:"toast:text,attr,omitempty"`
}

type BehaviorAlertOpts struct {
	Behavior     `xml:",any"`
	AlertOptions []AlertOption `xml:"alert:option,omitempty"`
}

type AlertOption struct {
	Label    string    `xml:"alert:label,attr,omitempty"`
	Behavior *Behavior `xml:",omitempty"`
}

type SwipeRow struct {
	XMLName    xml.Name `xml:"swipe:row"`
	XmlnsSwipe string   `xml:"xmlns:swipe,attr,omitempty"`
	Style      string   `xml:"style,attr,omitempty"`
}

type SwipeRowParams struct {
	SwipeRow     `xml:",any"`
	SwipeMain    SwipeMainParams //`xml:"swipe:main"` //omitempty pointer???????
	SwipeButtons []SwipeButton   `xml:"swipe:button,omitempty"`
}

type SwipeMain struct {
	XMLName xml.Name `xml:"swipe:main"`
	//Content  string    `xml:",chardata"`
}

type SwipeMainParams struct {
	SwipeMain `xml:",any"`
	View      *View `xml:"view,omitempty"`
}

type SwipeButton struct {
	//XMLName    xml.Name `xml:"swipe:button"`
	View *View `xml:"view,omitempty"`
}

type Spinner struct {
	XMLName xml.Name `xml:"spinner"`
}
