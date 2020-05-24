package handlers

import (
	"encoding/json"
	"github.com/krippendorf/myrig-services/globals"
	"log"
	"net/http"
)

func DubugHandler(route globals.Route) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Printf("DEBUG served for user: %s", getAuthMyrigKey(r))

		appStatus := AppStatus{Status: "DEBUG - Hello, " + GetAuthMyrigUser(r) + "!", Connected: true}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(appStatus); err != nil {
			panic(err)
		}
	}
}
