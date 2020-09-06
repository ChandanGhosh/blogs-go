package middlewares

import (
	"log"
	"net/http"
)

// SetLogger ..
func SetLogger(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\t%s %s%s %s\n", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}
