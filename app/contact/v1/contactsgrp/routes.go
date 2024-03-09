package contactsgrp

import (
	"net/http"

	"github.com/adamwoolhether/htmxbook/business/contacts"
	"github.com/adamwoolhether/htmxbook/foundation/logger"
	"github.com/adamwoolhether/htmxbook/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(build string, log *logger.Logger, app *web.App) {
	const version = "v1"

	core := contacts.NewCore(log)
	handlers := New(build, log, core)
	app.Handle(http.MethodGet, version, "/contacts", handlers.Query)
	app.Handle(http.MethodGet, version, "/contacts/{id}/view", handlers.QueryByID)
	app.Handle(http.MethodGet, version, "/contacts/{id}/edit", handlers.Update)
}
