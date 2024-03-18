package contactsgrp

import (
	"context"
	"net/http"
	"strconv"

	fe "github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/contacts"
	"github.com/adamwoolhether/hypermedia/business/contacts"
	"github.com/adamwoolhether/hypermedia/foundation/logger"
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

const defaultRows = 10

// Handlers manages the set of check points.
type Handlers struct {
	log  *logger.Logger
	core *contacts.Core
	//sessions *session.Store
}

func New(log *logger.Logger, core *contacts.Core) *Handlers {
	return &Handlers{
		log:  log,
		core: core,
		//sessions: store,
	}
}

func (h *Handlers) RootRedirect(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	web.Redirect(w, r, "/mobile/contacts")

	return nil
}

func (h *Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	query := web.QueryString(r, "q")
	page := 1
	pageStr := web.QueryString(r, "page")
	if len(pageStr) > 0 {
		p, err := strconv.Atoi(pageStr)
		if err == nil {
			page = p
		}
	}

	rowsOnly := false
	rowsOnlyQS := web.QueryString(r, "rows_only")
	if len(rowsOnlyQS) > 0 {
		ro, err := strconv.ParseBool(rowsOnlyQS)
		if err == nil {
			rowsOnly = ro
		}
	}

	contacts, err := h.core.Query(ctx, query, page, defaultRows)
	if err != nil {
		return err
	}

	if rowsOnly {
		return web.RenderXML(ctx, w, fe.Rows(contacts, page), http.StatusOK)
	}

	return web.RenderXML(ctx, w, fe.Index(contacts, page), http.StatusOK)
}

func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := web.Param(r, "id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	contact, err := h.core.QueryByID(ctx, id)
	if err != nil {
		return err
	}

	return web.RenderXML(ctx, w, fe.Show(contact), http.StatusOK)
}
