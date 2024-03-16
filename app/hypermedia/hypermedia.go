package hypermedia

import (
	"context"
	"net/http"

	"github.com/adamwoolhether/hypermedia/app/hypermedia/handlers/contactsgrp"
	"github.com/adamwoolhether/hypermedia/business/contacts"
	"github.com/adamwoolhether/hypermedia/foundation/logger"
	"github.com/adamwoolhether/hypermedia/foundation/session"
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

const prefix = ""

type Config struct {
	Log      *logger.Logger
	Session  *session.Store
	Contacts *contacts.Core
}

func Routes(app *web.App, cfg Config) {
	app.Handle(http.MethodGet, prefix, "/", root)

	contactsGrp := contactsgrp.New(cfg.Log, cfg.Contacts, cfg.Session)
	app.Handle(http.MethodGet, prefix, "/contacts", contactsGrp.Query)
	app.Handle(http.MethodDelete, prefix, "/contacts", contactsGrp.DeleteBatch)
	app.Handle(http.MethodGet, prefix, "/contacts/count", contactsGrp.Count)
	app.Handle(http.MethodGet, prefix, "/contacts/new", contactsGrp.CreateForm)
	app.Handle(http.MethodPost, prefix, "/contacts/new", contactsGrp.Create)
	app.Handle(http.MethodGet, prefix, "/contacts/{id}", contactsGrp.QueryByID)
	app.Handle(http.MethodGet, prefix, "/contacts/{id}/email", contactsGrp.ValidateEmail)
	app.Handle(http.MethodGet, prefix, "/contacts/{id}/edit", contactsGrp.UpdateForm)
	app.Handle(http.MethodPost, prefix, "/contacts/{id}/edit", contactsGrp.Update)
	app.Handle(http.MethodDelete, prefix, "/contacts/{id}", contactsGrp.Delete)
	app.Handle(http.MethodPost, prefix, "/contacts/archive", contactsGrp.Archive)
	app.Handle(http.MethodDelete, prefix, "/contacts/archive", contactsGrp.ArchiveRm)
	app.Handle(http.MethodGet, prefix, "/contacts/archive", contactsGrp.ArchivePoll)
	app.Handle(http.MethodGet, prefix, "/contacts/archive/file", contactsGrp.ArchiveDL)
	// To test very slow responses.
	app.Handle(http.MethodGet, prefix, "/contacts/slow", contactsGrp.Slow)
}

func root(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	web.Redirect(w, r, "/contacts")

	return nil
}
