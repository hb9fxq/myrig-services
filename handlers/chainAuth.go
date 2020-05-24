package handlers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/krippendorf/myrig-services/globals"
	"log"
	"net/http"
)

func HttpAuthChainHandler(inner http.Handler, route *globals.Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		myrigRequestSerial := fmt.Sprintf("mr-public-%s", uuid.New().String())

		if route.IsPublic {
			w.Header().Set("X-myrig-requestserial", myrigRequestSerial)
			r.Header.Set("X-myrig-requestserial", myrigRequestSerial)

			log.Printf(
				"AUTH:PUBLIC Method: %s RequestURI:%s Route:%s IP:%s RequestSerial:%s",
				r.Method,
				r.RequestURI,
				route.Name,
				GetIP(r),
				myrigRequestSerial,
			)

			inner.ServeHTTP(w, r)
			return
		}

		var apiKey = getAuthMyrigKey(r)
		var apiUser = GetAuthMyrigUser(r)

		var userAuthed = false

		for _, rKey := range route.AppCtx.ApiKeys {
			if rKey.User == apiUser && rKey.Key == apiKey {
				userAuthed = true
				myrigRequestSerial = fmt.Sprintf("mr-%s-%s", rKey.User, uuid.New().String())
				w.Header().Set("X-myrig-requestserial", myrigRequestSerial)
				r.Header.Set("X-myrig-requestserial", myrigRequestSerial)
				break
			}
		}

		if userAuthed {
			log.Printf(
				"AUTH:AUTHORIZED User:%s Method: %s RequestURI:%s Route:%s IP:%s RequestSerial:%s",
				apiUser,
				r.Method,
				r.RequestURI,
				route.Name,
				GetIP(r),
				myrigRequestSerial,
			)

			inner.ServeHTTP(w, r)
		} else {
			log.Printf(
				"NOT AUTHORIZED (401): Method: %s RequestURI:%s Route:%s IP:%s",
				r.Method,
				r.RequestURI,
				route.Name,
				GetIP(r),
			)

			w.WriteHeader(http.StatusUnauthorized)
		}

	}
}

func GetAuthMyrigUser(r *http.Request) string {
	return r.Header.Get("X-myrig-user")
}

func GetAuthMyrigRequestSerial(r *http.Request) string {
	return r.Header.Get("X-myrig-requestserial")
}

func getAuthMyrigKey(r *http.Request) string {
	return r.Header.Get("X-myrig-key")
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
