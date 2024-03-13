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

	fe "github.com/adamwoolhether/hypermedia/app/frontend/view/contacts"
	"github.com/adamwoolhether/hypermedia/foundation/session"
	"github.com/adamwoolhether/hypermedia/foundation/validate"

	"github.com/adamwoolhether/hypermedia/business/contacts"
	"github.com/adamwoolhether/hypermedia/foundation/logger"
	"github.com/adamwoolhether/hypermedia/foundation/web"
)

// Handlers manages the set of check points.
type Handlers struct {
	log      *logger.Logger
	core     *contacts.Core
	sessions *session.Store
}

func New(build string, log *logger.Logger, core *contacts.Core, store *session.Store) *Handlers {
	return &Handlers{
		log:      log,
		core:     core,
		sessions: store,
	}
}

func (h *Handlers) CreateForm(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	return fe.NewForm(fe.NewContact{}).Render(ctx, w)
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

		return fe.NewForm(newContact).Render(ctx, w)
	}

	err := h.core.Create(ctx, newContact.ToDB())
	if err != nil {
		newContact.InternalErrors = err.Error()
		return fe.NewForm(newContact).Render(ctx, w)
	}

	if err := h.sessions.AddFlash(w, r, "Created contact!"); err != nil {
		h.log.Error(ctx, "adding flash", "err", err)
	}

	web.Redirect(w, r, "/v1/contacts")

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

	contacts, err := h.core.Query(ctx, query, page)
	if err != nil {
		return err
	}

	// If the specific trigger header is present, then
	// we know we only need to update a very specific
	// part of the page, so we just return the rows.
	if r.Header.Get("HX-Trigger") == "search" {
		return fe.Rows(contacts).Render(ctx, w)
	}

	flashCtx := h.sessions.GetFlashCtx(w, r)
	return fe.Index(query, page, contacts, h.core.ArchivePoll(ctx)).Render(flashCtx, w)
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

	return fe.ShowByID(contact).Render(ctx, w)
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
	c := fe.UpdateContact{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
		Email:     contact.Email,
		//FieldErrs:      fe.ContactErrors{},
		//InternalErrors: "",
	}

	return fe.EditByID(c).Render(ctx, w)
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

		return fe.EditByID(uc).Render(ctx, w)
	}

	err = h.core.Update(ctx, uc.ToDB())
	if err != nil {
		// Or do failure flash here?
		uc.InternalErrors = err.Error()
		return fe.EditByID(uc).Render(ctx, w)
	}

	if err := h.sessions.AddFlash(w, r, "Updated contact!"); err != nil {
		h.log.Error(ctx, "adding flash", "err", err)
	}

	web.Redirect(w, r, "/v1/contacts"+userID)

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

	web.Redirect(w, r, "/v1/contacts")
	return nil
}

func (h *Handlers) DeleteBatch(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// Not sure why the req isn't settings the form... Should investigate later.
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)

	data, err := url.ParseQuery(bodyString)
	if err != nil {
		return err
	}

	toDelete := data["selected_contact_ids"]
	for _, id := range toDelete {
		id, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		if err := h.core.Delete(ctx, id); err != nil {
			h.log.Error(ctx, "deleting contact", "id", id, "err", err)
		}
	}

	if err := h.sessions.AddFlash(w, r, "Deleted contacts!"); err != nil {
		h.log.Error(ctx, "adding flash", "err", err)
	}

	contacts, err := h.core.Query(ctx, "", 1)
	if err != nil {
		return err
	}

	flashCtx := h.sessions.GetFlashCtx(w, r)
	return fe.Index("", 1, contacts, h.core.ArchivePoll(ctx)).Render(flashCtx, w)
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

	return fe.Archive(h.core.ArchivePoll(ctx)).Render(ctx, w)
}

func (h *Handlers) ArchiveRm(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	h.core.ArchiveRm(ctx)

	return fe.Archive(h.core.ArchivePoll(ctx)).Render(ctx, w)
}

func (h *Handlers) ArchivePoll(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	return fe.Archive(h.core.ArchivePoll(ctx)).Render(ctx, w)
}

func (h *Handlers) ArchiveDL(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	file := h.core.ArchiveFile(ctx)

	f, err := os.Open(file)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, f)
	if err != nil {
		return err
	}

	return nil
}
