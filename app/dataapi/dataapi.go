package dataapi

import (
	"net/http"

	"github.com/adamwoolhether/hypermedia/app/dataapi/v1/contactsgrp"
	"github.com/adamwoolhether/hypermedia/business/contacts"
	"github.com/adamwoolhether/hypermedia/foundation/logger"
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

type Config struct {
	Log      *logger.Logger
	Contacts *contacts.Core
}

func Routes(app *web.App, cfg Config) {
	// V1
	const v1 = "api/v1"
	contactsGrp := contactsgrp.New(cfg.Log, cfg.Contacts)
	app.Handle(http.MethodGet, v1, "/contacts", contactsGrp.Query)
	app.Handle(http.MethodPost, v1, "/contacts", contactsGrp.Create)
	app.Handle(http.MethodGet, v1, "/contacts/{id}", contactsGrp.QueryByID)
	app.Handle(http.MethodPut, v1, "/contacts/{id}", contactsGrp.Update)
	app.Handle(http.MethodDelete, v1, "/contacts/{id}", contactsGrp.Delete)
}
