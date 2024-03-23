package contacts

import (
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/layout"
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xmlmodel"
)

func Deleted() xmlmodel.View {
	deleted := xmlmodel.View{
		Xmlns: layout.Namespace,
		Behavior: []xmlmodel.Behavior{
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
