package v1

import (
	"context"
	"net/http"

	"github.com/adamwoolhether/htmxbook/app/contact/v1/contactsgrp"
	"github.com/adamwoolhether/htmxbook/foundation/logger"
	"github.com/adamwoolhether/htmxbook/foundation/web"
)

// RegisterRoutes adds specific routes for this group.
func RegisterRoutes(build string, log *logger.Logger, app *web.App) {
	//const version = "v1"

	app.Handle(http.MethodGet, "", "/", root)

	contactsgrp.Routes(build, log, app)
}

func root(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	web.Redirect(w, r, "/v1/contacts")

	return nil
}
