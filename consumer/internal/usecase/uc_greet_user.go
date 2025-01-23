package usecase

import (
	"fmt"
	"strconv"

	"github.com/mesirendon/contract-testing/consumer/internal/model"
)

type (
	userGetter interface {
		GetUser(id int) (model.User, error)
	}

	GreetUser struct {
		userService userGetter
	}
)

func NewGreetUser(userService userGetter) *GreetUser {
	return &GreetUser{
		userService: userService,
	}
}

func (uc *GreetUser) GreetUser(id string) (string, error) {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return "", fmt.Errorf("converting id to integer: %w", err)
	}

	user, err := uc.userService.GetUser(uid)
	if err != nil {
		return "", fmt.Errorf("fetching user from user service: %w", err)
	}

	return fmt.Sprintf("Hello %s %s!", user.FirstName, user.LastName), nil
}
