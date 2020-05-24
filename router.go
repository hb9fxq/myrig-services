package main

import (
	"github.com/krippendorf/myrig-services/globals"
	"github.com/krippendorf/myrig-services/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {

	var routes = []globals.Route{
		globals.Route{
			"Index",
			"GET",
			"/",
			true,
			handlers.IndexHandler,
			appCtx,
		},
		globals.Route{
			"rotor",
			"GET",
			"/rotor",
			true,
			handlers.RotorHandler,
			appCtx,
		}, globals.Route{
			"debug",
			"GET",
			"/debug",
			false,
			handlers.DubugHandler,
			appCtx,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.Handler(route)
		handler = addCommonHandlers(handler, route)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

		router.Use()
	}
	return router
}

func addCommonHandlers(handler http.Handler, route globals.Route) http.Handler {
	handler = handlers.HttpAuthChainHandler(appCtx, handler, route)
	handler = handlers.LoggerChainHandler(handler, route.Name)
	return handler
}
