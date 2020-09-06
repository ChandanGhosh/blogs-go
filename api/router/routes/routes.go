package routes

import (
	"net/http"

	"github.com/chandanghosh/blog/api/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	URI     string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

// LoadRoutes ..
func LoadRoutes() []Route {
	return userRoutes
}

// SetupRoutes ..
func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range LoadRoutes() {
		r.HandleFunc(route.URI, middlewares.SetLogger(middlewares.SetJSONContentType(route.Handler))).Methods(route.Method)
	}
	return r
}
