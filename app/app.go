package app

import (
	"github.com/adamwoolhether/hypermedia/app/dataapi"
	"github.com/adamwoolhether/hypermedia/app/hypermedia"
	"github.com/adamwoolhether/hypermedia/business/contacts"
	"github.com/adamwoolhether/hypermedia/business/web/mux"
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

// Routes constructs the add value which provides the implementation
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouteAdder interface.
func (add) Add(app *web.App, cfg mux.WebAppConfig) {
	contactsCore := contacts.NewCore(cfg.Log)

	hypermedia.Routes(app, hypermedia.Config{
		Log:      cfg.Log,
		Session:  cfg.Session,
		Contacts: contactsCore,
	})

	dataapi.Routes(app, dataapi.Config{
		Log:      cfg.Log,
		Contacts: contactsCore,
	})
}
