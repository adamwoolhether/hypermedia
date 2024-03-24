package main

/*
- Better flash handling: how to get flashes from session directly in template?
- Smooth progress bar transition, needed to use JS.
- debounce for mobile hyperview not working
- how to check email validation as user typing in mobile app?
*/
import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/adamwoolhether/hypermedia/app/dataapi"
	"github.com/adamwoolhether/hypermedia/app/hypermedia"
	fe "github.com/adamwoolhether/hypermedia/app/hypermedia/web"
	"github.com/adamwoolhether/hypermedia/business/contacts"
	"github.com/adamwoolhether/hypermedia/business/web/mux"
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
	app := mux.WebApp(
		mux.WebAppConfig{
			Build:    build,
			Shutdown: shutdown,
			Log:      log,
			Session:  cookieStore,
		}, Routes(), mux.WithStaticFS(fe.Static()))

	api := http.Server{
		Addr:    "192.168.1.71:42069",
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

func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouteAdder interface.
func (add) Add(app *web.App, cfg mux.WebAppConfig) {
	contactsCore := contacts.NewCore(cfg.Log)

	hypermedia.Routes(app, hypermedia.Config{
		Log:      cfg.Log,
		Session:  cfg.Session,
		Contacts: contactsCore,
	})

	dataapi.Routes(app, dataapi.Config{
		Log:      cfg.Log,
		Contacts: contactsCore,
	})
}
