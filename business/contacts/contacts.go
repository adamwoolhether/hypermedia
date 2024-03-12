package contacts

import (
	"context"
	"errors"
	"os"
	"slices"
	"strings"
	"sync"

	"github.com/go-json-experiment/json"

	"github.com/adamwoolhether/hypermedia/foundation/logger"
)

const fileDB = "business/contacts/contacts.json"

// Core manages the set of API's for user access.
type Core struct {
	log *logger.Logger
	db  []Contact
	mu  sync.RWMutex
}

// NewCore constructs a core for the user api access.
func NewCore(log *logger.Logger) *Core {
	bytes, err := os.ReadFile(fileDB)
	if err != nil {
		panic(err)
	}

	var contacts []Contact
	if err := json.Unmarshal(bytes, &contacts); err != nil {
		panic(err)
	}

	c := &Core{
		log: log,
		db:  contacts,
	}

	return c
}

const pageSize = 10

func (c *Core) Query(ctx context.Context, query string, page int) ([]Contact, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	start := (page - 1) * pageSize
	end := start + pageSize

	var results []Contact

	if query == "" {
		results = c.db
	} else {
		c.log.Info(ctx, "searching", "query", query)

		for i := range c.db {
			if strings.Contains(c.db[i].FirstName, query) || strings.Contains(c.db[i].LastName, query) || strings.Contains(c.db[i].Email, query) || strings.Contains(c.db[i].Phone, query) {
				results = append(results, c.db[i])
			}
		}
	}

	if start > len(c.db) {
		return []Contact{}, nil
	}
	if end > len(results) {
		end = len(results)
	}

	return results[start:end], nil
}

func (c *Core) QueryByID(ctx context.Context, id int) (Contact, error) {
	c.log.Info(ctx, "searching by id", "id", id)

	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, contact := range c.db {
		if contact.ID == id {
			return contact, nil
		}
	}

	return Contact{}, errors.New("not found")
}

func (c *Core) UniqueEmail(ctx context.Context, id int, email string) bool {
	c.log.Info(ctx, "checking unique email", "email", email)

	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, contact := range c.db {
		if contact.Email == email && contact.ID != id {
			return false
		}
	}

	return true
}

func (c *Core) Create(ctx context.Context, newContact Contact) error {
	c.log.Info(ctx, "creating", "newContact", newContact)

	c.mu.Lock()
	defer c.mu.Unlock()
	latestID := c.db[len(c.db)-1].ID

	newContact.ID = latestID + 1

	c.db = append(c.db, newContact)

	return nil
}

func (c *Core) Update(ctx context.Context, contact Contact) error {
	c.log.Info(ctx, "updating", "contact", contact)

	c.mu.Lock()
	defer c.mu.Unlock()

	for i := range c.db {
		if c.db[i].ID == contact.ID {
			c.db[i] = contact

			return nil
		}
	}

	return errors.New("contact not found")
}

func (c *Core) Delete(ctx context.Context, id int) error {
	c.log.Info(ctx, "deleting", "id", id)

	c.mu.Lock()
	defer c.mu.Unlock()

	for i := range c.db {
		if c.db[i].ID == id {
			c.db = slices.Delete(c.db, i, i+1)
			return nil
		}
	}

	return errors.New("contact not found")
}

func (c *Core) Count(ctx context.Context) int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.db)
}
