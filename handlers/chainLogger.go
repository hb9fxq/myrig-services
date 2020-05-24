package handlers

import (
	"github.com/krippendorf/myrig-services/globals"
	"log"
	"net/http"
	"time"
)

func LoggerChainHandler(inner http.Handler, route *globals.Route) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s %s\t%s\t%s\t%s",
			GetAuthMyrigRequestSerial(r),
			r.Method,
			r.RequestURI,
			route.Name,
			time.Since(start),
		)
	}
}
