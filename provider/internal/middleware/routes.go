package middleware

import (
	"net/http"

	"github.com/mesirendon/contract-testing/provider/internal/dependency"
	"github.com/mesirendon/contract-testing/provider/internal/model"
)

func setRoutes(mux *http.ServeMux, db *map[int]model.User) {
	mux.HandleFunc("/user/", commonMiddleware(func(w http.ResponseWriter, r *http.Request) {
		dependency.UserGetterHandler(w, r, db)
	}))
}
