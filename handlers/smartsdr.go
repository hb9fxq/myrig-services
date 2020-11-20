package handlers

import (
	"encoding/json"
	"github.com/hb9fxq/myrig-services/globals"
	"net/http"
)

func SliceHandler(route *globals.Route) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(route.AppCtx.SmartSdrSlice); err != nil {
			panic(err)
		}
	}
}

func PanadapterHandler(route *globals.Route) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(route.AppCtx.SmartSdrPanadapter); err != nil {
			panic(err)
		}
	}
}
