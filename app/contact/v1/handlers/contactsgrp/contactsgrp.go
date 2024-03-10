package contactsgrp

import (
	"context"
	"net/http"

	"github.com/adamwoolhether/hypermedia/app/frontend/view/contacts"

	"github.com/adamwoolhether/hypermedia/business/contacts"
	"github.com/adamwoolhether/hypermedia/foundation/logger"
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

// Handlers manages the set of check points.
type Handlers struct {
	log  *logger.Logger
	core *contacts.Core
	//db json.RawMessage
}

func New(build string, log *logger.Logger, core *contacts.Core) *Handlers {
	return &Handlers{
		log:  log,
		core: core,
	}
}

func (h *Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	query := web.QueryString(r, "q")

	contacts, err := h.core.Query(ctx, query)
	if err != nil {
		return err
	}

	return index.Index(query, contacts).Render(ctx, w)
}

func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := web.Param(r, "id")

	h.log.Error(ctx, "UNIMPLEMENTED", "id", userID)
	return nil
}

func (h *Handlers) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := web.Param(r, "id")

	h.log.Error(ctx, "UNIMPLEMENTED", "id", userID)
	return nil
}
