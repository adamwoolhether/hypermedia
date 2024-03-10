package contactsgrp

import (
	"net/http"

	"github.com/adamwoolhether/hypermedia/business/contacts"
	"github.com/adamwoolhether/hypermedia/foundation/logger"
	"github.com/adamwoolhether/hypermedia/foundation/session"
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(build string, log *logger.Logger, store *session.Store, app *web.App) {
	const version = "v1"

	core := contacts.NewCore(log)
	handlers := New(build, log, core, store)
	app.Handle(http.MethodGet, version, "/contacts", handlers.Query)
	app.Handle(http.MethodGet, version, "/contacts/new", handlers.CreateForm)
	app.Handle(http.MethodPost, version, "/contacts/new", handlers.Create)
	app.Handle(http.MethodGet, version, "/contacts/{id}/view", handlers.QueryByID)
	app.Handle(http.MethodGet, version, "/contacts/{id}/edit", handlers.Update)
}
