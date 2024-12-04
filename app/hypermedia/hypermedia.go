package hypermedia

import (
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

type Config struct {
}

func Routes(app *web.App, cfg Config) {
	const root = ""

	// webRoutes(app, cfg)
}

// func webRoutes(app *web.App, cfg Config) {
// 	const root = ""
//
// 	contactsGrp := webhandlers.New(cfg.Log, cfg.Contacts, cfg.Session)
// 	app.Handle(http.MethodGet, root, "/contacts", contactsGrp.Query)
// 	app.Handle(http.MethodDelete, root, "/contacts", contactsGrp.DeleteBatch)
// 	app.Handle(http.MethodGet, root, "/contacts/count", contactsGrp.Count)
// 	app.Handle(http.MethodGet, root, "/contacts/new", contactsGrp.CreateForm)
// 	app.Handle(http.MethodPost, root, "/contacts/new", contactsGrp.Create)
// 	app.Handle(http.MethodGet, root, "/contacts/{id}", contactsGrp.QueryByID)
// 	app.Handle(http.MethodGet, root, "/contacts/{id}/email", contactsGrp.ValidateEmail)
// 	app.Handle(http.MethodGet, root, "/contacts/{id}/edit", contactsGrp.UpdateForm)
// 	app.Handle(http.MethodPost, root, "/contacts/{id}/edit", contactsGrp.Update)
// 	app.Handle(http.MethodDelete, root, "/contacts/{id}", contactsGrp.Delete)
// 	app.Handle(http.MethodPost, root, "/contacts/archive", contactsGrp.Archive)
// 	app.Handle(http.MethodDelete, root, "/contacts/archive", contactsGrp.ArchiveRm)
// 	app.Handle(http.MethodGet, root, "/contacts/archive", contactsGrp.ArchivePoll)
// 	app.Handle(http.MethodGet, root, "/contacts/archive/file", contactsGrp.ArchiveDL)
// 	// To test very slow responses.
// 	app.Handle(http.MethodGet, root, "/contacts/slow", contactsGrp.Slow)
// }
