package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mesirendon/contract-testing/provider/internal/model"
)

type (
	userGetter interface {
		GetUser(id string) (model.User, error)
	}

	GetUser struct {
		userGetter userGetter
	}
)

func NewGetUser(userGetter userGetter) *GetUser {
	return &GetUser{
		userGetter: userGetter,
	}
}

func (h GetUser) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	a := strings.Split(r.URL.Path, "/")
	id := a[len(a)-1]

	u, err := h.userGetter.GetUser(id)
	if err != nil {
		switch err.Error() {
		case "not found":
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}
	w.WriteHeader(http.StatusOK)
	usr, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(usr)
}
