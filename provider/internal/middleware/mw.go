package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mesirendon/contract-testing/provider/internal/model"
)

const (
	timeFormat = "2006-01-02T15:04"
)

func GetHTTPHandler(p *map[int]model.User) *http.ServeMux {
	mux := http.NewServeMux()

	setRoutes(mux, p)

	return mux
}

func commonMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return withCorrelationID(isAuthenticated(f))
}

func withCorrelationID(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid.New()
		w.Header().Set("X-Api-Correlation-Id", uuid.String())
		h.ServeHTTP(w, r)
	}
}

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

func getAuthToken() string {
	return fmt.Sprintf("Bearer %s", time.Now().Format(timeFormat))
}
