package usecase

import (
	"fmt"
	"strconv"
	"time"

	"github.com/mesirendon/contract-testing/consumer/internal/model"
)

type (
	userGetter interface {
		GetUser(id int) (model.User, error)
		SetToken(token string)
	}

	GreetUser struct {
		userService userGetter
		now         func() time.Time
	}
)

var (
	timeFormat = "2006-01-02T15:04"
)

func NewGreetUser(userService userGetter) *GreetUser {
	return &GreetUser{
		userService: userService,
		now:         time.Now,
	}
}

func (uc *GreetUser) GreetUser(id string) (string, error) {
	uid, err := strconv.Atoi(id)
	if err != nil {
		return "", fmt.Errorf("converting id to integer: %w", err)
	}

	now := uc.now().Format(timeFormat)
	uc.userService.SetToken(now)

	user, err := uc.userService.GetUser(uid)
	if err != nil {
		return "", fmt.Errorf("fetching user from user service: %w", err)
	}

	return fmt.Sprintf("Hello %s %s!", user.FirstName, user.LastName), nil
}
