package handlers

import (
	"context"
	"net/http"

	"github.com/adamwoolhether/hypermedia/app/contact/v1/handlers/contactsgrp"
	v1 "github.com/adamwoolhether/hypermedia/business/web/v1"
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

// Routes constructs the add value which provides the implementation
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouteAdder interface.
func (add) Add(app *web.App, cfg v1.APIMuxConfig) {
	app.Handle(http.MethodGet, "", "/", root)

	contactsgrp.Routes(cfg.Build, cfg.Log, cfg.Session, app)
}

func root(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	web.Redirect(w, r, "/v1/contacts")

	return nil
}
