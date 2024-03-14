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
	app.Handle(http.MethodDelete, version, "/contacts", handlers.DeleteBatch)
	app.Handle(http.MethodGet, version, "/contacts/count", handlers.Count)
	app.Handle(http.MethodGet, version, "/contacts/new", handlers.CreateForm)
	app.Handle(http.MethodPost, version, "/contacts/new", handlers.Create)
	app.Handle(http.MethodGet, version, "/contacts/{id}", handlers.QueryByID)
	app.Handle(http.MethodGet, version, "/contacts/{id}/email", handlers.ValidateEmail)
	app.Handle(http.MethodGet, version, "/contacts/{id}/edit", handlers.UpdateForm)
	app.Handle(http.MethodPost, version, "/contacts/{id}/edit", handlers.Update)
	app.Handle(http.MethodDelete, version, "/contacts/{id}", handlers.Delete)
	app.Handle(http.MethodPost, version, "/contacts/archive", handlers.Archive)
	app.Handle(http.MethodDelete, version, "/contacts/archive", handlers.ArchiveRm)
	app.Handle(http.MethodGet, version, "/contacts/archive", handlers.ArchivePoll)
	app.Handle(http.MethodGet, version, "/contacts/archive/file", handlers.ArchiveDL)

	// To test very slow responses.
	app.Handle(http.MethodGet, version, "/contacts/slow", handlers.Slow)
}
