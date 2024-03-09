package web

import "net/http"

func Redirect(w http.ResponseWriter, r *http.Request, path string) {
	rawQuery := r.URL.RawQuery
	if rawQuery != "" {
		path += "?" + rawQuery
	}

	http.Redirect(w, r, path, http.StatusSeeOther)
}
