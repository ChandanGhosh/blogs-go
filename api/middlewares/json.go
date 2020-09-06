package middlewares

import (
	"net/http"
)

// SetJSONContentType ..
func SetJSONContentType(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		next(w, r)
	}
}
