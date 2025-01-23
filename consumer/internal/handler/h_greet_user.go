package handler

import (
	"net/http"
	"strings"
)

type (
	userGreeter interface {
		GreetUser(id string) (string, error)
	}

	GreetUser struct {
		userGreeter userGreeter
	}
)

func NewGreetUser(userGreeter userGreeter) *GreetUser {
	return &GreetUser{
		userGreeter: userGreeter,
	}
}

func (h *GreetUser) GreetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	a := strings.Split(r.URL.Path, "/")
	id := a[len(a)-1]

	g, err := h.userGreeter.GreetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(g))
	}
}
