package contacts

import (
	"slices"

	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/layout"
	"github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/xml"
)

func Deleted(toastMsg string) xml.View {
	toasts := layout.ShowToasts(toastMsg)
	deleteEvents := []xml.Behavior{
		{
			Trigger:   "load",
			Action:    "dispatch-event",
			EventName: "contact-updated",
		},
		{
			Trigger: "load",
			Action:  "back",
		},
	}

	deleted := xml.View{
		Xmlns:    xml.Namespace,
		Behavior: slices.Concat(toasts, deleteEvents),
	}

	return deleted
}
