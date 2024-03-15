package contactsgrp

import (
	"net/http"

	"github.com/adamwoolhether/hypermedia/business/contacts"
	"github.com/adamwoolhether/hypermedia/foundation/logger"
	"github.com/adamwoolhether/hypermedia/foundation/session"
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(prefix string, log *logger.Logger, store *session.Store, app *web.App) {

	core := contacts.NewCore(log)
	handlers := New(log, core, store)
	app.Handle(http.MethodGet, prefix, "/contacts", handlers.Query)
	app.Handle(http.MethodDelete, prefix, "/contacts", handlers.DeleteBatch)
	app.Handle(http.MethodGet, prefix, "/contacts/count", handlers.Count)
	app.Handle(http.MethodGet, prefix, "/contacts/new", handlers.CreateForm)
	app.Handle(http.MethodPost, prefix, "/contacts/new", handlers.Create)
	app.Handle(http.MethodGet, prefix, "/contacts/{id}", handlers.QueryByID)
	app.Handle(http.MethodGet, prefix, "/contacts/{id}/email", handlers.ValidateEmail)
	app.Handle(http.MethodGet, prefix, "/contacts/{id}/edit", handlers.UpdateForm)
	app.Handle(http.MethodPost, prefix, "/contacts/{id}/edit", handlers.Update)
	app.Handle(http.MethodDelete, prefix, "/contacts/{id}", handlers.Delete)
	app.Handle(http.MethodPost, prefix, "/contacts/archive", handlers.Archive)
	app.Handle(http.MethodDelete, prefix, "/contacts/archive", handlers.ArchiveRm)
	app.Handle(http.MethodGet, prefix, "/contacts/archive", handlers.ArchivePoll)
	app.Handle(http.MethodGet, prefix, "/contacts/archive/file", handlers.ArchiveDL)

	// To test very slow responses.
	app.Handle(http.MethodGet, prefix, "/contacts/slow", handlers.Slow)
}
