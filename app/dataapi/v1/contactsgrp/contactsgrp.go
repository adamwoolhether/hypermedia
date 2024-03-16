package contactsgrp

import (
	"context"
	"math"
	"net/http"
	"strconv"

	"github.com/adamwoolhether/hypermedia/business/contacts"
	"github.com/adamwoolhether/hypermedia/business/web/response"
	"github.com/adamwoolhether/hypermedia/foundation/logger"
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

type Handlers struct {
	log  *logger.Logger
	core *contacts.Core
}

func New(log *logger.Logger, core *contacts.Core) *Handlers {
	return &Handlers{
		log:  log,
		core: core,
	}
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

	rows := math.MaxInt
	rowsStr := web.QueryString(r, "rows")
	if len(rowsStr) > 0 {
		r, err := strconv.Atoi(rowsStr)
		if err == nil {
			rows = r
		}
	}

	contacts, err := h.core.Query(ctx, query, page, rows)
	if err != nil {
		return response.NewError(err, http.StatusInternalServerError)
	}

	resp := newResponse(contacts, h.core.Count(), page, rows)

	return web.Respond(ctx, w, resp, http.StatusOK)
}
