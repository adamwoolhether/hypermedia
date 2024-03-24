package layout

import "github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xml"

func ShowToasts(messages ...string) []xml.Behavior {
	if len(messages) < 1 {
		return nil
	}

	behaviors := make([]xml.Behavior, len(messages))

	for i, msg := range messages {
		behaviors[i] = xml.Behavior{
			XmlnsToast: xml.NamespaceToast,
			Trigger:    "load",
			Action:     "show-toast",
			ToastText:  msg,
		}
	}

	return behaviors
}
