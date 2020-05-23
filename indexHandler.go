package main

import (
	"encoding/json"
	"net/http"
)

type AppStatus struct {
	Status    string `json:"status"`
	Connected bool   `json:"connected"`
}

func Index(w http.ResponseWriter, r *http.Request) {

	appStatus := AppStatus{Status: "Active", Connected: true}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(appStatus); err != nil {
		panic(err)
	}
}
