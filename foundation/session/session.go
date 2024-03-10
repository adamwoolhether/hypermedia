package session

import (
	"context"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type contextKey int

const FlashContextKey contextKey = iota + 1

var (
	sessionKey = "default_session"
	flashKey   = "flash"
)

type Store struct {
	cs *sessions.CookieStore
}

func New(secretKey string) *Store {
	var key []byte
	if secretKey != "" {
		key = []byte(secretKey)
	} else {
		key = securecookie.GenerateRandomKey(32)
	}

	cs := sessions.NewCookieStore(key)

	store := Store{cs: cs}

	return &store
}

func (s *Store) AddFlash(w http.ResponseWriter, r *http.Request, value string) error {
	sesh, err := s.cs.Get(r, sessionKey)
	if err != nil {
		return err
	}

	sesh.AddFlash(value, flashKey)
	return sesh.Save(r, w)
}

func (s *Store) GetFlashCtx(w http.ResponseWriter, r *http.Request) context.Context {
	sesh, err := s.cs.Get(r, sessionKey)
	if err != nil {
		return r.Context()
	}

	flashes := sesh.Flashes(flashKey)
	if err := sesh.Save(r, w); err != nil {
		return r.Context()
	}

	flashStr := make([]string, 0, len(flashes))
	for _, flash := range flashes {
		str, ok := flash.(string)
		if !ok {
			continue
		}

		flashStr = append(flashStr, str)
	}

	ctx := context.WithValue(r.Context(), FlashContextKey, flashStr)

	return ctx
}
