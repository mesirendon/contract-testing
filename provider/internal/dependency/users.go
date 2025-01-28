package dependency

import (
	"net/http"

	"github.com/mesirendon/contract-testing/provider/internal/handler"
	"github.com/mesirendon/contract-testing/provider/internal/model"
	"github.com/mesirendon/contract-testing/provider/internal/usecase"
)

func UserGetterHandler(w http.ResponseWriter, r *http.Request, db *map[int]model.User) {
	uc := usecase.NewGetUser(*db)
	h := handler.NewGetUser(uc)

	h.GetUser(w, r)
}
