package main

import (
	"github.com/hb9fxq/myrig-services/globals"
	"github.com/hb9fxq/myrig-services/handlers"
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
			globals.GlobalAppCtx,
		},
		globals.Route{
			"rotor",
			"GET",
			"/rotor",
			true,
			handlers.RotorHandler,
			globals.GlobalAppCtx,
		},
		globals.Route{
			"rotor",
			"GET",
			"/ant",
			true,
			handlers.AntHandler,
			globals.GlobalAppCtx,
		}, globals.Route{
			"rotor",
			"GET",
			"/kiwi",
			true,
			handlers.KiwistatusHandler,
			globals.GlobalAppCtx,
		}, globals.Route{
			"debug",
			"GET",
			"/debug",
			false,
			handlers.DebugHandler,
			globals.GlobalAppCtx,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for idx, _ := range routes {
		route := routes[idx]

		var handler http.Handler
		handler = route.Handler(&route)
		handler = addCommonHandlers(handler, &route)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

		router.Use()
	}
	return router
}

func addCommonHandlers(handler http.Handler, route *globals.Route) http.Handler {

	handler = handlers.LoggerChainHandler(handler, route)
	handler = handlers.HttpAuthChainHandler(handler, route)

	return handler
}
