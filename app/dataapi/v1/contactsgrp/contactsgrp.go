package contactsgrp

import (
	"context"
	"net/http"

	"github.com/adamwoolhether/hypermedia/business/contacts"
	"github.com/adamwoolhether/hypermedia/foundation/logger"
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

	return nil
}
