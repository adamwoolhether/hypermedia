package mux

import (
	"net/http"
	"os"

	middleware "github.com/adamwoolhether/hypermedia/business/web/middleware"
	"github.com/adamwoolhether/hypermedia/foundation/logger"
	"github.com/adamwoolhether/hypermedia/foundation/session"
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

type RouteAdder interface {
	Add(app *web.App, cfg WebAppConfig)
}

// WebAppConfig contains all the mandatory systems required by handlers.
type WebAppConfig struct {
	Build    string
	Shutdown chan os.Signal
	Log      *logger.Logger
	Session  *session.Store
}

// WebApp constructs a http.Handler with all application routes defined.
func WebApp(cfg WebAppConfig, routes RouteAdder, options ...func(opts *Options)) http.Handler {
	var opts Options
	for _, option := range options {
		option(&opts)
	}
	app := web.NewApp(
		cfg.Shutdown,
		middleware.Logger(cfg.Log),
		middleware.Errors(cfg.Log),
		middleware.Panics(),
	)

	if opts.corsOrigin != "" {
		app.EnableCORS(middleware.Cors(opts.corsOrigin))
	}

	if opts.staticFS != nil {
		app.HandleNoMiddleware(http.MethodGet, "", "/static/*", opts.staticFS)
	}

	routes.Add(app, cfg)

	return app
}
