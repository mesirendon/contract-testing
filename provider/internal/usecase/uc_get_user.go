package usecase

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/mesirendon/contract-testing/provider/internal/model"
)

type (
	GetUser struct {
		db map[int]model.User
	}
)

func NewGetUser(db map[int]model.User) *GetUser {
	return &GetUser{
		db: db,
	}
}

func (uc *GetUser) GetUser(id string) (model.User, error) {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return model.User{}, fmt.Errorf("converting id: %w", err)
	}

	u, ok := uc.db[uid]
	if !ok {
		return model.User{}, errors.New("not found")
	}

	return u, nil
}
