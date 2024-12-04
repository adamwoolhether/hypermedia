package hypermedia

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	mobilehandlers "github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/handlers/contactsgrp"
	webhandlers "github.com/adamwoolhether/hypermedia/app/hypermedia/web/handlers/contactsgrp"
	"github.com/adamwoolhether/hypermedia/business/contacts"
	"github.com/adamwoolhether/hypermedia/business/web/response"
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
	const root = ""

	rootRedirect := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		accept := r.Header.Get("Accept")
		fmt.Println("accept", accept)
		switch {
		case strings.Contains(accept, web.HXMLMime):
			web.Redirect(w, r, "/mobile/contacts")
		case strings.Contains(accept, web.HTMLMime):
			web.Redirect(w, r, "/contacts")
		default:
			return response.NewError(errors.New("invalid accept headers"), http.StatusBadRequest)
		}

		return nil
	}

	app.Handle(http.MethodGet, root, "/", rootRedirect)

	webRoutes(app, cfg)
	mobileRoutes(app, cfg)
}

func webRoutes(app *web.App, cfg Config) {
	const root = ""

	contactsGrp := webhandlers.New(cfg.Log, cfg.Contacts, cfg.Session)
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

	mobileContactsGrp := mobilehandlers.New(cfg.Log, cfg.Contacts)
	app.Handle(http.MethodGet, mobile, "/contacts", mobileContactsGrp.Query)
	app.Handle(http.MethodGet, mobile, "/contacts/new", mobileContactsGrp.CreateForm)
	app.Handle(http.MethodPost, mobile, "/contacts/new", mobileContactsGrp.Create)
	app.Handle(http.MethodGet, mobile, "/contacts/{id}", mobileContactsGrp.QueryByID)
	app.Handle(http.MethodGet, mobile, "/contacts/{id}/edit", mobileContactsGrp.UpdateForm)
	app.Handle(http.MethodPost, mobile, "/contacts/{id}/edit", mobileContactsGrp.Update)
	app.Handle(http.MethodPost, mobile, "/contacts/{id}/delete", mobileContactsGrp.Delete) // hyperview currently doesn't support DELETE
	app.Handle(http.MethodGet, mobile, "/contacts/{id}/email", mobileContactsGrp.ValidateEmail)
}
