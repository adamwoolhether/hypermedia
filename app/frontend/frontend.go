package frontend

import (
	"embed"
	"net/http"
)

//go:embed static
var publicFS embed.FS

func Static() http.Handler {
	return http.FileServerFS(publicFS)
}
