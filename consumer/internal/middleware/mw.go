package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const (
	timeFormat = "2006-01-02T15:04"
)

func isAuthenticated(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == getAuthToken() {
			h.ServeHTTP(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}

func withCorrelationID(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid.New()
		w.Header().Set("X-Api-Correlation-Id", uuid.String())
		h.ServeHTTP(w, r)
	}
}

func commonMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return withCorrelationID(isAuthenticated(f))
}

func getAuthToken() string {
	return fmt.Sprintf("Bearer %s", time.Now().Format(timeFormat))
}

func GetHTTPHandler() *http.ServeMux {
	mux := http.NewServeMux()

	setRoutes(mux)

	return mux
}
