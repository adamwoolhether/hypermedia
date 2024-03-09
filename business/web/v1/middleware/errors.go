package mid

import (
	"context"
	"net/http"

	"github.com/adamwoolhether/hypermedia/business/web/v1/response"
	"github.com/adamwoolhether/hypermedia/foundation/logger"
	"github.com/adamwoolhether/hypermedia/foundation/validate"
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

// Errors handles errors coming out of the call chain. It detects normal
// application errors which are used to respond to the client in a
// uniform way. Unexpected errors (status >= 500) are logged.
func Errors(log *logger.Logger) web.Middleware {

	// This is the actual middleware function to be executed.
	m := func(handler web.Handler) web.Handler {

		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			// Run the next handler and catch any propagated error.
			if err := handler(ctx, w, r); err != nil {

				// Log the error.
				log.Error(ctx, "message", "msg", err)

				var er response.ErrorDocument
				var status int
				switch {
				case response.IsError(err):
					reqErr := response.GetError(err)
					if validate.IsFieldErrors(reqErr.Err) {
						fieldErrors := validate.GetFieldErrors(reqErr.Err)
						er = response.ErrorDocument{
							Error:  "data validation error",
							Fields: fieldErrors.Fields(),
						}
						status = reqErr.Status
						break
					}

					er = response.ErrorDocument{
						Error: reqErr.Error(),
					}
					status = reqErr.Status
				//case auth.IsAuthError(err):
				//	er = response.ErrorDocument{
				//		Error: http.StatusText(http.StatusUnauthorized),
				//	}
				//	status = http.StatusUnauthorized
				default:
					er = response.ErrorDocument{
						Error: http.StatusText(http.StatusInternalServerError),
					}
					status = http.StatusInternalServerError
				}

				// Response with the error back to the client.
				if err = web.Respond(ctx, w, er, status); err != nil {
					return err
				}

				// If we receive the shutdown err we need to return it
				// back to teh base handler to shut down the service.
				if web.IsShutdown(err) {
					return err
				}
			}

			// The error has been handled, so we stop propagating it.
			return nil
		}

		return h
	}

	return m
}
