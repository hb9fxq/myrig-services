package handlers

import (
	"encoding/json"
	"github.com/hb9fxq/myrig-services/globals"
	"net/http"
)

type DebugInfo struct {
	AuthUser           string
	RotorUrl           string
	MyrigRequestSerial string
}

func DebugHandler(route *globals.Route) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		debugInfo := DebugInfo{AuthUser: GetAuthMyrigUser(r), RotorUrl: route.AppCtx.RotorstatusUrl, MyrigRequestSerial: GetAuthMyrigRequestSerial(r)}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(debugInfo); err != nil {
			panic(err)
		}
	}
}
