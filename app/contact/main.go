package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/adamwoolhether/hypermedia/app/contact/v1/handlers"
	"github.com/adamwoolhether/hypermedia/app/frontend"
	v1 "github.com/adamwoolhether/hypermedia/business/web/v1"
	"github.com/adamwoolhether/hypermedia/foundation/logger"
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

	app := v1.APIMux(
		v1.APIMuxConfig{
			Build:    build,
			Shutdown: shutdown,
			Log:      log,
		}, handlers.Routes(), v1.WithStaticFS(frontend.Static()))

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
