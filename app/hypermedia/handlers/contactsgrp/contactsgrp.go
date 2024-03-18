package contactsgrp

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	fe "github.com/adamwoolhether/hypermedia/app/hypermedia/frontend/view/contacts"
	"github.com/adamwoolhether/hypermedia/business/contacts"
	"github.com/adamwoolhether/hypermedia/foundation/logger"
	"github.com/adamwoolhether/hypermedia/foundation/session"
	"github.com/adamwoolhether/hypermedia/foundation/validate"
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

const defaultRows = 10

// Handlers manages the set of check points.
type Handlers struct {
	log      *logger.Logger
	core     *contacts.Core
	sessions *session.Store
}

func New(log *logger.Logger, core *contacts.Core, store *session.Store) *Handlers {
	return &Handlers{
		log:      log,
		core:     core,
		sessions: store,
	}
}

func (h *Handlers) RootRedirect(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	web.Redirect(w, r, "/contacts")

	return nil
}

func (h *Handlers) CreateForm(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	return web.RenderHTML(ctx, w, fe.NewForm(fe.NewContact{}), http.StatusOK)
}

func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	newContact := fe.NewContact{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Phone:     r.FormValue("phone"),
		Email:     r.FormValue("email"),
		//FieldErrs: ,
	}

	if err := validate.Check(newContact); err != nil {
		fieldErrs := validate.GetFieldErrors(err)

		newContact.FieldErrs = fe.ContactErrors{
			FirstName: fieldErrs.Fields()["first_name"],
			LastName:  fieldErrs.Fields()["last_name"],
			Phone:     fieldErrs.Fields()["phone"],
			Email:     fieldErrs.Fields()["email"],
		}

		return web.RenderHTML(ctx, w, fe.FormFields(newContact), http.StatusBadRequest)
	}

	err := h.core.Create(ctx, newContact.ToDB())
	if err != nil {
		newContact.InternalErrors = err.Error()

		return web.RenderHTML(ctx, w, fe.NewForm(newContact), http.StatusInternalServerError)
	}

	if err := h.sessions.AddFlash(w, r, "Created contact!"); err != nil {
		h.log.Error(ctx, "adding flash", "err", err)
	}

	web.Redirect(w, r, "/contacts")

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

	// Simulate a slow query so we can see our beautiful spinner at work.
	if len(query) > 0 {
		time.Sleep(500 * time.Millisecond)
	}

	contacts, err := h.core.Query(ctx, query, page, defaultRows)
	if err != nil {
		return err
	}

	// If the specific trigger header is present, then
	// we know we only need to update a very specific
	// part of the page, so we just return the rows.
	if r.Header.Get("HX-Trigger") == "search" {
		return web.RenderHTML(ctx, w, fe.Rows(contactsToWeb(contacts)), http.StatusOK)
	}

	flashCtx := h.sessions.GetFlashCtx(w, r)

	templComponent := fe.Index(query, page, contactsToWeb(contacts), h.core.ArchivePoll(ctx))
	return web.RenderHTML(flashCtx, w, templComponent, http.StatusOK)
}

func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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

	return web.RenderHTML(ctx, w, fe.ShowByID(contactToWeb(contact)), http.StatusOK)
}

func (h *Handlers) ValidateEmail(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := web.Param(r, "id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		// better err handling
		return err
	}

	email := r.FormValue("email")

	if !h.core.UniqueEmail(ctx, id, email) {
		_, err := w.Write([]byte("This email is taken"))
		if err != nil {
			return err
		}
	}

	return nil
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

	// embed the db contact into this to access the fields.
	uc := fe.UpdateContact{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
		Email:     contact.Email,
		//FieldErrs:      fe.ContactErrors{},
		//InternalErrors: "",
	}

	return web.RenderHTML(ctx, w, fe.EditByID(uc), http.StatusOK)
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

		return web.RenderHTML(ctx, w, fe.EditFields(uc), http.StatusBadRequest)
	}

	err = h.core.Update(ctx, uc.ToDB())
	if err != nil {
		// Or do failure flash here?
		uc.InternalErrors = err.Error()
		return web.RenderHTML(ctx, w, fe.EditByID(uc), http.StatusInternalServerError)
	}

	if err := h.sessions.AddFlash(w, r, "Updated contact!"); err != nil {
		h.log.Error(ctx, "adding flash", "err", err)
	}

	web.Redirect(w, r, "/contacts"+userID)

	return nil
}

func (h *Handlers) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID := web.Param(r, "id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	err = h.core.Delete(ctx, id)
	if err != nil {
		return err
	}

	if r.Header.Get("HX-Trigger") != "delete-btn" {
		_, err := w.Write([]byte(""))

		return err
	}

	if err := h.sessions.AddFlash(w, r, "Deleted contact!"); err != nil {
		h.log.Error(ctx, "adding flash", "err", err)
	}

	web.Redirect(w, r, "/contacts")
	return nil
}

func (h *Handlers) DeleteBatch(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// Go doesn't allow us to use `r.ParseForm()` on DELETE requests, so we manually parse the form data from body.
	// We could also just set this as a POST handler, but that seems weird.
	ids, err := parseDeleteForm(r)
	if err != nil {
		return err
	}

	for _, id := range ids {
		if err := h.core.Delete(ctx, id); err != nil {
			h.log.Error(ctx, "deleting contact", "id", id, "err", err)
		}
	}

	if err := h.sessions.AddFlash(w, r, "Deleted contacts!"); err != nil {
		h.log.Error(ctx, "adding flash", "err", err)
	}

	contacts, err := h.core.Query(ctx, "", 1, defaultRows)
	if err != nil {
		return err
	}

	flashCtx := h.sessions.GetFlashCtx(w, r)

	templComponent := fe.Index("", 1, contactsToWeb(contacts), h.core.ArchivePoll(ctx))
	return web.RenderHTML(flashCtx, w, templComponent, http.StatusOK)
}

func (h *Handlers) Count(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// Pretend this is a slow req to demonstrate lazy loading.
	time.Sleep(2000 * time.Millisecond)
	count := h.core.Count()

	retStr := fmt.Sprintf("( %d total Contacts )", count)

	_, err := w.Write([]byte(retStr))

	return err
}

func (h *Handlers) Archive(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if err := h.core.Archive(ctx); err != nil {
		return err
	}

	return web.RenderHTML(ctx, w, fe.Archive(h.core.ArchivePoll(ctx)), http.StatusOK)
}

func (h *Handlers) ArchiveRm(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	h.core.ArchiveRm(ctx)

	return web.RenderHTML(ctx, w, fe.Archive(h.core.ArchivePoll(ctx)), http.StatusNoContent)
}

func (h *Handlers) ArchivePoll(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	return web.RenderHTML(ctx, w, fe.Archive(h.core.ArchivePoll(ctx)), http.StatusOK)
}

func (h *Handlers) ArchiveDL(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	file := h.core.ArchiveFile(ctx)

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", "attachment; filename=\"archive.json\"")

	_, err = io.Copy(w, f)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handlers) Slow(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	i := 0

	for i < 10 {
		select {
		case <-t.C:
			h.log.Info(ctx, "slow req tick", "i", i)
			i++
		case <-ctx.Done():
			h.log.Info(ctx, "slow req canceled", "err", ctx.Err())
			return nil
		}
	}

	h.log.Info(ctx, "slow req done")

	return h.Query(ctx, w, r)
}

// /////////////////////////////////////////////////////////////////

func parseDeleteForm(r *http.Request) ([]int, error) {
	maxFormSize := int64(10 << 20)
	bodyBytes, err := io.ReadAll(io.LimitReader(r.Body, maxFormSize))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if n, _ := io.ReadFull(r.Body, make([]byte, 1)); n == 1 {
		return nil, fmt.Errorf("request body too large")
	}

	bodyString := string(bodyBytes)

	data, err := url.ParseQuery(bodyString)
	if err != nil {
		return nil, err
	}

	toDelete := data["selected_contact_ids"]
	ids := make([]int, len(toDelete))

	for i, id := range toDelete {
		id, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}

		ids[i] = id
	}

	return ids, nil
}
