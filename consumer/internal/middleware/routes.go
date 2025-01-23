package middleware

import (
	"net/http"

	"github.com/mesirendon/contract-testing/consumer/internal/dependency"
)

func setRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/user/", commonMiddleware(dependency.GreeterHandler))
}
