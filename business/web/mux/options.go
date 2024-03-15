package mux

import (
	"context"
	"net/http"

	"github.com/adamwoolhether/hypermedia/foundation/web"
)

// Options represents optional parameters.
type Options struct {
	corsOrigin string
	staticFS   web.Handler
}

// WithCORS provides configuration options for CORS.
func WithCORS(origin string) func(opts *Options) {
	return func(opts *Options) {
		opts.corsOrigin = origin
	}
}

func WithStaticFS(fs http.Handler) func(opts *Options) {
	f := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		fs.ServeHTTP(w, r)

		return nil
	}

	return func(opts *Options) {
		opts.staticFS = f
	}
}
