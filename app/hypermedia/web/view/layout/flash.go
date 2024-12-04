package layout

import (
	"context"

	"github.com/adamwoolhether/hypermedia/foundation/session"
)

func getFlash(ctx context.Context) []string {
	flashes, ok := ctx.Value(session.FlashContextKey).([]string)
	if !ok {
		return []string{}
	}

	return flashes
}
