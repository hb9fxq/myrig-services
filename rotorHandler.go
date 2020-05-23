package main

import (
	"encoding/json"
	"net/http"
)

func Rotor(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(appCtx.RotorStatus); err != nil {
		panic(err)
	}
}
