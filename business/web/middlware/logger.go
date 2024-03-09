package mid

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/adamwoolhether/htmxbook/foundation/logger"
	"github.com/adamwoolhether/htmxbook/foundation/web"
)

// Logger writes some information about the request to the logs.
// Format: TraceID : (200) GET /foo -> IP ADDR (latency)
func Logger(log *logger.Logger) web.Middleware {

	// This is the actual middleware function to be executed.
	m := func(handler web.Handler) web.Handler {

		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			v := web.GetValues(ctx)

			path := r.URL.Path
			if r.URL.RawQuery != "" {
				path = fmt.Sprintf("%s?%s", path, r.URL.RawQuery)
			}

			log.Info(ctx, "request started", "method", r.Method, "path", path, "remoteaddr", r.RemoteAddr)

			// Call the next handler.
			err := handler(ctx, w, r)

			log.Info(ctx, "request completed", "method", r.Method, "path", path, "remoteaddr", r.RemoteAddr, "statusCode", v.StatusCode, "since", time.Since(v.Now))

			// Return the error, so it can be handled further up the chain.
			return err
		}

		return h
	}

	return m
}
