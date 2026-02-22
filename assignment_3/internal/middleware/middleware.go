package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s | %s | %s", start.Format(time.RFC3339), r.Method, r.URL.Path)
	})
}

func APIKeyAuthWithKey(next http.Handler, validKey string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-KEY")
		if key == "" || key != validKey {
			http.Error(w, "Unauthorized - Invalid API Key", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
