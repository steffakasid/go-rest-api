package main

import (
	"net/http"

	"github.com/steffakasid/rest-api/handler"
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
