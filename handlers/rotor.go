package handlers

import (
	"encoding/json"
	"github.com/hb9fxq/myrig-services/globals"
	"net/http"
)

func RotorHandler(route *globals.Route) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		w.WriteHeader(http.StatusOK)

		m := make(map[string]*globals.RotorStatusType)
		m["Optibeam"] = route.AppCtx.RotorStatus
		m["MagneticLoop"] = route.AppCtx.LoopRotorStatus

		if err := json.NewEncoder(w).Encode(m); err != nil {
			panic(err)
		}
	}
}
