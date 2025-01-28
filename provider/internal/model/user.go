package model

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Type      string `json:"type"`
	ID        int    `json:"id"`
}
