package handlers

import (
	"encoding/json"
	"github.com/krippendorf/myrig-services/globals"
	"net/http"
)

func RotorHandler(route globals.Route) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(route.AppCtx.RotorStatus); err != nil {
			panic(err)
		}
	}
}
