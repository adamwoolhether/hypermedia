package contactsgrp

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	fe "github.com/adamwoolhether/hypermedia/app/hypermedia/mobile/view/contacts"
	"github.com/adamwoolhether/hypermedia/business/contacts"
	"github.com/adamwoolhether/hypermedia/foundation/logger"
	"github.com/adamwoolhether/hypermedia/foundation/validate"
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

const defaultRows = 20

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
		return web.RenderXML(ctx, w, fe.Rows(contactsToMobile(contacts), page), http.StatusOK)
	}

	return web.RenderXML(ctx, w, fe.Index(contactsToMobile(contacts), page), http.StatusOK)
}

func (h *Handlers) CreateForm(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	return web.RenderXML(ctx, w, fe.New(fe.UpdateContact{}), http.StatusOK)
}

func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	newContact := fe.UpdateContact{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Phone:     r.FormValue("phone"),
		Email:     r.FormValue("email"),
	}

	if err := validate.Check(newContact); err != nil {
		fieldErrs := validate.GetFieldErrors(err)

		newContact.FieldErrs = fe.ContactErrors{
			FirstName: fieldErrs.Fields()["first_name"],
			LastName:  fieldErrs.Fields()["last_name"],
			Phone:     fieldErrs.Fields()["phone"],
			Email:     fieldErrs.Fields()["email"],
		}

		return web.RenderXML(ctx, w, fe.FormFields(newContact, false), http.StatusBadRequest)
	}

	created, err := h.core.Create(ctx, newContact.ToDB())
	if err != nil {
		newContact.InternalErrors = err.Error()

		return web.RenderXML(ctx, w, fe.FormFields(newContact, false), http.StatusInternalServerError)
	}

	newContact.ID = created.ID

	return web.RenderXML(ctx, w, fe.FormFields(newContact, true, "Contact added"), http.StatusCreated)
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

	return web.RenderXML(ctx, w, fe.Show(contactToMobile(contact)), http.StatusOK)
}

func (h *Handlers) UpdateForm(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := web.Param(r, "id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		// better err handling
		return err
	}

	contact, err := h.core.QueryByID(ctx, id)
	if err != nil {
		return err
	}

	uc := fe.UpdateContact{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
		Email:     contact.Email,
	}

	return web.RenderXML(ctx, w, fe.Edit(uc), http.StatusOK)
}

func (h *Handlers) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := web.Param(r, "id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	contact, err := h.core.QueryByID(ctx, id)
	if err != nil {
		return err
	}

	uc := fe.UpdateContact{
		ID:        contact.ID,
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Phone:     r.FormValue("phone"),
		Email:     r.FormValue("email"),
		//FieldErrs:      fe.ContactErrors{},
		//InternalErrors: "",
	}

	if err := validate.Check(uc); err != nil {
		fieldErrs := validate.GetFieldErrors(err)

		uc.FieldErrs = fe.ContactErrors{
			FirstName: fieldErrs.Fields()["first_name"],
			LastName:  fieldErrs.Fields()["last_name"],
			Phone:     fieldErrs.Fields()["phone"],
			Email:     fieldErrs.Fields()["email"],
		}

		return web.RenderXML(ctx, w, fe.FormFields(uc, false), http.StatusOK)
	}

	err = h.core.Update(ctx, uc.ToDB())
	if err != nil {
		// Or do failure flash here?
		uc.InternalErrors = err.Error()
		return web.RenderXML(ctx, w, fe.FormFields(uc, false), http.StatusInternalServerError)
	}

	//if err := h.sessions.AddFlash(w, r, "Updated contact!"); err != nil {
	//	h.log.Error(ctx, "adding flash", "err", err)
	//}

	return web.RenderXML(ctx, w, fe.FormFields(uc, true, "Contact updated"), http.StatusOK)
}

func (h *Handlers) ValidateEmail(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := web.Param(r, "id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	email := r.FormValue("email")

	uc := fe.UpdateContact{
		ID:    id,
		Email: email,
	}

	if !h.core.UniqueEmail(ctx, id, strings.ToLower(email)) {
		uc.FieldErrs = fe.ContactErrors{Email: "This email is taken"}
		return web.RenderXML(ctx, w, fe.CheckEmailErr(uc), http.StatusBadRequest)
	}

	// We need to return the `text` element with empty error.
	return web.RenderXML(ctx, w, fe.CheckEmailErr(uc), http.StatusOK)
}

func (h *Handlers) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := web.Param(r, "id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	if err := h.core.Delete(ctx, id); err != nil {
		return err
	}

	return web.RenderXML(ctx, w, fe.Deleted("Contact deleted"), http.StatusNoContent)
}
