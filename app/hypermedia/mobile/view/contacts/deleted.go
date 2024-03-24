package contacts

import (
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xml"
)

func Deleted() xml.View {
	deleted := xml.View{
		Xmlns: xml.Namespace,
		Behavior: []xml.Behavior{
			{
				Trigger:   "load",
				Action:    "dispatch-event",
				EventName: "contact-updated",
			},
			{
				Trigger: "load",
				Action:  "back",
			},
		},
	}

	return deleted
}
