package main

/*
- Why bulk delete is not sent as form data?
- Better flash handling: how to get flashes from session directly in template?
- Move the struct conversion funcs.
- Smooth progress bar transition, needed to use JS.
- Fix text above progress bar moving it downwards.
// How to prevent default in hyperscript?? Alt-S
*/
import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/adamwoolhether/hypermedia/app/contact/hypermedia/frontend"
	"github.com/adamwoolhether/hypermedia/app/contact/hypermedia/handlers"
	web2 "github.com/adamwoolhether/hypermedia/business/web"
	"github.com/adamwoolhether/hypermedia/foundation/logger"
	"github.com/adamwoolhether/hypermedia/foundation/session"
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

func main() {
	traceIDFunc := func(ctx context.Context) string {
		return web.GetTraceID(ctx)
	}

	log := logger.New(os.Stdout, logger.LevelInfo, "HTMX", traceIDFunc)

	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "msg", err)
		os.Exit(1)
	}
}

const build = "test"

func run(ctx context.Context, log *logger.Logger) error {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, os.Kill)

	cookieStore := session.New("super-secret-key-for-now")
	app := web2.APIMux(
		web2.APIMuxConfig{
			Build:    build,
			Shutdown: shutdown,
			Log:      log,
			Session:  cookieStore,
		}, handlers.Routes(), web2.WithStaticFS(frontend.Static()))

	api := http.Server{
		Addr:    "localhost:42069",
		Handler: app,
		//ReadTimeout:  cfg.Web.ReadTimeout,
		//WriteTimeout: cfg.Web.WriteTimeout,
		//IdleTimeout:  cfg.Web.IdleTimeout,
		ErrorLog: logger.NewStdLogger(log, logger.LevelError),
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Info(ctx, "startup", "status", "api router started", "host", api.Addr)

		serverErrors <- api.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Info(ctx, "shutdown", "status", "shutdown started", "signal", sig)
		defer log.Info(ctx, "shutdown", "status", "shutdown complete", "signal", sig)

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
