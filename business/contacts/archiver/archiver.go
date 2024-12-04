package archiver

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"
)

type ArchiveView struct {
	Status  Status `json:"status"`
	Percent string `json:"progress"`
}

type Archiver struct {
	mu       sync.RWMutex
	Status   Status
	Progress atomic.Uint32
}

type Status int

const (
	Waiting Status = iota + 1
	Running
	Complete
)

func New() *Archiver {
	a := Archiver{
		Status: Waiting,
	}

	return &a
}

func (a *Archiver) Run() error {
	a.mu.Lock()
	{
		a.Status = Running
	}
	a.mu.Unlock()

	jobFn := func(ctx context.Context) {
		for i := range 100 {
			delay, err := rand.Int(rand.Reader, big.NewInt(100))
			if err != nil {
				delay = big.NewInt(0)
			}

			time.Sleep(time.Duration(delay.Int64()) * time.Millisecond)

			a.Progress.Swap(uint32(i))
		}

		a.mu.Lock()
		defer a.mu.Unlock()
		if a.Status != Running {
			return
		}

		a.Status = Complete
	}

	go jobFn(context.Background())

	return nil
}

func (a *Archiver) Poll() ArchiveView {
	a.mu.RLock()
	defer a.mu.RUnlock()

	status := a.Status
	progress := a.Progress.Load()

	apiResp := ArchiveView{
		Status:  status,
		Percent: fmt.Sprintf("%d", progress),
	}

	return apiResp
}

func (a *Archiver) File() string {
	return "business/contacts/contacts.json"
}

func (a *Archiver) Reset() {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Status = Waiting
	a.Progress.Store(0)
}
