package contactsgrp

import (
	"context"
	"net/http"

	"github.com/adamwoolhether/htmxbook/app/frontend/view/index"
	"github.com/adamwoolhether/htmxbook/business/contacts"
	"github.com/adamwoolhether/htmxbook/foundation/logger"
	"github.com/adamwoolhether/htmxbook/foundation/web"
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
