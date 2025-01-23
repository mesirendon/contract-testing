package users

import "github.com/mesirendon/contract-testing/consumer/internal/model"

type user struct {
	FirstName string `json:"firstName" pact:"example=John"`
	LastName  string `json:"lastName"  pact:"example=Doe"`
	Username  string `json:"username"  pact:"example=drwho"`
	Type      string `json:"type"      pact:"example=admin,regex^(admin|user|guest)$"`
	ID        int    `json:"id"        pact:"example=10"`
}

func (u user) toModel() model.User {
	return model.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Type:      u.Type,
		ID:        u.ID,
	}
}
