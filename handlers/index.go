package handlers

import (
	"encoding/json"
	"github.com/krippendorf/myrig-services/globals"
	"log"
	"net/http"
)

type AppStatus struct {
	Status    string `json:"status"`
	Connected bool   `json:"connected"`
}

func IndexHandler(route globals.Route) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Printf("index served %s - route name: %s", route.AppCtx.RotorstatusUrl, route.Name)
		appStatus := AppStatus{Status: "Active", Connected: true}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(appStatus); err != nil {
			panic(err)
		}
	}
}
