package middlewares

import (
	"log"
	"net/http"
)

func RequestLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[REQUEST] %v %v \n", r.URL, r.Method)
		next.ServeHTTP(w, r)
	})
}
