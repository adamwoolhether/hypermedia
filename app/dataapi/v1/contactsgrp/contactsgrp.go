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

	return web.RespondJSON(ctx, w, resp, http.StatusOK)
}

func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var newContact newContact
	if err := web.Decode(r, &newContact); err != nil {
		return response.NewError(err, http.StatusBadRequest)
	}

	created, err := h.core.Create(ctx, newContact.toDB())
	if err != nil {
		return response.NewError(err, http.StatusInternalServerError)
	}

	return web.RespondJSON(ctx, w, created, http.StatusCreated)
}

func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := web.Param(r, "id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return response.NewError(err, http.StatusBadRequest)
	}

	contact, err := h.core.QueryByID(ctx, id)
	if err != nil {
		return response.NewError(err, http.StatusInternalServerError)
	}

	return web.RespondJSON(ctx, w, contactToAPI(contact), http.StatusOK)
}

func (h *Handlers) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var uc updateContact
	if err := web.Decode(r, &uc); err != nil {
		return response.NewError(err, http.StatusBadRequest)
	}

	userID := web.Param(r, "id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	contact, err := h.core.QueryByID(ctx, id)
	if err != nil {
		return err
	}

	if uc.FirstName != nil {
		contact.FirstName = *uc.FirstName
	}
	if uc.LastName != nil {
		contact.LastName = *uc.LastName
	}
	if uc.Email != nil {
		contact.Email = *uc.Email
	}
	if uc.Phone != nil {
		contact.Phone = *uc.Phone
	}

	if err := h.core.Update(ctx, contact); err != nil {
		return response.NewError(err, http.StatusInternalServerError)
	}

	return web.RespondJSON(ctx, w, contactToAPI(contact), http.StatusOK)
}

func (h *Handlers) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := web.Param(r, "id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	if err := h.core.Delete(ctx, id); err != nil {
		return response.NewError(err, http.StatusInternalServerError)
	}

	return web.RespondJSON(ctx, w, nil, http.StatusNoContent)
}
