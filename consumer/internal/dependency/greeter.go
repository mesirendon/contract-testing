package dependency

import (
	"net/http"

	"github.com/mesirendon/contract-testing/consumer/internal/handler"
	"github.com/mesirendon/contract-testing/consumer/internal/services/users"
	"github.com/mesirendon/contract-testing/consumer/internal/usecase"
)

func GreeterHandler(w http.ResponseWriter, r *http.Request) {
	us, err := users.NewUsersClient("http://localhost:8082")
	if err != nil {
		panic(err)
	}

	uc := usecase.NewGreetUser(us)
	h := handler.NewGreetUser(uc)

	h.GreetUser(w, r)
}
