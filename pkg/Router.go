package pkg

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/steffakasid/go-rest-api/pkg/handler"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/",
		handler.Index,
	},
	Route{
		"GET",
		"/health",
		handler.GetHealth,
	},
	Route{
		"POST",
		"/json",
		handler.ReceiveJson,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Handler(route.Handler)

	}
	return router
}
