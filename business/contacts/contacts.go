package contacts

import (
	"context"
	"os"
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

func (c *Core) Query(ctx context.Context, query string) ([]Contact, error) {
	if query == "" {
		c.mu.RLock()
		defer c.mu.RUnlock()
		return c.db, nil
	}

	c.log.Info(ctx, "searching", "query", query)

	var results []Contact

	c.mu.RLock()
	defer c.mu.RUnlock()
	for i := range c.db {
		if strings.Contains(c.db[i].First, query) || strings.Contains(c.db[i].Last, query) || strings.Contains(c.db[i].Email, query) || strings.Contains(c.db[i].Phone, query) {
			results = append(results, c.db[i])
		}
	}

	return results, nil
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
