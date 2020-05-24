package handlers

import (
	"github.com/krippendorf/myrig-services/globals"
	"log"
	"net/http"
)

func HttpAuthChainHandler(ctx *globals.ApplicationContext, inner http.Handler, route globals.Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if route.IsPublic {
			inner.ServeHTTP(w, r)
			return
		}

		var apiKey = getAuthMyrigKey(r)
		var apiUser = GetAuthMyrigUser(r)

		var userAuthed = false

		for _, rKey := range ctx.ApiKeys {
			if rKey.User == apiUser && rKey.Key == apiKey {
				userAuthed = true
				break
			}
		}

		if userAuthed {
			log.Printf(
				"AUTHORIZED User:%s Method: %s RequestURI:%s Route:%s IP:%s",
				apiUser,
				r.Method,
				r.RequestURI,
				route.Name,
				GetIP(r),
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

			http.Error(w, "Not authorized", 401)
		}

	}
}

func GetAuthMyrigUser(r *http.Request) string {
	return r.Header.Get("X-myrig-user")
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
