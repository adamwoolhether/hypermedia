package hypermedia

import (
	"net/http"

	"github.com/adamwoolhether/hypermedia/app/hypermedia/handlers/contactsgrp"
	mobileHandlers "github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/handlers/contactsgrp"
	"github.com/adamwoolhether/hypermedia/business/contacts"
	"github.com/adamwoolhether/hypermedia/foundation/logger"
	"github.com/adamwoolhether/hypermedia/foundation/session"
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

type Config struct {
	Log      *logger.Logger
	Session  *session.Store
	Contacts *contacts.Core
}

func Routes(app *web.App, cfg Config) {
	webRoutes(app, cfg)
	mobileRoutes(app, cfg)
}

func webRoutes(app *web.App, cfg Config) {
	const root = ""

	contactsGrp := contactsgrp.New(cfg.Log, cfg.Contacts, cfg.Session)
	app.Handle(http.MethodGet, root, "/", contactsGrp.RootRedirect)
	app.Handle(http.MethodGet, root, "/contacts", contactsGrp.Query)
	app.Handle(http.MethodDelete, root, "/contacts", contactsGrp.DeleteBatch)
	app.Handle(http.MethodGet, root, "/contacts/count", contactsGrp.Count)
	app.Handle(http.MethodGet, root, "/contacts/new", contactsGrp.CreateForm)
	app.Handle(http.MethodPost, root, "/contacts/new", contactsGrp.Create)
	app.Handle(http.MethodGet, root, "/contacts/{id}", contactsGrp.QueryByID)
	app.Handle(http.MethodGet, root, "/contacts/{id}/email", contactsGrp.ValidateEmail)
	app.Handle(http.MethodGet, root, "/contacts/{id}/edit", contactsGrp.UpdateForm)
	app.Handle(http.MethodPost, root, "/contacts/{id}/edit", contactsGrp.Update)
	app.Handle(http.MethodDelete, root, "/contacts/{id}", contactsGrp.Delete)
	app.Handle(http.MethodPost, root, "/contacts/archive", contactsGrp.Archive)
	app.Handle(http.MethodDelete, root, "/contacts/archive", contactsGrp.ArchiveRm)
	app.Handle(http.MethodGet, root, "/contacts/archive", contactsGrp.ArchivePoll)
	app.Handle(http.MethodGet, root, "/contacts/archive/file", contactsGrp.ArchiveDL)
	// To test very slow responses.
	app.Handle(http.MethodGet, root, "/contacts/slow", contactsGrp.Slow)
}

func mobileRoutes(app *web.App, cfg Config) {
	const mobile = "mobile"

	mobileContactsGrp := mobileHandlers.New(cfg.Log, cfg.Contacts)
	app.Handle(http.MethodGet, mobile, "/", mobileContactsGrp.RootRedirect)
	app.Handle(http.MethodGet, mobile, "/contacts", mobileContactsGrp.Query)
	app.Handle(http.MethodGet, mobile, "/contacts/{id}", mobileContactsGrp.QueryByID)
	app.Handle(http.MethodGet, mobile, "/contacts/{id}/edit", mobileContactsGrp.UpdateForm)
	app.Handle(http.MethodPost, mobile, "/contacts/{id}/edit", mobileContactsGrp.Update)
}
