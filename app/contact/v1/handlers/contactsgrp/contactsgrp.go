package contactsgrp

import (
	"context"
	"fmt"
	"net/http"

	fe "github.com/adamwoolhether/hypermedia/app/frontend/view/contacts"
	"github.com/adamwoolhether/hypermedia/foundation/validate"

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

	fmt.Println("VALDATE")
	if err := validate.Check(newContact); err != nil {
		fieldErrs := validate.GetFieldErrors(err)

		newContact.FieldErrs = fe.NewContactErrors{
			First: fieldErrs.Fields()["first_name"],
			Last:  fieldErrs.Fields()["last_name"],
			Phone: fieldErrs.Fields()["phone"],
			Email: fieldErrs.Fields()["email"],
		}

		return fe.NewForm(newContact).Render(ctx, w)
	}

	// do some validation here

	// need generalized error for internal stuff, duplicate users, etc.
	err := h.core.Create(ctx, newContact.ToDB())
	if err != nil {
		newContact.InternalErrors = err.Error()
		return fe.NewForm(newContact).Render(ctx, w)
	}

	// FLASH HERE, whatever that is
	web.Redirect(w, r, "/contacts")

	return nil
}

func (h *Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	query := web.QueryString(r, "q")

	contacts, err := h.core.Query(ctx, query)
	if err != nil {
		return err
	}

	return fe.Index(query, contacts).Render(ctx, w)
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
